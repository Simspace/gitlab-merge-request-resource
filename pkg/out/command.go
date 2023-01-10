package out

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Khan/genqlient/graphql"
	"github.com/simspace/gitlab-merge-request-resource/pkg/gitlab"
	"github.com/simspace/gitlab-merge-request-resource/pkg/models"
	gitlabv4 "github.com/xanzy/go-gitlab"
)

type Command struct {
	client   *graphql.Client
	clientv4 *gitlabv4.Client
}

func NewCommand(client *graphql.Client, clientv4 *gitlabv4.Client) *Command {
	return &Command{client, clientv4}
}

func (command *Command) Run(destination string, request Request) (Response, error) {
	repo := filepath.Join(destination, request.Params.Repository)
	err := os.MkdirAll(repo, 0755)
	if err != nil {
		return Response{}, err
	}

	err = os.Chdir(repo)
	if err != nil {
		return Response{}, err
	}

	file, err := os.ReadFile(".git/merge-request.json")
	if err != nil {
		return Response{}, err
	}

	var mr gitlab.MergeRequest
	err = json.Unmarshal(file, &mr)
	if err != nil {
		return Response{}, err
	}

	err = command.updateCommitStatus(request, mr)
	if err != nil {
		return Response{}, err
	}

	mr, err = command.updateLabels(request, mr)
	if err != nil {
		return Response{}, err
	}

	err = command.createNote(destination, request, mr)
	if err != nil {
		return Response{}, err
	}

	commit, err := mr.GetLatestCommit()
	if err != nil {
		return Response{}, err
	}

	response := Response{
		Version: models.Version{
			ID:        mr.Id,
			IID:       mr.Iid,
			UpdatedAt: &commit.AuthoredDate,
		},
		Metadata: buildMetadata(&mr),
	}

	return response, nil
}

func (command *Command) createNote(destination string, request Request, mr gitlab.MergeRequest) error {
	body, err := request.Params.Comment.ReadContent(destination)
	if err != nil {
		return err
	}

	if body != "" {
		in := gitlab.CreateNoteInput{
			Body:                    body,
			MergeRequestDiffHeadSha: mr.DiffHeadSha,
			NoteableId:              mr.Id,
		}
		_, err = gitlab.CreateNote(context.Background(), *command.client, in)
		if err != nil {
			return err
		}
	}
	return nil
}

// updateLabels appends labels to a merge request, if labels are supplied,
// and returns the updated MergeRequest object from the server
func (command *Command) updateLabels(request Request, mr gitlab.MergeRequest) (gitlab.MergeRequest, error) {
	updMr := mr
	if request.Params.Labels != nil {
		labelMap := map[string]string{}
		lresp, err := gitlab.ListLabels(context.Background(), *command.client, mr.GetProjectPath())
		for _, label := range lresp.Project.Labels.Nodes {
			labelMap[label.Title] = label.Id
		}
		labelIds := []string{}
		for _, label := range request.Params.Labels {
			if id, found := labelMap[label]; found {
				labelIds = append(labelIds, id)
			} else {
				return gitlab.MergeRequest{}, fmt.Errorf("could not find an existing label for %s", label)
			}
		}
		in := gitlab.MergeRequestSetLabelsInput{
			Iid:           mr.Iid,
			LabelIds:      labelIds,
			OperationMode: gitlab.MutationOperationModeAppend,
			ProjectPath:   mr.GetProjectPath(),
		}
		resp, err := gitlab.SetMergeRequestLabels(context.Background(), *command.client, in)
		if err != nil {
			return gitlab.MergeRequest{}, err
		}

		updMr = resp.MergeRequestSetLabels.GetMergeRequest()
	}
	return updMr, nil
}

func (command *Command) updateCommitStatus(request Request, mr gitlab.MergeRequest) error {
	if request.Params.Status != "" {
		state := gitlabv4.BuildState(gitlabv4.BuildStateValue(request.Params.Status))
		target := request.Source.GetTargetURL()
		name := request.Source.GetPipelineName()
		options := gitlabv4.SetCommitStatusOptions{
			Name:      &name,
			TargetURL: &target,
			State:     *state,
		}

		_, _, err := command.clientv4.Commits.SetCommitStatus(mr.SourceProjectId, mr.DiffHeadSha, &options)
		if err != nil {
			return err
		}
	}
	return nil
}

func buildMetadata(mr *gitlab.MergeRequest) models.Metadata {
	var labels []string
	for _, label := range mr.GetLabels().Nodes {
		labels = append(labels, label.Title)
	}

	return []models.MetadataField{
		{
			Name:  "id",
			Value: mr.Id,
		},
		{
			Name:  "iid",
			Value: mr.Iid,
		},
		{
			Name:  "sha",
			Value: mr.DiffHeadSha,
		},
		{
			Name:  "title",
			Value: mr.Title,
		},
		{
			Name:  "author",
			Value: mr.Author.Name,
		},
		{
			Name:  "source",
			Value: mr.SourceBranch,
		},
		{
			Name:  "target",
			Value: mr.TargetBranch,
		},
		{
			Name:  "url",
			Value: mr.WebUrl,
		},
		{
			Name:  "labels",
			Value: strings.Join(labels, ","),
		},
	}
}

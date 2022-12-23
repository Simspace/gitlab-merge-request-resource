package check

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/simspace/gitlab-merge-request-resource/pkg"
	"github.com/xanzy/go-gitlab"
)

type Command struct {
	client   *graphql.Client
	clientv4 *gitlab.Client
}

func NewCommand(client *graphql.Client, clientv4 *gitlab.Client) *Command {
	return &Command{
		client: client,
	}
}

func (command *Command) GetClient() graphql.Client {
	return *command.client
}

func (command *Command) Run(request Request) (Response, error) {
	labels := request.Source.Labels

	resp, err := pkg.GetProject(context.Background(), command.GetClient(), request.Source.GetProjectPath(), pkg.MergeRequestStateOpened)
	if err != nil {
		return Response{}, err
	}

	versions := make([]pkg.Version, 0)

	for _, mr := range resp.Project.GetMergeRequests().Nodes {
		if mr.DiffHeadSha == "" {
			continue
		}
		mrLabels := mr.GetLabels().Nodes
		if !matchLabels(labels, mrLabels) {
			continue
		}

		for _, commit := range mr.GetCommits().Nodes {
			if commit.Sha != mr.DiffHeadSha {
				continue
			}
			if strings.Contains(commit.Title, "[skip ci]") {
				continue
			}

			updatedAt := &commit.AuthoredDate

			if request.Version.UpdatedAt != nil && !updatedAt.After(*request.Version.UpdatedAt) {
				continue
			}

			statuses := commit.GetPipelines().Nodes

			// Only set status pending if no CI has already run on the commit
			if len(statuses) == 0 {
				name := request.Source.GetPipelineName()
				target := request.Source.GetTargetURL()
				options := gitlab.SetCommitStatusOptions{
					Name:      &name,
					TargetURL: &target,
					State:     gitlab.Pending,
				}
				_, _, _ = command.clientv4.Commits.SetCommitStatus(resp.Project.GetId(), commit.GetSha(), &options)
			}

			IIDStr, err := strconv.Atoi(mr.Iid)
			if err != nil {
				return Response{}, err
			}
			versions = append(versions, pkg.Version{ID: IIDStr, UpdatedAt: updatedAt})
		}

	}

	return versions, nil
}

func matchLabels(sourceLabels []string, mrLabels []pkg.Label) bool {
	if len(sourceLabels) == 0 {
		return true
	}
	if len(mrLabels) == 0 {
		return false
	}
	for _, label := range sourceLabels {
		for _, mrLabel := range mrLabels {
			if mrLabel.Title == label {
				return true
			}
		}
	}
	return false
}

func matchPathPatterns(api *gitlab.Client, mr *gitlab.MergeRequest, source pkg.Source) (bool, error) {

	if len(source.Paths) == 0 && len(source.IgnorePaths) == 0 {
		return true, nil
	}

	modified := 0

	versions, _, err := api.MergeRequests.GetMergeRequestDiffVersions(mr.ProjectID, mr.IID, nil)
	if err != nil {
		return false, err
	}

	if len(versions) > 0 {

		latest := versions[0].ID
		version, _, err := api.MergeRequests.GetSingleMergeRequestDiffVersion(mr.ProjectID, mr.IID, latest)
		if err != nil {
			return false, err
		}

		for _, d := range version.Diffs {
			if source.AcceptPath(d.OldPath) || source.AcceptPath(d.NewPath) {
				modified += 1
			}
		}
	}

	return modified > 0, nil
}

func getMostRecentUpdateTime(notes []*gitlab.Note, updatedAt *time.Time) *time.Time {
	for _, note := range notes {
		if strings.Contains(note.Body, "[trigger ci]") && updatedAt.Before(*note.UpdatedAt) {
			return note.UpdatedAt
		}
	}
	return updatedAt
}

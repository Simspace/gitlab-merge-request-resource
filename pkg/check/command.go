package check

import (
	"context"
	"strings"
	"time"

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
	return &Command{
		client:   client,
		clientv4: clientv4,
	}
}

func (command *Command) Run(request Request) (Response, error) {
	labels := request.Source.Labels

	resp, err := gitlab.ListMergeRequests(context.Background(), *command.client, request.Source.GetProjectPath(), gitlab.MergeRequestStateOpened)
	if err != nil {
		return Response{}, err
	}

	versions := make([]models.Version, 0)

	for _, mr := range resp.Project.GetMergeRequests().Nodes {
		if request.Source.SourceBranch != "" && mr.SourceBranch != request.Source.SourceBranch {
			continue
		}

		if request.Source.TargetBranch != "" && mr.TargetBranch != request.Source.TargetBranch {
			continue
		}

		if request.Source.SkipNotMergeable && mr.MergeStatusEnum != gitlab.MergeStatusCanBeMerged {
			continue
		}

		if request.Source.SkipWorkInProgress && mr.Draft {
			continue
		}

		mrLabels := mr.GetLabels().Nodes
		if !matchLabels(labels, mrLabels) {
			continue
		}

		match, err := matchPathPatterns(&mr, request.Source)
		if err != nil {
			return nil, err
		}

		if !match {
			continue
		}

		var updatedAt *time.Time

		commit, err := mr.GetLatestCommit()
		if err != nil {
			return Response{}, err
		}

		if strings.Contains(commit.Title, "[skip ci]") || strings.Contains(commit.Message, "[skip ci]") {
			continue
		}

		updatedAt = &commit.AuthoredDate

		// Enable manual triggering of the pipeline via special comments on the MR
		if !request.Source.SkipTriggerComment {
			notes := mr.GetNotes().Nodes
			updatedAt = getMostRecentUpdateTime(notes, updatedAt)
		}

		if request.Version.UpdatedAt != nil && !updatedAt.After(*request.Version.UpdatedAt) {
			continue
		}

		statuses := commit.GetPipelines().Nodes

		// Only set status pending if no CI has already run on the commit
		if len(statuses) == 0 {
			name := request.Source.GetPipelineName()
			target := request.Source.GetTargetURL()
			options := gitlabv4.SetCommitStatusOptions{
				Name:      &name,
				TargetURL: &target,
				State:     gitlabv4.Pending,
			}
			projID := stripID(resp.Project.GetId())
			_, _, _ = command.clientv4.Commits.SetCommitStatus(projID, commit.GetSha(), &options)
		}

		versions = append(versions, models.Version{ID: mr.Id, IID: mr.Iid, UpdatedAt: updatedAt})
	}

	return versions, nil
}

func stripID(s string) string {
	return strings.ReplaceAll(s, "git://gitlab/Project/", "")
}

func matchLabels(sourceLabels []string, mrLabels []gitlab.Label) bool {
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

func matchPathPatterns(mr *gitlab.MergeRequest, source models.Source) (bool, error) {

	if len(source.Paths) == 0 && len(source.IgnorePaths) == 0 {
		return true, nil
	}

	modified := 0

	diffStats := mr.GetDiffStats()

	if len(diffStats) > 0 {
		for _, d := range diffStats {
			if source.AcceptPath(d.Path) {
				modified++
			}
		}
	}

	return modified > 0, nil
}

func getMostRecentUpdateTime(notes []gitlab.Note, updatedAt *time.Time) *time.Time {
	for _, note := range notes {
		if strings.Contains(note.Body, "[trigger ci]") && updatedAt.Before(note.UpdatedAt) {
			return &note.UpdatedAt
		}
	}
	return updatedAt
}

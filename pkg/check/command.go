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

		match, err := matchPathPatterns(&mr, request.Source)
		if err != nil {
			return nil, err
		}

		if !match {
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
				projID := stripID(resp.Project.GetId())
				_, _, _ = command.clientv4.Commits.SetCommitStatus(projID, commit.GetSha(), &options)
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

func stripID(s string) string {
	return strings.ReplaceAll(s, "git://gitlab/Project/", "")
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

func matchPathPatterns(mr *pkg.MergeRequest, source pkg.Source) (bool, error) {

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

func getMostRecentUpdateTime(notes []*gitlab.Note, updatedAt *time.Time) *time.Time {
	for _, note := range notes {
		if strings.Contains(note.Body, "[trigger ci]") && updatedAt.Before(*note.UpdatedAt) {
			return note.UpdatedAt
		}
	}
	return updatedAt
}

package in

import (
	"context"
	"encoding/json"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/Khan/genqlient/graphql"
	"github.com/simspace/gitlab-merge-request-resource/pkg/gitlab"
	"github.com/simspace/gitlab-merge-request-resource/pkg/models"
)

type Command struct {
	client *graphql.Client
	runner GitRunner
	agent  AgentRunner
}

func NewCommand(client *graphql.Client) *Command {
	return &Command{
		client,
		NewRunner(),
		NewAgentRunner(),
	}
}

func (command *Command) WithRunner(runner GitRunner) *Command {
	command.runner = runner
	return command
}

func (command *Command) Run(destination string, request Request) (Response, error) {
	err := os.MkdirAll(destination, 0755)
	if err != nil {
		return Response{}, err
	}

	user, err := gitlab.GetCurrentUser(context.Background(), *command.client)
	err = command.runner.Run("config", "--global", "user.email", user.CurrentUser.GetPublicEmail())
	if err != nil {
		return Response{}, err
	}

	err = command.runner.Run("config", "--global", "user.name", user.CurrentUser.GetName())
	if err != nil {
		return Response{}, err
	}

	resp, err := gitlab.GetMergeRequest(context.Background(), *command.client, request.Version.ID)
	if err != nil {
		return Response{}, err
	}

	mr := resp.GetMergeRequest()

	target, err := command.createRepositoryURL(mr.GetTargetProject().HttpUrlToRepo, request.Source.PrivateToken)
	if err != nil {
		return Response{}, err
	}
	source, err := command.createRepositoryURL(mr.GetSourceProject().HttpUrlToRepo, request.Source.PrivateToken)
	if err != nil {
		return Response{}, err
	}

	commit, err := mr.GetLatestCommit()
	if err != nil {
		return Response{}, err
	}

	err = command.runner.Run("clone", "-c", "http.sslVerify="+strconv.FormatBool(!request.Source.Insecure), "-o", "target", "-b", mr.TargetBranch, target.String(), destination)
	if err != nil {
		return Response{}, err
	}

	if (request.Source.SshKeys != nil) && (len(request.Source.SshKeys) != 0) {
		err = command.runner.Run("config", "--global", "core.sshCommand", "ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no")
		if err != nil {
			return Response{}, err
		}
		err = command.agent.Start()
		if err != nil {
			return Response{}, err
		}
		for _, key := range request.Source.SshKeys {
			err = command.agent.AddKey(key)
			if err != nil {
				return Response{}, err
			}
		}
	}

	err = os.Chdir(destination)
	if err != nil {
		return Response{}, err
	}

	err = command.runner.Run("remote", "add", "source", source.String())
	if err != nil {
		return Response{}, err
	}

	err = command.runner.Run("remote", "update")
	if err != nil {
		return Response{}, err
	}

	err = command.runner.Run("merge", "--no-ff", "--no-commit", mr.DiffHeadSha)
	if err != nil {
		return Response{}, err
	}

	if request.Source.Recursive {
		err = command.runner.Run("submodule", "update", "--init", "--recursive")
		if err != nil {
			return Response{}, err
		}
	}

	notes, err := json.Marshal(mr)
	if err != nil {
		return Response{}, err
	}
	err = os.WriteFile(".git/merge-request.json", notes, 0644)
	if err != nil {
		return Response{}, err
	}

	err = os.WriteFile(".git/merge-request-source-branch", []byte(mr.SourceBranch), 0644)
	if err != nil {
		return Response{}, err
	}

	err = os.Mkdir(".git/resource", 0755)
	if err != nil {
		return Response{}, err
	}

	var changedFiles []string
	for _, change := range mr.DiffStats {
		changedFiles = append(changedFiles, change.Path)
	}

	err = os.WriteFile(".git/resource/changed_files", []byte(strings.Join(changedFiles, "\n")), 0644)
	if err != nil {
		return Response{}, err
	}

	response := Response{Version: request.Version, Metadata: buildMetadata(mr, commit)}

	return response, nil
}

func (command *Command) createRepositoryURL(uri, token string) (*url.URL, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	u.User = url.UserPassword("gitlab-ci-token", token)

	return u, nil
}

func buildMetadata(mr gitlab.MergeRequest, commit gitlab.Commit) models.Metadata {
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
			Name:  "message",
			Value: commit.Title,
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
	}
}

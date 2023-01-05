package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/simspace/gitlab-merge-request-resource/pkg/check"
	"github.com/simspace/gitlab-merge-request-resource/pkg/common"
	"github.com/xanzy/go-gitlab"
)

func main() {

	var request check.Request
	inputRequest(&request)

	httpClient := http.Client{
		Transport: &common.AuthedTransport{
			Key:     request.Source.PrivateToken,
			Wrapped: http.DefaultTransport,
		},
	}
	client := graphql.NewClient("https://gitlab.com/api/graphql", &httpClient)
	clientv4, err := gitlab.NewClient(request.Source.PrivateToken, gitlab.WithHTTPClient(common.GetDefaultClient(request.Source.Insecure)), gitlab.WithBaseURL(request.Source.GetBaseURL()))
	if err != nil {
		common.Fatal("initializing gitlab client", err)
	}
	command := check.NewCommand(&client, clientv4)
	response, err := command.Run(request)
	if err != nil {
		common.Fatal("running command", err)
	}

	outputResponse(response)
}

func inputRequest(request *check.Request) {
	if err := json.NewDecoder(os.Stdin).Decode(request); err != nil {
		common.Fatal("reading request from stdin", err)
	}
}

func outputResponse(response check.Response) {
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		common.Fatal("writing response to stdout", err)
	}
}

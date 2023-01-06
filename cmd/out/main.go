package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/simspace/gitlab-merge-request-resource/pkg/common"
	"github.com/simspace/gitlab-merge-request-resource/pkg/out"
	gitlabv4 "github.com/xanzy/go-gitlab"
)

func main() {

	if len(os.Args) < 2 {
		println("usage: " + os.Args[0] + " <destination>")
		os.Exit(1)
	}

	destination := os.Args[1]

	var request out.Request
	inputRequest(&request)

	httpClient := http.Client{
		Transport: &common.AuthedTransport{
			Key:     request.Source.PrivateToken,
			Wrapped: http.DefaultTransport,
		},
	}
	client := graphql.NewClient("https://gitlab.com/api/graphql", &httpClient)
	clientv4, err := gitlabv4.NewClient(request.Source.PrivateToken, gitlabv4.WithHTTPClient(common.GetDefaultClient(request.Source.Insecure)), gitlabv4.WithBaseURL(request.Source.GetBaseURL()))
	if err != nil {
		common.Fatal("initializing gitlab client", err)
	}

	command := out.NewCommand(&client, clientv4)
	response, err := command.Run(destination, request)
	if err != nil {
		common.Fatal("running command", err)
	}

	outputResponse(response)
}

func inputRequest(request *out.Request) {
	if err := json.NewDecoder(os.Stdin).Decode(request); err != nil {
		common.Fatal("reading request from stdin", err)
	}
}

func outputResponse(response out.Response) {
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		common.Fatal("writing response to stdout", err)
	}
}

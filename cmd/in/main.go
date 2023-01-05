package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/simspace/gitlab-merge-request-resource/pkg/common"
	"github.com/simspace/gitlab-merge-request-resource/pkg/in"
)

func main() {

	if len(os.Args) < 2 {
		println("usage: " + os.Args[0] + " <destination>")
		os.Exit(1)
	}

	destination := os.Args[1]

	var request in.Request
	inputRequest(&request)

	httpClient := http.Client{
		Transport: &common.AuthedTransport{
			Key:     request.Source.PrivateToken,
			Wrapped: http.DefaultTransport,
		},
	}
	client := graphql.NewClient("https://gitlab.com/api/graphql", &httpClient)

	command := in.NewCommand(&client)
	response, err := command.Run(destination, request)
	if err != nil {
		common.Fatal("running command", err)
	}

	outputResponse(response)
}

func inputRequest(request *in.Request) {
	if err := json.NewDecoder(os.Stdin).Decode(request); err != nil {
		common.Fatal("reading request from stdin", err)
	}
}

func outputResponse(response in.Response) {
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		common.Fatal("writing response to stdout", err)
	}
}

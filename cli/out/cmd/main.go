package main

import (
	"encoding/json"
	"github.com/simspace/gitlab-merge-request-resource/pkg"
	"github.com/simspace/gitlab-merge-request-resource/pkg/out"
	"github.com/xanzy/go-gitlab"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		println("usage: " + os.Args[0] + " <destination>")
		os.Exit(1)
	}

	destination := os.Args[1]

	var request out.Request
	inputRequest(&request)

	client, err := gitlab.NewClient(request.Source.PrivateToken, gitlab.WithHTTPClient(pkg.GetDefaultClient(request.Source.Insecure)), gitlab.WithBaseURL(request.Source.GetBaseURL()))
	if err != nil {
		pkg.Fatal("initializing gitlab client", err)
	}

	command := out.NewCommand(client)
	response, err := command.Run(destination, request)
	if err != nil {
		pkg.Fatal("running command", err)
	}

	outputResponse(response)
}

func inputRequest(request *out.Request) {
	if err := json.NewDecoder(os.Stdin).Decode(request); err != nil {
		pkg.Fatal("reading request from stdin", err)
	}
}

func outputResponse(response out.Response) {
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		pkg.Fatal("writing response to stdout", err)
	}
}

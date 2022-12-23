package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/simspace/gitlab-merge-request-resource/pkg"
	"github.com/simspace/gitlab-merge-request-resource/pkg/check"
	"github.com/xanzy/go-gitlab"
)

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "bearer "+t.key)
	return t.wrapped.RoundTrip(req)
}

func main() {

	var request check.Request
	inputRequest(&request)

	httpClient := http.Client{
		Transport: &authedTransport{
			key:     request.Source.PrivateToken,
			wrapped: http.DefaultTransport,
		},
	}
	client := graphql.NewClient("https://gitlab.com/api/graphql", &httpClient)
	clientv4, err := gitlab.NewClient(request.Source.PrivateToken, gitlab.WithHTTPClient(pkg.GetDefaultClient(request.Source.Insecure)), gitlab.WithBaseURL(request.Source.GetBaseURL()))
	if err != nil {
		pkg.Fatal("initializing gitlab client", err)
	}
	command := check.NewCommand(&client, clientv4)
	response, err := command.Run(request)
	if err != nil {
		pkg.Fatal("running command", err)
	}

	outputResponse(response)
}

func inputRequest(request *check.Request) {
	if err := json.NewDecoder(os.Stdin).Decode(request); err != nil {
		pkg.Fatal("reading request from stdin", err)
	}
}

func outputResponse(response check.Response) {
	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		pkg.Fatal("writing response to stdout", err)
	}
}

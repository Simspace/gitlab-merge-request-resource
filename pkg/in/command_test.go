package in_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Khan/genqlient/graphql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/simspace/gitlab-merge-request-resource/pkg/gitlab"
	"github.com/simspace/gitlab-merge-request-resource/pkg/in"
	"github.com/simspace/gitlab-merge-request-resource/pkg/models"
)

var _ = Describe("In", func() {

	var (
		t           time.Time
		mux         *http.ServeMux
		server      *httptest.Server
		command     *in.Command
		root        *url.URL
		destination string
	)

	type changes []struct {
		OldPath     string `json:"old_path"`
		NewPath     string `json:"new_path"`
		AMode       string `json:"a_mode"`
		BMode       string `json:"b_mode"`
		Diff        string `json:"diff"`
		NewFile     bool   `json:"new_file"`
		RenamedFile bool   `json:"renamed_file"`
		DeletedFile bool   `json:"deleted_file"`
	}

	BeforeEach(func() {
		destination, _ = os.MkdirTemp("", "gitlab-merge-request-resource-in")
		t, _ = time.Parse(time.RFC3339, "2022-01-01T08:00:00Z")
		mux = http.NewServeMux()
		server = httptest.NewServer(mux)
		root, _ = url.Parse(server.URL)
		context, _ := url.Parse("/api/graphql")
		base := root.ResolveReference(context)
		client := graphql.NewClient(base.String(), nil)
		command = in.NewCommand(&client).WithRunner(newMockRunner(destination))

	})

	AfterEach(func() {
		defer os.Remove(destination)
		server.Close()
	})

	Describe("Run", func() {

		BeforeEach(func() {
			mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
				project, _ := url.Parse("namespace/project.git")
				uri := root.ResolveReference(project)
				mrResponse := gitlab.GetMergeRequestResponse{
					MergeRequest: gitlab.MergeRequest{
						Iid:         "88",
						Id:          "99",
						Author:      gitlab.MergeRequestAuthor{Name: "Tester"},
						DiffHeadSha: "abc",
						DiffStats: []gitlab.MergeRequestDiffStats{
							{
								Path: "/foo",
							},
							{
								Path: "/bar",
							},
						},
						SourceBranch: "source-branch",
						SourceProject: gitlab.MergeRequestSourceProject{
							HttpUrlToRepo: uri.String(),
						},
						TargetBranch: "target-branch",
						TargetProject: gitlab.MergeRequestTargetProject{
							HttpUrlToRepo: uri.String(),
						},
						Commits: gitlab.MergeRequestCommitsCommitConnection{
							Nodes: []gitlab.Commit{
								{
									Sha:          "abc",
									AuthoredDate: t,
								},
							},
						},
					},
				}
				userResponse := gitlab.GetCurrentUserResponse{
					CurrentUser: gitlab.User{
						Name:        "testuser",
						PublicEmail: "tester@test.com",
					},
				}

				body, _ := ioutil.ReadAll(r.Body)
				graphqlReq := graphql.Request{}
				_ = json.Unmarshal(body, &graphqlReq)

				var data interface{}

				if graphqlReq.OpName == "GetMergeRequest" {
					data = mrResponse
				} else {
					data = userResponse
				}
				output, _ := json.Marshal(graphql.Response{Data: data})
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(output)
			})
		})

		Context("When it has a minimal valid configuration", func() {

			It("Should clone repository", func() {
				project, _ := url.Parse("namespace/project.git")
				uri := root.ResolveReference(project)

				request := in.Request{
					Source: models.Source{
						URI:          uri.String(),
						PrivateToken: "$",
					},
					Version: models.Version{ID: "git://gitlab/MergeRequest/1"},
				}

				response, err := command.Run(destination, request)
				Expect(err).Should(BeNil())
				Expect(response.Metadata[0].Name).To(Equal("id"))
				Expect(response.Metadata[0].Value).To(Equal("99"))
				_, err = os.Stat(filepath.Join(destination, ".git", "merge-request.json"))
				Expect(err).Should(BeNil())
				_, err = os.Stat(filepath.Join(destination, ".git", "resource", "changed_files"))
				Expect(err).Should(BeNil())
				sb, _ := os.ReadFile(filepath.Join(destination, ".git", "merge-request-source-branch"))
				Expect(string(sb)).Should(Equal("source-branch"))
			})
			It("Should write file changes", func() {
				project, _ := url.Parse("namespace/project.git")
				uri := root.ResolveReference(project)

				request := in.Request{
					Source: models.Source{
						URI:          uri.String(),
						PrivateToken: "$",
					},
					Version: models.Version{ID: "git://gitlab/MergeRequest/1"},
				}

				_, err := command.Run(destination, request)
				Expect(err).Should(BeNil())
				_, err = os.Stat(filepath.Join(destination, ".git", "resource", "changed_files"))
				Expect(err).Should(BeNil())
				cf, _ := os.ReadFile(filepath.Join(destination, ".git", "resource", "changed_files"))
				Expect(string(cf)).Should(Equal("/foo\n/bar"))
			})
		})

	})

})

func newMockRunner(destination string) in.GitRunner {
	os.MkdirAll(filepath.Join(destination, ".git"), 0755)
	return mockRunner{destination}
}

type mockRunner struct {
	destination string
}

func (mock mockRunner) Run(args ...string) error {
	fmt.Printf("mock: git %s\n", strings.Join(args, " "))
	return nil
}

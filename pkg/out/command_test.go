package out_test

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/Khan/genqlient/graphql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/simspace/gitlab-merge-request-resource/pkg/gitlab"
	"github.com/simspace/gitlab-merge-request-resource/pkg/models"
	"github.com/simspace/gitlab-merge-request-resource/pkg/out"
	gitlabv4 "github.com/xanzy/go-gitlab"
)

var _ = Describe("Out", func() {
	var (
		mux         *http.ServeMux
		server      *httptest.Server
		root        *url.URL
		command     *out.Command
		destination string
	)

	BeforeEach(func() {

		mux = http.NewServeMux()
		server = httptest.NewServer(mux)
		root, _ = url.Parse(server.URL)
		context, _ := url.Parse("/api/graphql")
		base := root.ResolveReference(context)
		client := graphql.NewClient(base.String(), nil)
		contextv4, _ := url.Parse("/api/v4")
		basev4 := root.ResolveReference(contextv4)
		clientv4, _ := gitlabv4.NewClient("$", gitlabv4.WithBaseURL(basev4.String()))
		destination, _ = os.MkdirTemp("", "gitlab-merge-request-resource-out")

		_ = os.Setenv("ATC_EXTERNAL_URL", "https://concourse-ci.company.ltd")
		_ = os.Setenv("BUILD_TEAM_NAME", "winner")
		_ = os.Setenv("BUILD_PIPELINE_NAME", "baltic")
		_ = os.Setenv("BUILD_JOB_NAME", "release")
		_ = os.Setenv("BUILD_NAME", "1")

		command = out.NewCommand(&client, clientv4)
	})

	AfterEach(func() {
		os.Remove(destination)
		server.Close()
	})

	Describe("Only update status", func() {

		BeforeEach(func() {
			mr := gitlab.MergeRequest{
				Id:              "gid://gitlab/MergeRequest/1",
				Iid:             "42",
				DiffHeadSha:     "abc",
				SourceProjectId: 1,
				Author:          gitlab.MergeRequestAuthor{Name: "john"},
				Commits: gitlab.MergeRequestCommitsCommitConnection{
					Nodes: []gitlab.Commit{
						{
							Sha:          "abc",
							AuthoredDate: time.Time{},
						},
					},
				},
			}
			content, _ := json.Marshal(mr)

			_ = os.Mkdir(path.Join(destination, "repo"), 0755)
			_ = os.Mkdir(path.Join(destination, "repo", ".git"), 0755)
			_ = os.WriteFile(path.Join(destination, "repo", ".git", "merge-request.json"), content, 0644)
		})

		It("Sets the commit status", func() {
			project, _ := url.Parse("namespace/project.git")
			uri := root.ResolveReference(project)

			request := out.Request{
				Source: models.Source{URI: uri.String()},
				Params: out.Params{
					Repository: "repo",
					Status:     "running",
				},
			}

			mux.HandleFunc("/api/v4/projects/1/statuses/abc", func(w http.ResponseWriter, r *http.Request) {
				body, _ := io.ReadAll(r.Body)
				Expect(string(body)).To(ContainSubstring(`"state":"running"`))
				status := gitlabv4.CommitStatus{ID: 1, SHA: "abc"}
				output, _ := json.Marshal(status)
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write(output)
			})

			response, err := command.Run(destination, request)
			Expect(err).Should(BeNil())
			Expect(response.Version.IID).To(Equal("42"))
		})

	})

	Describe("Only update labels", func() {

		BeforeEach(func() {
			mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
				lresp := gitlab.ListLabelsResponse{
					Project: gitlab.ListLabelsProject{
						Labels: gitlab.ListLabelsProjectLabelsLabelConnection{
							Nodes: []gitlab.Label{
								{
									Id:    "1",
									Title: "test-label",
								},
							},
						},
					},
				}
				sresp := gitlab.SetMergeRequestLabelsResponse{
					MergeRequestSetLabels: gitlab.SetMergeRequestLabelsMergeRequestSetLabelsMergeRequestSetLabelsPayload{
						MergeRequest: gitlab.MergeRequest{
							Id:          "gid//gitlab/MergeRequest/1",
							Iid:         "42",
							DiffHeadSha: "abc",
							Author:      gitlab.MergeRequestAuthor{Name: "john"},
							Commits: gitlab.MergeRequestCommitsCommitConnection{
								Nodes: []gitlab.Commit{
									{
										Sha:          "abc",
										AuthoredDate: time.Time{},
									},
								},
							},
							Labels: gitlab.MergeRequestLabelsLabelConnection{
								Nodes: []gitlab.Label{
									{
										Title: "test-label",
									},
									{
										Title: "existing-label",
									},
								},
							},
						},
					},
				}

				body, _ := ioutil.ReadAll(r.Body)
				graphqlReq := graphql.Request{}
				_ = json.Unmarshal(body, &graphqlReq)

				var data interface{}

				if graphqlReq.OpName == "ListLabels" {
					data = lresp
				} else {
					data = sresp
				}
				output, _ := json.Marshal(graphql.Response{Data: &data})
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(output)
			})

			mr := gitlab.MergeRequest{
				Id:          "gid://MergeRequest/1",
				Iid:         "42",
				DiffHeadSha: "abc",
				SourceProject: gitlab.MergeRequestSourceProject{
					HttpUrlToRepo: "https://gitlab.com/namespace/project.git",
				},
				SourceProjectId: 1,
				Labels: gitlab.MergeRequestLabelsLabelConnection{
					Nodes: []gitlab.Label{
						{
							Title: "existing-label",
						},
					},
				},
				Author: gitlab.MergeRequestAuthor{Name: "john"},
				Commits: gitlab.MergeRequestCommitsCommitConnection{
					Nodes: []gitlab.Commit{
						{
							Sha:          "abc",
							AuthoredDate: time.Time{},
						},
					},
				},
			}
			content, _ := json.Marshal(mr)

			_ = os.Mkdir(path.Join(destination, "repo"), 0755)
			_ = os.Mkdir(path.Join(destination, "repo", ".git"), 0755)
			_ = os.WriteFile(path.Join(destination, "repo", ".git", "merge-request.json"), content, 0644)
		})

		It("Updates the labels", func() {
			project, _ := url.Parse("namespace/project.git")
			uri := root.ResolveReference(project)

			request := out.Request{
				Source: models.Source{URI: uri.String()},
				Params: out.Params{
					Repository: "repo",
					Labels:     []string{"test-label"},
				},
			}

			response, err := command.Run(destination, request)
			Expect(err).Should(BeNil())
			Expect(response.Metadata[8].Value).To(ContainSubstring("test-label"))
		})

	})

	Describe("Only add comment", func() {

		BeforeEach(func() {
			mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
				resp := gitlab.CreateNoteResponse{}
				output, _ := json.Marshal(graphql.Response{Data: &resp})
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(output)
			})

			mr := gitlab.MergeRequest{
				Id:              "gid://MergeRequest/1",
				Iid:             "42",
				DiffHeadSha:     "abc",
				SourceProjectId: 1,
				Labels: gitlab.MergeRequestLabelsLabelConnection{
					Nodes: []gitlab.Label{},
				},
				Author: gitlab.MergeRequestAuthor{Name: "john"},
				Commits: gitlab.MergeRequestCommitsCommitConnection{
					Nodes: []gitlab.Commit{
						{
							Sha:          "abc",
							AuthoredDate: time.Time{},
						},
					},
				},
			}
			content, _ := json.Marshal(mr)

			_ = os.Mkdir(path.Join(destination, "repo"), 0755)
			_ = os.Mkdir(path.Join(destination, "repo", ".git"), 0755)
			_ = os.WriteFile(path.Join(destination, "repo", ".git", "merge-request.json"), content, 0644)
			_ = os.WriteFile(path.Join(destination, "comment.txt"), []byte("lorem ipsum"), 0644)
		})

		It("Does not throw an error", func() {

			project, _ := url.Parse("namespace/project.git")
			uri := root.ResolveReference(project)

			request := out.Request{
				Source: models.Source{URI: uri.String()},
				Params: out.Params{
					Repository: "repo",
					Comment:    out.Comment{FilePath: "comment.txt", Text: "new comment, $FILE_CONTENT"},
				},
			}

			_, err := command.Run(destination, request)
			Expect(err).Should(BeNil())
		})

	})

})

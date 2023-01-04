package check_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"time"

	"github.com/Khan/genqlient/graphql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/simspace/gitlab-merge-request-resource/pkg/check"
	"github.com/simspace/gitlab-merge-request-resource/pkg/gitlab"
	"github.com/simspace/gitlab-merge-request-resource/pkg/models"
	gitlabv4 "github.com/xanzy/go-gitlab"
)

var _ = Describe("Check", func() {

	var (
		t       time.Time
		updTime time.Time
		mux     *http.ServeMux
		command *check.Command
		root    *url.URL
	)

	BeforeEach(func() {
		t, _ = time.Parse(time.RFC3339, "2023-01-01T08:00:00Z")
		mux = http.NewServeMux()
		server := httptest.NewServer(mux)
		root, _ = url.Parse(server.URL)
		context, _ := url.Parse("/api/graphql")
		base := root.ResolveReference(context)
		client := graphql.NewClient(base.String(), nil)
		contextv4, _ := url.Parse("/api/v4")
		basev4 := root.ResolveReference(contextv4)
		clientv4, _ := gitlabv4.NewClient("$", gitlabv4.WithBaseURL(basev4.String()))

		_ = os.Setenv("ATC_EXTERNAL_URL", "https://concourse-ci.company.ltd")
		_ = os.Setenv("BUILD_TEAM_NAME", "winner")
		_ = os.Setenv("BUILD_PIPELINE_NAME", "baltic")
		_ = os.Setenv("BUILD_JOB_NAME", "release")
		_ = os.Setenv("BUILD_NAME", "1")

		command = check.NewCommand(&client, clientv4)
	})

	Describe("Run", func() {

		Context("When it has a minimal valid configuration", func() {

			BeforeEach(func() {
				mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
					projResp := gitlab.GetProjectResponse{
						Project: gitlab.Project{
							Id: "gid://gitlab/Project/1",
							MergeRequests: gitlab.ProjectMergeRequestsMergeRequestConnection{
								Nodes: []gitlab.MergeRequest{
									{
										Id:          "git://gitlab/MergeRequest/1234",
										Iid:         "88",
										Title:       "abc",
										DiffHeadSha: "abc",
										DiffStats:   []gitlab.MergeRequestDiffStats{},
										Commits: gitlab.MergeRequestCommitsCommitConnection{
											Nodes: []gitlab.Commit{
												{
													AuthoredDate: t,
													Sha:          "abc",
												},
											},
										},
										Labels: gitlab.MergeRequestLabelsLabelConnection{
											Nodes: []gitlab.Label{},
										},
									},
								},
							},
						},
					}
					output, _ := json.Marshal(graphql.Response{Data: projResp})
					w.Header().Set("content-type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write(output)
				})

				mux.HandleFunc("/api/v4/projects/42/repository/commits/abc/statuses", func(w http.ResponseWriter, r *http.Request) {
					statuses := []gitlabv4.CommitStatus{}
					output, _ := json.Marshal(statuses)
					w.Header().Set("content-type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write(output)
				})
			})

			It("Should return a single version", func() {

				project, _ := url.Parse("namespace/project.git")
				uri := root.ResolveReference(project)

				request := check.Request{
					Source: models.Source{
						URI:          uri.String(),
						PrivateToken: "$",
					},
				}

				response, err := command.Run(request)
				Expect(err).Should(BeNil())
				Expect(len(response)).To(Equal(1))
				Expect(response[0].ID).To(Equal(88))
				Expect(response[0].UpdatedAt).To(Equal(&t))
			})
		})

		Context("When it contains an invalid project uri", func() {

			BeforeEach(func() {
				mux.HandleFunc("/api/v4/projects/namespace/project/merge_requests", http.NotFound)
			})

			It("Should error when project uri is invalid", func() {

				project, _ := url.Parse("invalid/project.git")
				uri := root.ResolveReference(project)

				request := check.Request{
					Source: models.Source{
						URI:          uri.String(),
						PrivateToken: "$",
					},
				}

				response, err := command.Run(request)
				Expect(response).To(Equal(check.Response{}))
				Expect(err).NotTo(BeNil())
			})
		})

		Context("When it has a manual trigger comments", func() {

			BeforeEach(func() {
				updTime, _ = time.Parse(time.RFC3339, "2023-01-01T08:30:00Z")
				mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
					projResp := gitlab.GetProjectResponse{
						Project: gitlab.Project{
							Id: "gid://gitlab/Project/1",
							MergeRequests: gitlab.ProjectMergeRequestsMergeRequestConnection{
								Nodes: []gitlab.MergeRequest{
									{
										Id:          "git://gitlab/MergeRequest/1234",
										Iid:         "88",
										Title:       "abc",
										DiffHeadSha: "abc",
										DiffStats:   []gitlab.MergeRequestDiffStats{},
										Commits: gitlab.MergeRequestCommitsCommitConnection{
											Nodes: []gitlab.Commit{
												{
													AuthoredDate: t,
													Sha:          "abc",
													Pipelines: gitlab.CommitPipelinesPipelineConnection{
														Nodes: []gitlab.Pipeline{
															{
																Sha:    "abc",
																Status: gitlab.PipelineStatusEnumFailed,
															},
														},
													},
												},
											},
										},
										Notes: gitlab.MergeRequestNotesNoteConnection{
											Nodes: []gitlab.Note{
												{
													Body:      "[trigger ci]",
													UpdatedAt: updTime,
												},
											},
										},
									},
								},
							},
						},
					}
					output, _ := json.Marshal(graphql.Response{Data: projResp})
					w.Header().Set("content-type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write(output)
				})
			})

			It("Should have an updated timestamp", func() {

				project, _ := url.Parse("namespace/project.git")
				uri := root.ResolveReference(project)

				request := check.Request{
					Source: models.Source{
						URI:          uri.String(),
						PrivateToken: "$",
					},
				}

				response, err := command.Run(request)
				Expect(err).Should(BeNil())
				Expect(response[0].UpdatedAt).To(Equal(&updTime))
			})
		})

		Context("When it has a label specified", func() {

			BeforeEach(func() {
				updTime, _ = time.Parse(time.RFC3339, "2023-01-01T08:30:00Z")
				mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
					projResp := gitlab.GetProjectResponse{
						Project: gitlab.Project{
							Id: "gid://gitlab/Project/1",
							MergeRequests: gitlab.ProjectMergeRequestsMergeRequestConnection{
								Nodes: []gitlab.MergeRequest{
									{
										Id:          "git://gitlab/MergeRequest/1234",
										Iid:         "88",
										Title:       "abc",
										DiffHeadSha: "abc",
										DiffStats:   []gitlab.MergeRequestDiffStats{},
										Commits: gitlab.MergeRequestCommitsCommitConnection{
											Nodes: []gitlab.Commit{
												{
													AuthoredDate: t,
													Sha:          "abc",
												},
											},
										},
									},
									{
										Id:          "git://gitlab/MergeRequest/1235",
										Iid:         "89",
										Title:       "abc label",
										DiffHeadSha: "abc",
										DiffStats:   []gitlab.MergeRequestDiffStats{},
										Commits: gitlab.MergeRequestCommitsCommitConnection{
											Nodes: []gitlab.Commit{
												{
													AuthoredDate: t,
													Sha:          "abc",
												},
											},
										},
										Labels: gitlab.MergeRequestLabelsLabelConnection{
											Nodes: []gitlab.Label{
												{
													Title: "labelled",
												},
											},
										},
									},
								},
							},
						},
					}
					output, _ := json.Marshal(graphql.Response{Data: projResp})
					w.Header().Set("content-type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write(output)
				})

				mux.HandleFunc("/api/v4/projects/42/repository/commits/abc/statuses", func(w http.ResponseWriter, r *http.Request) {
					statuses := []gitlabv4.CommitStatus{}
					output, _ := json.Marshal(statuses)
					w.Header().Set("content-type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write(output)
				})
			})

			It("Should only return labelled merge requests", func() {

				project, _ := url.Parse("namespace/project.git")
				uri := root.ResolveReference(project)

				request := check.Request{
					Source: models.Source{
						URI:          uri.String(),
						PrivateToken: "$",
						Labels: []string{
							"labelled",
						},
					},
				}

				response, err := command.Run(request)
				Expect(err).Should(BeNil())
				Expect(len(response)).To(Equal(1))
				Expect(response[0].ID).To(Equal(89))
			})
		})
	})
})

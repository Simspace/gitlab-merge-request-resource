package in_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/simspace/gitlab-merge-request-resource/pkg/in"
	"github.com/simspace/gitlab-merge-request-resource/pkg/models"
	"github.com/xanzy/go-gitlab"
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
		context, _ := url.Parse("/api/v4")
		base := root.ResolveReference(context)
		client, _ := gitlab.NewClient("$", gitlab.WithBaseURL(base.String()))
		command = in.NewCommand(client).WithRunner(newMockRunner(destination))

	})

	AfterEach(func() {
		defer os.Remove(destination)
		server.Close()
	})

	Describe("Run", func() {

		BeforeEach(func() {
			mux.HandleFunc("/api/v4/projects/namespace/project/merge_requests/1/changes", func(w http.ResponseWriter, r *http.Request) {
				mr := gitlab.MergeRequest{
					IID:             88,
					ID:              99,
					SHA:             "abc",
					ProjectID:       42,
					TargetProjectID: 42,
					SourceProjectID: 42,
					SourceBranch:    "source-branch",
					TargetBranch:    "target-branch",
					Author:          &gitlab.BasicUser{Name: "Tester"},
					Changes: changes{
						{
							OldPath:     "/foo",
							NewPath:     "/foo",
							AMode:       "",
							BMode:       "",
							Diff:        "",
							NewFile:     false,
							RenamedFile: false,
							DeletedFile: false,
						},
						{
							OldPath:     "/bar",
							NewPath:     "/bar",
							AMode:       "",
							BMode:       "",
							Diff:        "",
							NewFile:     false,
							RenamedFile: false,
							DeletedFile: false,
						},
					},
				}
				output, _ := json.Marshal(mr)
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(output)
			})
			mux.HandleFunc("/api/v4/projects/42", func(w http.ResponseWriter, r *http.Request) {
				project, _ := url.Parse("namespace/project.git")
				uri := root.ResolveReference(project)
				mr := gitlab.Project{HTTPURLToRepo: uri.String()}
				output, _ := json.Marshal(mr)
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(output)
			})
			mux.HandleFunc("/api/v4/projects/42/repository/commits/abc", func(w http.ResponseWriter, r *http.Request) {
				commit := gitlab.Commit{CommittedDate: &t}
				output, _ := json.Marshal(commit)
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(output)
			})
			mux.HandleFunc("/api/v4/user", func(w http.ResponseWriter, r *http.Request) {
				user := gitlab.User{
					Username: "test",
					Email:    "test@example.com",
				}
				output, _ := json.Marshal(user)
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
					Version: models.Version{ID: 1},
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
					Version: models.Version{ID: 1},
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

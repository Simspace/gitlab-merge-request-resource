package gitlab_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/simspace/gitlab-merge-request-resource/pkg/gitlab"
)

var _ = Describe("Gitlab", func() {
	var createResp gitlab.CreateNoteResponse
	var expErr error

	Describe("WrapErrors", func() {
		Context("when an error is returned", func() {
			BeforeEach(func() {
				createResp = gitlab.CreateNoteResponse{
					CreateNote: gitlab.CreateNoteCreateNoteCreateNotePayload{
						Errors: []string{"error1"},
					},
				}
				expErr = errors.New("error1")
			})

			It("should return the wrapped error", func() {
				_, err := gitlab.WrapErrors(createResp, nil)
				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError(expErr))
			})
		})
		Context("when multiple errors are returned", func() {
			BeforeEach(func() {
				createResp = gitlab.CreateNoteResponse{
					CreateNote: gitlab.CreateNoteCreateNoteCreateNotePayload{
						Errors: []string{"error1", "error2"},
					},
				}
				expErr = errors.New("multiple API errors: error1,error2")
			})

			It("should return all wrapped errors", func() {
				_, err := gitlab.WrapErrors(createResp, nil)
				Expect(err).NotTo(BeNil())
				Expect(err).To(MatchError(expErr))
			})
		})
	})
})

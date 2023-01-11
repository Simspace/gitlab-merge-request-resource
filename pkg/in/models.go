package in

import (
	"github.com/simspace/gitlab-merge-request-resource/pkg/models"
)

type Request struct {
	Source  models.Source  `json:"source"`
	Version models.Version `json:"version"`
}

type Response struct {
	Version  models.Version  `json:"version"`
	Metadata models.Metadata `json:"metadata"`
}

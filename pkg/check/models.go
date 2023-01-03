package check

import (
	"github.com/simspace/gitlab-merge-request-resource/pkg/models"
)

type Request struct {
	Source  models.Source  `json:"source"`
	Version models.Version `json:"version,omitempty"`
}

type Response []models.Version

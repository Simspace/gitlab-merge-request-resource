package check

import (
	. "github.com/simspace/gitlab-merge-request-resource/pkg"
)

type Request struct {
	Source  Source  `json:"source"`
	Version Version `json:"version,omitempty"`
}

type Response []Version

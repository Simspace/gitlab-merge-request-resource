package pkg

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Raw GraphQL queries here are auto-generated into GoLang functions,
// returning typed values that corresponde to the GraphQL types.
const (
	_ = `# @genqlient
fragment Pipeline on Pipeline {
  sha
  status
}

fragment Commit on Commit {
  id
  authoredDate
  sha
  title
  pipelines {
	# @genqlient(flatten: true)
    nodes {
      ...Pipeline
    }
  }
}

fragment Label on Label {
  title
}

fragment MergeRequest on MergeRequest {
  id
  iid
  title
  diffHeadSha
	diffStats {
	  path
	}
  commits {
	# @genqlient(flatten: true)
    nodes {
      ...Commit
    }
  }
  labels {
	# @genqlient(flatten: true)
    nodes {
	  ...Label
	}
  }
}

query GetProject($project: ID!, $state: MergeRequestState!) {
  project(fullPath: $project) {
    id
    mergeRequests(state: $state, sort: UPDATED_ASC) {
	  # @genqlient(flatten: true)
      nodes {
        ...MergeRequest
      }
    }
  }
}`
)

func Fatal(doing string, err error) {
	fmt.Fprintf(os.Stderr, "error %s: %s\n", doing, err)
	os.Exit(1)
}

func GetDefaultClient(insecure bool) *http.Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: insecure}
	return http.DefaultClient
}

func matchPath(patterns []string, path string) bool {
	for _, pattern := range patterns {
		ok, _ := filepath.Match(pattern, path)
		if ok {
			return true
		}
		if isInsidePath(pattern, path) {
			return true
		}
	}
	return false
}

func isInsidePath(parent, child string) bool {
	if parent == child {
		return true
	}

	// we add a trailing slash so that we only get prefix matches on a
	// directory separator
	parentWithTrailingSlash := parent
	if !strings.HasSuffix(parentWithTrailingSlash, string(filepath.Separator)) {
		parentWithTrailingSlash += string(filepath.Separator)
	}

	return strings.HasPrefix(child, parentWithTrailingSlash)
}

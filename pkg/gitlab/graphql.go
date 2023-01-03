package gitlab

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

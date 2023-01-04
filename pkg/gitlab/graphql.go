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
  message
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

fragment Note on Note {
  body
  updatedAt
}

fragment MergeRequest on MergeRequest {
  id
  iid
  title
  diffHeadSha
  draft
  mergeStatusEnum
  sourceBranch
  targetBranch
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
  notes {
    # @genqlient(flatten: true)
    nodes {
      ...Note
    }
  }
}

fragment Project on Project {
  id
  mergeRequests(state: $state, sort: UPDATED_ASC) {
    # @genqlient(flatten: true)
    nodes {
      ...MergeRequest
    }
  }
}

query GetProject($project: ID!, $state: MergeRequestState!) {
  # @genqlient(flatten: true)
  project(fullPath: $project) {
    ...Project
  }
}`
)

package gitlab

// Raw GraphQL queries here are auto-generated into GoLang functions,
// returning typed values that corresponde to the GraphQL types.
const (
	_ = `# @genqlient
fragment Pipeline on Pipeline {
  sha
  status
}`

	_ = `# @genqlient
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
}`

	_ = `# @genqlient
fragment Label on Label {
  title
}`

	_ = `# @genqlient
fragment Note on Note {
  body
  updatedAt
}`

	_ = `# @genqlient
fragment MergeRequest on MergeRequest {
  id
  iid
  title
  author {
    name
  }
  diffHeadSha
  diffStats {
    path
  }
  draft
  mergeStatusEnum
  sourceBranch
  sourceProjectId
  sourceProject {
    httpUrlToRepo
  }
  targetBranch
  targetProject {
    httpUrlToRepo
  }
  webUrl
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
 }`

	_ = `# @genqlient
fragment Project on Project {
  id
  mergeRequests(state: $state, sort: UPDATED_ASC) {
    # @genqlient(flatten: true)
    nodes {
      ...MergeRequest
    }
  }
}`

	_ = `# @genqlient
query ListMergeRequests($project: ID!, $state: MergeRequestState!) {
  # @genqlient(flatten: true)
  project(fullPath: $project) {
    ...Project
  }
}`

	_ = `# @genqlient
query GetCurrentUser {
  # @genqlient(typename: User)
  currentUser{
    name
    publicEmail
  }
}`

	_ = `# @genqlient
query GetMergeRequest($id: MergeRequestID!) {
  # @genqlient(flatten: true)
  mergeRequest(id: $id) {
    ...MergeRequest
  }
}`
)

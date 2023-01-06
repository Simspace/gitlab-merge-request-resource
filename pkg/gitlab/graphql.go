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
  id
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
query ListLabels($project: ID!) {
  project(fullPath: $project) {
    labels {
      # @genqlient(flatten: true)
      nodes {
        ...Label
      }
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

	_ = `# @genqlient
# @genqlient(for: "CreateNoteInput.confidential", omitempty: true)
# @genqlient(for: "CreateNoteInput.clientMutationId", omitempty: true)
# @genqlient(for: "CreateNoteInput.discussionId", omitempty: true)
# @genqlient(for: "CreateNoteInput.internal", omitempty: true)
mutation CreateNote(
  $input: CreateNoteInput!
) {
  createNote(input: $input) {
    errors
  }
}`

	_ = `# @genqlient
# @genqlient(for: "MergeRequestSetLabelsInput.clientMutationId", omitempty: true)
# @genqlient(for: "MergeRequestSetLabelsInput.operationMode", omitempty: true)
mutation SetMergeRequestLabels(
  $input: MergeRequestSetLabelsInput!
) {
  mergeRequestSetLabels(input: $input) {
    errors
    # @genqlient(flatten: true)
    mergeRequest {
      ...MergeRequest
    }
  }
}`
)

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package gitlab

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// Commit includes the GraphQL fields of Commit requested by the fragment Commit.
type Commit struct {
	// ID (global ID) of the commit.
	Id string `json:"id"`
	// Timestamp of when the commit was authored.
	AuthoredDate time.Time `json:"authoredDate"`
	// Raw commit message.
	Message string `json:"message"`
	// SHA1 ID of the commit.
	Sha string `json:"sha"`
	// Title of the commit message.
	Title string `json:"title"`
	// Pipelines of the commit ordered latest first.
	Pipelines CommitPipelinesPipelineConnection `json:"pipelines"`
}

// GetId returns Commit.Id, and is useful for accessing the field via an interface.
func (v *Commit) GetId() string { return v.Id }

// GetAuthoredDate returns Commit.AuthoredDate, and is useful for accessing the field via an interface.
func (v *Commit) GetAuthoredDate() time.Time { return v.AuthoredDate }

// GetMessage returns Commit.Message, and is useful for accessing the field via an interface.
func (v *Commit) GetMessage() string { return v.Message }

// GetSha returns Commit.Sha, and is useful for accessing the field via an interface.
func (v *Commit) GetSha() string { return v.Sha }

// GetTitle returns Commit.Title, and is useful for accessing the field via an interface.
func (v *Commit) GetTitle() string { return v.Title }

// GetPipelines returns Commit.Pipelines, and is useful for accessing the field via an interface.
func (v *Commit) GetPipelines() CommitPipelinesPipelineConnection { return v.Pipelines }

// CommitPipelinesPipelineConnection includes the requested fields of the GraphQL type PipelineConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Pipeline.
type CommitPipelinesPipelineConnection struct {
	// A list of nodes.
	Nodes []Pipeline `json:"nodes"`
}

// GetNodes returns CommitPipelinesPipelineConnection.Nodes, and is useful for accessing the field via an interface.
func (v *CommitPipelinesPipelineConnection) GetNodes() []Pipeline { return v.Nodes }

// CreateNoteCreateNoteCreateNotePayload includes the requested fields of the GraphQL type CreateNotePayload.
// The GraphQL type's documentation follows.
//
// Autogenerated return type of CreateNote
type CreateNoteCreateNoteCreateNotePayload struct {
	// Errors encountered during execution of the mutation.
	Errors []string `json:"errors"`
}

// GetErrors returns CreateNoteCreateNoteCreateNotePayload.Errors, and is useful for accessing the field via an interface.
func (v *CreateNoteCreateNoteCreateNotePayload) GetErrors() []string { return v.Errors }

// Autogenerated input type of CreateNote
type CreateNoteInput struct {
	// Global ID of the resource to add a note to.
	NoteableId string `json:"noteableId"`
	// Content of the note.
	Body string `json:"body"`
	// Internal flag for a note. Default is false.
	Internal bool `json:"internal,omitempty"`
	// A unique identifier for the client performing the mutation.
	ClientMutationId string `json:"clientMutationId,omitempty"`
	// Global ID of the discussion the note is in reply to.
	DiscussionId string `json:"discussionId,omitempty"`
	// SHA of the head commit which is used to ensure that the merge request has not been updated since the request was sent.
	MergeRequestDiffHeadSha string `json:"mergeRequestDiffHeadSha"`
}

// GetNoteableId returns CreateNoteInput.NoteableId, and is useful for accessing the field via an interface.
func (v *CreateNoteInput) GetNoteableId() string { return v.NoteableId }

// GetBody returns CreateNoteInput.Body, and is useful for accessing the field via an interface.
func (v *CreateNoteInput) GetBody() string { return v.Body }

// GetInternal returns CreateNoteInput.Internal, and is useful for accessing the field via an interface.
func (v *CreateNoteInput) GetInternal() bool { return v.Internal }

// GetClientMutationId returns CreateNoteInput.ClientMutationId, and is useful for accessing the field via an interface.
func (v *CreateNoteInput) GetClientMutationId() string { return v.ClientMutationId }

// GetDiscussionId returns CreateNoteInput.DiscussionId, and is useful for accessing the field via an interface.
func (v *CreateNoteInput) GetDiscussionId() string { return v.DiscussionId }

// GetMergeRequestDiffHeadSha returns CreateNoteInput.MergeRequestDiffHeadSha, and is useful for accessing the field via an interface.
func (v *CreateNoteInput) GetMergeRequestDiffHeadSha() string { return v.MergeRequestDiffHeadSha }

// CreateNoteResponse is returned by CreateNote on success.
type CreateNoteResponse struct {
	// Creates a Note.
	// If the body of the Note contains only quick actions,
	// the Note will be destroyed during an update, and no Note will be
	// returned.
	CreateNote CreateNoteCreateNoteCreateNotePayload `json:"createNote"`
}

// GetCreateNote returns CreateNoteResponse.CreateNote, and is useful for accessing the field via an interface.
func (v *CreateNoteResponse) GetCreateNote() CreateNoteCreateNoteCreateNotePayload {
	return v.CreateNote
}

// GetCurrentUserResponse is returned by GetCurrentUser on success.
type GetCurrentUserResponse struct {
	// Get information about current user.
	CurrentUser User `json:"currentUser"`
}

// GetCurrentUser returns GetCurrentUserResponse.CurrentUser, and is useful for accessing the field via an interface.
func (v *GetCurrentUserResponse) GetCurrentUser() User { return v.CurrentUser }

// GetMergeRequestResponse is returned by GetMergeRequest on success.
type GetMergeRequestResponse struct {
	// Find a merge request.
	MergeRequest MergeRequest `json:"mergeRequest"`
}

// GetMergeRequest returns GetMergeRequestResponse.MergeRequest, and is useful for accessing the field via an interface.
func (v *GetMergeRequestResponse) GetMergeRequest() MergeRequest { return v.MergeRequest }

// Label includes the GraphQL fields of Label requested by the fragment Label.
type Label struct {
	// Label ID.
	Id string `json:"id"`
	// Content of the label.
	Title string `json:"title"`
}

// GetId returns Label.Id, and is useful for accessing the field via an interface.
func (v *Label) GetId() string { return v.Id }

// GetTitle returns Label.Title, and is useful for accessing the field via an interface.
func (v *Label) GetTitle() string { return v.Title }

// ListLabelsProject includes the requested fields of the GraphQL type Project.
type ListLabelsProject struct {
	// Labels available on this project.
	Labels ListLabelsProjectLabelsLabelConnection `json:"labels"`
}

// GetLabels returns ListLabelsProject.Labels, and is useful for accessing the field via an interface.
func (v *ListLabelsProject) GetLabels() ListLabelsProjectLabelsLabelConnection { return v.Labels }

// ListLabelsProjectLabelsLabelConnection includes the requested fields of the GraphQL type LabelConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Label.
type ListLabelsProjectLabelsLabelConnection struct {
	// A list of nodes.
	Nodes []Label `json:"nodes"`
}

// GetNodes returns ListLabelsProjectLabelsLabelConnection.Nodes, and is useful for accessing the field via an interface.
func (v *ListLabelsProjectLabelsLabelConnection) GetNodes() []Label { return v.Nodes }

// ListLabelsResponse is returned by ListLabels on success.
type ListLabelsResponse struct {
	// Find a project.
	Project ListLabelsProject `json:"project"`
}

// GetProject returns ListLabelsResponse.Project, and is useful for accessing the field via an interface.
func (v *ListLabelsResponse) GetProject() ListLabelsProject { return v.Project }

// ListMergeRequestsResponse is returned by ListMergeRequests on success.
type ListMergeRequestsResponse struct {
	// Find a project.
	Project Project `json:"project"`
}

// GetProject returns ListMergeRequestsResponse.Project, and is useful for accessing the field via an interface.
func (v *ListMergeRequestsResponse) GetProject() Project { return v.Project }

// MergeRequest includes the GraphQL fields of MergeRequest requested by the fragment MergeRequest.
type MergeRequest struct {
	// ID of the merge request.
	Id string `json:"id"`
	// Internal ID of the merge request.
	Iid string `json:"iid"`
	// Title of the merge request.
	Title string `json:"title"`
	// User who created this merge request.
	Author MergeRequestAuthor `json:"author"`
	// Diff head SHA of the merge request.
	DiffHeadSha string `json:"diffHeadSha"`
	// Details about which files were changed in this merge request.
	DiffStats []MergeRequestDiffStats `json:"diffStats"`
	// Indicates if the merge request is a draft.
	Draft bool `json:"draft"`
	// Merge status of the merge request.
	MergeStatusEnum MergeStatus `json:"mergeStatusEnum"`
	// Source branch of the merge request.
	SourceBranch string `json:"sourceBranch"`
	// ID of the merge request source project.
	SourceProjectId int `json:"sourceProjectId"`
	// Source project of the merge request.
	SourceProject MergeRequestSourceProject `json:"sourceProject"`
	// Target branch of the merge request.
	TargetBranch string `json:"targetBranch"`
	// Target project of the merge request.
	TargetProject MergeRequestTargetProject `json:"targetProject"`
	// Web URL of the merge request.
	WebUrl string `json:"webUrl"`
	// Merge request commits.
	Commits MergeRequestCommitsCommitConnection `json:"commits"`
	// Labels of the merge request.
	Labels MergeRequestLabelsLabelConnection `json:"labels"`
	// All notes on this noteable.
	Notes MergeRequestNotesNoteConnection `json:"notes"`
}

// GetId returns MergeRequest.Id, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetId() string { return v.Id }

// GetIid returns MergeRequest.Iid, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetIid() string { return v.Iid }

// GetTitle returns MergeRequest.Title, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetTitle() string { return v.Title }

// GetAuthor returns MergeRequest.Author, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetAuthor() MergeRequestAuthor { return v.Author }

// GetDiffHeadSha returns MergeRequest.DiffHeadSha, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetDiffHeadSha() string { return v.DiffHeadSha }

// GetDiffStats returns MergeRequest.DiffStats, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetDiffStats() []MergeRequestDiffStats { return v.DiffStats }

// GetDraft returns MergeRequest.Draft, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetDraft() bool { return v.Draft }

// GetMergeStatusEnum returns MergeRequest.MergeStatusEnum, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetMergeStatusEnum() MergeStatus { return v.MergeStatusEnum }

// GetSourceBranch returns MergeRequest.SourceBranch, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetSourceBranch() string { return v.SourceBranch }

// GetSourceProjectId returns MergeRequest.SourceProjectId, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetSourceProjectId() int { return v.SourceProjectId }

// GetSourceProject returns MergeRequest.SourceProject, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetSourceProject() MergeRequestSourceProject { return v.SourceProject }

// GetTargetBranch returns MergeRequest.TargetBranch, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetTargetBranch() string { return v.TargetBranch }

// GetTargetProject returns MergeRequest.TargetProject, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetTargetProject() MergeRequestTargetProject { return v.TargetProject }

// GetWebUrl returns MergeRequest.WebUrl, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetWebUrl() string { return v.WebUrl }

// GetCommits returns MergeRequest.Commits, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetCommits() MergeRequestCommitsCommitConnection { return v.Commits }

// GetLabels returns MergeRequest.Labels, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetLabels() MergeRequestLabelsLabelConnection { return v.Labels }

// GetNotes returns MergeRequest.Notes, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetNotes() MergeRequestNotesNoteConnection { return v.Notes }

// MergeRequestAuthor includes the requested fields of the GraphQL type MergeRequestAuthor.
// The GraphQL type's documentation follows.
//
// The author of the merge request.
type MergeRequestAuthor struct {
	// Human-readable name of the user. Returns `****` if the user is a project bot
	// and the requester does not have permission to view the project.
	Name string `json:"name"`
}

// GetName returns MergeRequestAuthor.Name, and is useful for accessing the field via an interface.
func (v *MergeRequestAuthor) GetName() string { return v.Name }

// MergeRequestCommitsCommitConnection includes the requested fields of the GraphQL type CommitConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Commit.
type MergeRequestCommitsCommitConnection struct {
	// A list of nodes.
	Nodes []Commit `json:"nodes"`
}

// GetNodes returns MergeRequestCommitsCommitConnection.Nodes, and is useful for accessing the field via an interface.
func (v *MergeRequestCommitsCommitConnection) GetNodes() []Commit { return v.Nodes }

// MergeRequestDiffStats includes the requested fields of the GraphQL type DiffStats.
// The GraphQL type's documentation follows.
//
// Changes to a single file
type MergeRequestDiffStats struct {
	// File path, relative to repository root.
	Path string `json:"path"`
}

// GetPath returns MergeRequestDiffStats.Path, and is useful for accessing the field via an interface.
func (v *MergeRequestDiffStats) GetPath() string { return v.Path }

// MergeRequestLabelsLabelConnection includes the requested fields of the GraphQL type LabelConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Label.
type MergeRequestLabelsLabelConnection struct {
	// A list of nodes.
	Nodes []Label `json:"nodes"`
}

// GetNodes returns MergeRequestLabelsLabelConnection.Nodes, and is useful for accessing the field via an interface.
func (v *MergeRequestLabelsLabelConnection) GetNodes() []Label { return v.Nodes }

// MergeRequestNotesNoteConnection includes the requested fields of the GraphQL type NoteConnection.
// The GraphQL type's documentation follows.
//
// The connection type for Note.
type MergeRequestNotesNoteConnection struct {
	// A list of nodes.
	Nodes []Note `json:"nodes"`
}

// GetNodes returns MergeRequestNotesNoteConnection.Nodes, and is useful for accessing the field via an interface.
func (v *MergeRequestNotesNoteConnection) GetNodes() []Note { return v.Nodes }

// Autogenerated input type of MergeRequestSetLabels
type MergeRequestSetLabelsInput struct {
	// Project the merge request to mutate is in.
	ProjectPath string `json:"projectPath"`
	// IID of the merge request to mutate.
	Iid string `json:"iid"`
	// A unique identifier for the client performing the mutation.
	ClientMutationId string `json:"clientMutationId,omitempty"`
	// Label IDs to set. Replaces existing labels by default.
	LabelIds []string `json:"labelIds"`
	// Changes the operation mode. Defaults to REPLACE.
	OperationMode MutationOperationMode `json:"operationMode,omitempty"`
}

// GetProjectPath returns MergeRequestSetLabelsInput.ProjectPath, and is useful for accessing the field via an interface.
func (v *MergeRequestSetLabelsInput) GetProjectPath() string { return v.ProjectPath }

// GetIid returns MergeRequestSetLabelsInput.Iid, and is useful for accessing the field via an interface.
func (v *MergeRequestSetLabelsInput) GetIid() string { return v.Iid }

// GetClientMutationId returns MergeRequestSetLabelsInput.ClientMutationId, and is useful for accessing the field via an interface.
func (v *MergeRequestSetLabelsInput) GetClientMutationId() string { return v.ClientMutationId }

// GetLabelIds returns MergeRequestSetLabelsInput.LabelIds, and is useful for accessing the field via an interface.
func (v *MergeRequestSetLabelsInput) GetLabelIds() []string { return v.LabelIds }

// GetOperationMode returns MergeRequestSetLabelsInput.OperationMode, and is useful for accessing the field via an interface.
func (v *MergeRequestSetLabelsInput) GetOperationMode() MutationOperationMode { return v.OperationMode }

// MergeRequestSourceProject includes the requested fields of the GraphQL type Project.
type MergeRequestSourceProject struct {
	// URL to connect to the project via HTTPS.
	HttpUrlToRepo string `json:"httpUrlToRepo"`
}

// GetHttpUrlToRepo returns MergeRequestSourceProject.HttpUrlToRepo, and is useful for accessing the field via an interface.
func (v *MergeRequestSourceProject) GetHttpUrlToRepo() string { return v.HttpUrlToRepo }

// State of a GitLab merge request
type MergeRequestState string

const (
	// Merge request has been merged.
	MergeRequestStateMerged MergeRequestState = "merged"
	// In open state.
	MergeRequestStateOpened MergeRequestState = "opened"
	// In closed state.
	MergeRequestStateClosed MergeRequestState = "closed"
	// Discussion has been locked.
	MergeRequestStateLocked MergeRequestState = "locked"
	// All available.
	MergeRequestStateAll MergeRequestState = "all"
)

// MergeRequestTargetProject includes the requested fields of the GraphQL type Project.
type MergeRequestTargetProject struct {
	// URL to connect to the project via HTTPS.
	HttpUrlToRepo string `json:"httpUrlToRepo"`
}

// GetHttpUrlToRepo returns MergeRequestTargetProject.HttpUrlToRepo, and is useful for accessing the field via an interface.
func (v *MergeRequestTargetProject) GetHttpUrlToRepo() string { return v.HttpUrlToRepo }

// Representation of whether a GitLab merge request can be merged.
type MergeStatus string

const (
	// Merge status has not been checked.
	MergeStatusUnchecked MergeStatus = "UNCHECKED"
	// Currently checking for mergeability.
	MergeStatusChecking MergeStatus = "CHECKING"
	// There are no conflicts between the source and target branches.
	MergeStatusCanBeMerged MergeStatus = "CAN_BE_MERGED"
	// There are conflicts between the source and target branches.
	MergeStatusCannotBeMerged MergeStatus = "CANNOT_BE_MERGED"
	// Currently unchecked. The previous state was `CANNOT_BE_MERGED`.
	MergeStatusCannotBeMergedRecheck MergeStatus = "CANNOT_BE_MERGED_RECHECK"
)

// Different toggles for changing mutator behavior
type MutationOperationMode string

const (
	// Performs a replace operation.
	MutationOperationModeReplace MutationOperationMode = "REPLACE"
	// Performs an append operation.
	MutationOperationModeAppend MutationOperationMode = "APPEND"
	// Performs a removal operation.
	MutationOperationModeRemove MutationOperationMode = "REMOVE"
)

// Note includes the GraphQL fields of Note requested by the fragment Note.
type Note struct {
	// Content of the note.
	Body string `json:"body"`
	// Timestamp of the note's last activity.
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetBody returns Note.Body, and is useful for accessing the field via an interface.
func (v *Note) GetBody() string { return v.Body }

// GetUpdatedAt returns Note.UpdatedAt, and is useful for accessing the field via an interface.
func (v *Note) GetUpdatedAt() time.Time { return v.UpdatedAt }

// Pipeline includes the GraphQL fields of Pipeline requested by the fragment Pipeline.
type Pipeline struct {
	// SHA of the pipeline's commit.
	Sha string `json:"sha"`
	// Status of the pipeline (CREATED, WAITING_FOR_RESOURCE, PREPARING, PENDING,
	// RUNNING, FAILED, SUCCESS, CANCELED, SKIPPED, MANUAL, SCHEDULED)
	Status PipelineStatusEnum `json:"status"`
}

// GetSha returns Pipeline.Sha, and is useful for accessing the field via an interface.
func (v *Pipeline) GetSha() string { return v.Sha }

// GetStatus returns Pipeline.Status, and is useful for accessing the field via an interface.
func (v *Pipeline) GetStatus() PipelineStatusEnum { return v.Status }

type PipelineStatusEnum string

const (
	// Pipeline has been created.
	PipelineStatusEnumCreated PipelineStatusEnum = "CREATED"
	// A resource (for example, a runner) that the pipeline requires to run is unavailable.
	PipelineStatusEnumWaitingForResource PipelineStatusEnum = "WAITING_FOR_RESOURCE"
	// Pipeline is preparing to run.
	PipelineStatusEnumPreparing PipelineStatusEnum = "PREPARING"
	// Pipeline has not started running yet.
	PipelineStatusEnumPending PipelineStatusEnum = "PENDING"
	// Pipeline is running.
	PipelineStatusEnumRunning PipelineStatusEnum = "RUNNING"
	// At least one stage of the pipeline failed.
	PipelineStatusEnumFailed PipelineStatusEnum = "FAILED"
	// Pipeline completed successfully.
	PipelineStatusEnumSuccess PipelineStatusEnum = "SUCCESS"
	// Pipeline was canceled before completion.
	PipelineStatusEnumCanceled PipelineStatusEnum = "CANCELED"
	// Pipeline was skipped.
	PipelineStatusEnumSkipped PipelineStatusEnum = "SKIPPED"
	// Pipeline needs to be manually started.
	PipelineStatusEnumManual PipelineStatusEnum = "MANUAL"
	// Pipeline is scheduled to run.
	PipelineStatusEnumScheduled PipelineStatusEnum = "SCHEDULED"
)

// Project includes the GraphQL fields of Project requested by the fragment Project.
type Project struct {
	// ID of the project.
	Id string `json:"id"`
	// Merge requests of the project.
	MergeRequests ProjectMergeRequestsMergeRequestConnection `json:"mergeRequests"`
}

// GetId returns Project.Id, and is useful for accessing the field via an interface.
func (v *Project) GetId() string { return v.Id }

// GetMergeRequests returns Project.MergeRequests, and is useful for accessing the field via an interface.
func (v *Project) GetMergeRequests() ProjectMergeRequestsMergeRequestConnection {
	return v.MergeRequests
}

// ProjectMergeRequestsMergeRequestConnection includes the requested fields of the GraphQL type MergeRequestConnection.
// The GraphQL type's documentation follows.
//
// The connection type for MergeRequest.
type ProjectMergeRequestsMergeRequestConnection struct {
	// A list of nodes.
	Nodes []MergeRequest `json:"nodes"`
}

// GetNodes returns ProjectMergeRequestsMergeRequestConnection.Nodes, and is useful for accessing the field via an interface.
func (v *ProjectMergeRequestsMergeRequestConnection) GetNodes() []MergeRequest { return v.Nodes }

// SetMergeRequestLabelsMergeRequestSetLabelsMergeRequestSetLabelsPayload includes the requested fields of the GraphQL type MergeRequestSetLabelsPayload.
// The GraphQL type's documentation follows.
//
// Autogenerated return type of MergeRequestSetLabels
type SetMergeRequestLabelsMergeRequestSetLabelsMergeRequestSetLabelsPayload struct {
	// Errors encountered during execution of the mutation.
	Errors []string `json:"errors"`
	// Merge request after mutation.
	MergeRequest MergeRequest `json:"mergeRequest"`
}

// GetErrors returns SetMergeRequestLabelsMergeRequestSetLabelsMergeRequestSetLabelsPayload.Errors, and is useful for accessing the field via an interface.
func (v *SetMergeRequestLabelsMergeRequestSetLabelsMergeRequestSetLabelsPayload) GetErrors() []string {
	return v.Errors
}

// GetMergeRequest returns SetMergeRequestLabelsMergeRequestSetLabelsMergeRequestSetLabelsPayload.MergeRequest, and is useful for accessing the field via an interface.
func (v *SetMergeRequestLabelsMergeRequestSetLabelsMergeRequestSetLabelsPayload) GetMergeRequest() MergeRequest {
	return v.MergeRequest
}

// SetMergeRequestLabelsResponse is returned by SetMergeRequestLabels on success.
type SetMergeRequestLabelsResponse struct {
	MergeRequestSetLabels SetMergeRequestLabelsMergeRequestSetLabelsMergeRequestSetLabelsPayload `json:"mergeRequestSetLabels"`
}

// GetMergeRequestSetLabels returns SetMergeRequestLabelsResponse.MergeRequestSetLabels, and is useful for accessing the field via an interface.
func (v *SetMergeRequestLabelsResponse) GetMergeRequestSetLabels() SetMergeRequestLabelsMergeRequestSetLabelsMergeRequestSetLabelsPayload {
	return v.MergeRequestSetLabels
}

// User includes the requested fields of the GraphQL type UserCore.
// The GraphQL type's documentation follows.
//
// Core represention of a GitLab user.
type User struct {
	// Human-readable name of the user. Returns `****` if the user is a project bot
	// and the requester does not have permission to view the project.
	Name string `json:"name"`
	// User's public email.
	PublicEmail string `json:"publicEmail"`
}

// GetName returns User.Name, and is useful for accessing the field via an interface.
func (v *User) GetName() string { return v.Name }

// GetPublicEmail returns User.PublicEmail, and is useful for accessing the field via an interface.
func (v *User) GetPublicEmail() string { return v.PublicEmail }

// __CreateNoteInput is used internally by genqlient
type __CreateNoteInput struct {
	Input CreateNoteInput `json:"input"`
}

// GetInput returns __CreateNoteInput.Input, and is useful for accessing the field via an interface.
func (v *__CreateNoteInput) GetInput() CreateNoteInput { return v.Input }

// __GetMergeRequestInput is used internally by genqlient
type __GetMergeRequestInput struct {
	Id string `json:"id"`
}

// GetId returns __GetMergeRequestInput.Id, and is useful for accessing the field via an interface.
func (v *__GetMergeRequestInput) GetId() string { return v.Id }

// __ListLabelsInput is used internally by genqlient
type __ListLabelsInput struct {
	Project string `json:"project"`
}

// GetProject returns __ListLabelsInput.Project, and is useful for accessing the field via an interface.
func (v *__ListLabelsInput) GetProject() string { return v.Project }

// __ListMergeRequestsInput is used internally by genqlient
type __ListMergeRequestsInput struct {
	Project string            `json:"project"`
	State   MergeRequestState `json:"state"`
}

// GetProject returns __ListMergeRequestsInput.Project, and is useful for accessing the field via an interface.
func (v *__ListMergeRequestsInput) GetProject() string { return v.Project }

// GetState returns __ListMergeRequestsInput.State, and is useful for accessing the field via an interface.
func (v *__ListMergeRequestsInput) GetState() MergeRequestState { return v.State }

// __SetMergeRequestLabelsInput is used internally by genqlient
type __SetMergeRequestLabelsInput struct {
	Input MergeRequestSetLabelsInput `json:"input"`
}

// GetInput returns __SetMergeRequestLabelsInput.Input, and is useful for accessing the field via an interface.
func (v *__SetMergeRequestLabelsInput) GetInput() MergeRequestSetLabelsInput { return v.Input }

func CreateNote(
	ctx context.Context,
	client graphql.Client,
	input CreateNoteInput,
) (*CreateNoteResponse, error) {
	req := &graphql.Request{
		OpName: "CreateNote",
		Query: `
mutation CreateNote ($input: CreateNoteInput!) {
	createNote(input: $input) {
		errors
	}
}
`,
		Variables: &__CreateNoteInput{
			Input: input,
		},
	}
	var err error

	var data CreateNoteResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func GetCurrentUser(
	ctx context.Context,
	client graphql.Client,
) (*GetCurrentUserResponse, error) {
	req := &graphql.Request{
		OpName: "GetCurrentUser",
		Query: `
query GetCurrentUser {
	currentUser {
		name
		publicEmail
	}
}
`,
	}
	var err error

	var data GetCurrentUserResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func GetMergeRequest(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*GetMergeRequestResponse, error) {
	req := &graphql.Request{
		OpName: "GetMergeRequest",
		Query: `
query GetMergeRequest ($id: MergeRequestID!) {
	mergeRequest(id: $id) {
		... MergeRequest
	}
}
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
		nodes {
			... Commit
		}
	}
	labels {
		nodes {
			... Label
		}
	}
	notes {
		nodes {
			... Note
		}
	}
}
fragment Commit on Commit {
	id
	authoredDate
	message
	sha
	title
	pipelines {
		nodes {
			... Pipeline
		}
	}
}
fragment Label on Label {
	id
	title
}
fragment Note on Note {
	body
	updatedAt
}
fragment Pipeline on Pipeline {
	sha
	status
}
`,
		Variables: &__GetMergeRequestInput{
			Id: id,
		},
	}
	var err error

	var data GetMergeRequestResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func ListLabels(
	ctx context.Context,
	client graphql.Client,
	project string,
) (*ListLabelsResponse, error) {
	req := &graphql.Request{
		OpName: "ListLabels",
		Query: `
query ListLabels ($project: ID!) {
	project(fullPath: $project) {
		labels {
			nodes {
				... Label
			}
		}
	}
}
fragment Label on Label {
	id
	title
}
`,
		Variables: &__ListLabelsInput{
			Project: project,
		},
	}
	var err error

	var data ListLabelsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func ListMergeRequests(
	ctx context.Context,
	client graphql.Client,
	project string,
	state MergeRequestState,
) (*ListMergeRequestsResponse, error) {
	req := &graphql.Request{
		OpName: "ListMergeRequests",
		Query: `
query ListMergeRequests ($project: ID!, $state: MergeRequestState!) {
	project(fullPath: $project) {
		... Project
	}
}
fragment Project on Project {
	id
	mergeRequests(state: $state, sort: UPDATED_ASC) {
		nodes {
			... MergeRequest
		}
	}
}
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
		nodes {
			... Commit
		}
	}
	labels {
		nodes {
			... Label
		}
	}
	notes {
		nodes {
			... Note
		}
	}
}
fragment Commit on Commit {
	id
	authoredDate
	message
	sha
	title
	pipelines {
		nodes {
			... Pipeline
		}
	}
}
fragment Label on Label {
	id
	title
}
fragment Note on Note {
	body
	updatedAt
}
fragment Pipeline on Pipeline {
	sha
	status
}
`,
		Variables: &__ListMergeRequestsInput{
			Project: project,
			State:   state,
		},
	}
	var err error

	var data ListMergeRequestsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func SetMergeRequestLabels(
	ctx context.Context,
	client graphql.Client,
	input MergeRequestSetLabelsInput,
) (*SetMergeRequestLabelsResponse, error) {
	req := &graphql.Request{
		OpName: "SetMergeRequestLabels",
		Query: `
mutation SetMergeRequestLabels ($input: MergeRequestSetLabelsInput!) {
	mergeRequestSetLabels(input: $input) {
		errors
		mergeRequest {
			... MergeRequest
		}
	}
}
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
		nodes {
			... Commit
		}
	}
	labels {
		nodes {
			... Label
		}
	}
	notes {
		nodes {
			... Note
		}
	}
}
fragment Commit on Commit {
	id
	authoredDate
	message
	sha
	title
	pipelines {
		nodes {
			... Pipeline
		}
	}
}
fragment Label on Label {
	id
	title
}
fragment Note on Note {
	body
	updatedAt
}
fragment Pipeline on Pipeline {
	sha
	status
}
`,
		Variables: &__SetMergeRequestLabelsInput{
			Input: input,
		},
	}
	var err error

	var data SetMergeRequestLabelsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

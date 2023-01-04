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

// GetProjectResponse is returned by GetProject on success.
type GetProjectResponse struct {
	// Find a project.
	Project Project `json:"project"`
}

// GetProject returns GetProjectResponse.Project, and is useful for accessing the field via an interface.
func (v *GetProjectResponse) GetProject() Project { return v.Project }

// Label includes the GraphQL fields of Label requested by the fragment Label.
type Label struct {
	// Content of the label.
	Title string `json:"title"`
}

// GetTitle returns Label.Title, and is useful for accessing the field via an interface.
func (v *Label) GetTitle() string { return v.Title }

// MergeRequest includes the GraphQL fields of MergeRequest requested by the fragment MergeRequest.
type MergeRequest struct {
	// ID of the merge request.
	Id string `json:"id"`
	// Internal ID of the merge request.
	Iid string `json:"iid"`
	// Title of the merge request.
	Title string `json:"title"`
	// Diff head SHA of the merge request.
	DiffHeadSha string `json:"diffHeadSha"`
	// Source branch of the merge request.
	SourceBranch string `json:"sourceBranch"`
	// Target branch of the merge request.
	TargetBranch string `json:"targetBranch"`
	// Details about which files were changed in this merge request.
	DiffStats []MergeRequestDiffStats `json:"diffStats"`
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

// GetDiffHeadSha returns MergeRequest.DiffHeadSha, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetDiffHeadSha() string { return v.DiffHeadSha }

// GetSourceBranch returns MergeRequest.SourceBranch, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetSourceBranch() string { return v.SourceBranch }

// GetTargetBranch returns MergeRequest.TargetBranch, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetTargetBranch() string { return v.TargetBranch }

// GetDiffStats returns MergeRequest.DiffStats, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetDiffStats() []MergeRequestDiffStats { return v.DiffStats }

// GetCommits returns MergeRequest.Commits, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetCommits() MergeRequestCommitsCommitConnection { return v.Commits }

// GetLabels returns MergeRequest.Labels, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetLabels() MergeRequestLabelsLabelConnection { return v.Labels }

// GetNotes returns MergeRequest.Notes, and is useful for accessing the field via an interface.
func (v *MergeRequest) GetNotes() MergeRequestNotesNoteConnection { return v.Notes }

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

// State of a GitLab merge request
type MergeRequestState string

const (
	// All available.
	MergeRequestStateAll MergeRequestState = "all"
	// In closed state.
	MergeRequestStateClosed MergeRequestState = "closed"
	// Discussion has been locked.
	MergeRequestStateLocked MergeRequestState = "locked"
	// Merge request has been merged.
	MergeRequestStateMerged MergeRequestState = "merged"
	// In open state.
	MergeRequestStateOpened MergeRequestState = "opened"
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
	// Pipeline was canceled before completion.
	PipelineStatusEnumCanceled PipelineStatusEnum = "CANCELED"
	// Pipeline has been created.
	PipelineStatusEnumCreated PipelineStatusEnum = "CREATED"
	// At least one stage of the pipeline failed.
	PipelineStatusEnumFailed PipelineStatusEnum = "FAILED"
	// Pipeline needs to be manually started.
	PipelineStatusEnumManual PipelineStatusEnum = "MANUAL"
	// Pipeline has not started running yet.
	PipelineStatusEnumPending PipelineStatusEnum = "PENDING"
	// Pipeline is preparing to run.
	PipelineStatusEnumPreparing PipelineStatusEnum = "PREPARING"
	// Pipeline is running.
	PipelineStatusEnumRunning PipelineStatusEnum = "RUNNING"
	// Pipeline is scheduled to run.
	PipelineStatusEnumScheduled PipelineStatusEnum = "SCHEDULED"
	// Pipeline was skipped.
	PipelineStatusEnumSkipped PipelineStatusEnum = "SKIPPED"
	// Pipeline completed successfully.
	PipelineStatusEnumSuccess PipelineStatusEnum = "SUCCESS"
	// A resource (for example, a runner) that the pipeline requires to run is unavailable.
	PipelineStatusEnumWaitingForResource PipelineStatusEnum = "WAITING_FOR_RESOURCE"
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

// __GetProjectInput is used internally by genqlient
type __GetProjectInput struct {
	Project string            `json:"project"`
	State   MergeRequestState `json:"state"`
}

// GetProject returns __GetProjectInput.Project, and is useful for accessing the field via an interface.
func (v *__GetProjectInput) GetProject() string { return v.Project }

// GetState returns __GetProjectInput.State, and is useful for accessing the field via an interface.
func (v *__GetProjectInput) GetState() MergeRequestState { return v.State }

func GetProject(
	ctx context.Context,
	client graphql.Client,
	project string,
	state MergeRequestState,
) (*GetProjectResponse, error) {
	req := &graphql.Request{
		OpName: "GetProject",
		Query: `
query GetProject ($project: ID!, $state: MergeRequestState!) {
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
	diffHeadSha
	sourceBranch
	targetBranch
	diffStats {
		path
	}
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
		Variables: &__GetProjectInput{
			Project: project,
			State:   state,
		},
	}
	var err error

	var data GetProjectResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

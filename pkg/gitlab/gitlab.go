package gitlab

import (
	"errors"
	"regexp"
)

// GetLatestCommit returns the commit matching the HEAD SHA of the parent merge request
func (mr *MergeRequest) GetLatestCommit() (Commit, error) {
	for _, commit := range mr.GetCommits().Nodes {
		if commit.Sha != mr.DiffHeadSha {
			continue
		}
		return commit, nil
	}
	return Commit{}, errors.New("could not find latest commit")
}

// GetProjectPath extracts project path from URI (repository URL).
func (mr *MergeRequest) GetProjectPath() string {
	r, _ := regexp.Compile("(https?|ssh)://([^/]*)/(.*)\\.git$")
	return r.FindStringSubmatch(mr.SourceProject.HttpUrlToRepo)[3]
}

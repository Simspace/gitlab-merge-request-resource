package gitlab

import "errors"

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

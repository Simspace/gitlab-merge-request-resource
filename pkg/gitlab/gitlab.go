package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// GetLatestCommit returns the commit matching the HEAD SHA of the merge request
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

// WrapErrors extracts error messages from Gitlab mutation API responses and
// returns them as Go error types along with the returned response.
func WrapErrors(i interface{}, e error) (interface{}, error) {
	if e != nil {
		return i, e
	}

	// Marshal the object into JSON and then unmarshal into a map to easily do a
	// deep search for "errors" keys
	j, err := json.Marshal(i)
	if err != nil {
		return i, err
	}

	var jm map[string]interface{}

	err = json.Unmarshal(j, &jm)
	if err != nil {
		return i, err
	}

	errs := deepGetErrs(jm)

	if len(errs) == 1 {
		return i, errors.New(errs[0])
	}

	if len(errs) > 1 {
		return i, fmt.Errorf("multiple API errors: %s", strings.Join(errs, ","))
	}

	return i, nil
}

func deepGetErrs(m map[string]interface{}) []string {
	errs := []string{}
	for k := range m {
		if k == "errors" {
			if e, ok := m[k].([]interface{}); ok {
				for _, s := range e {
					errs = append(errs, s.(string))
				}
				return errs
			}
			// We should never hit this, but return an error from *this* func if we do.
			return []string{"could not read errors"}
		}
		if v, ok := m[k].(map[string]interface{}); ok {
			errs = deepGetErrs(v)
		}
	}
	return errs
}

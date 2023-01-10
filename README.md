# GitLab Merge Request Concourse Resource

A concourse resource to check for new merge requests on GitLab and update the merge request status.

The provided Docker image is Alpine-based.

## Source Configuration

```yaml
resource_types:
- name: merge-request
  type: docker-image
  source:
    repository: simspace/gitlab-merge-request-resource

resources:
- name: merge-request
  type: merge-request
  source:
    uri: https://gitlab.com/myname/myproject.git
    private_token: XXX
```

* `uri`: (required) The location of the repository (required)
* `private_token`: (required) Your GitLab user's private token (required, can be found in your profile settings)
* `insecure`: When set to `true`, SSL verification is turned off 
* `skip_work_in_progress`: When set to `true`, merge requests mark as work in progress (WIP) will be skipped. Default `false`
* `skip_not_mergeable`: When set to `true`, merge requests not marked as mergeable will be skipped. Default `false`
* `skip_trigger_comment`: When set to `true`, the resource will not look up for `[trigger ci]` merge request comments to manually trigger builds. Default `false`
* `concourse_url`: When set, this url will be used to override `ATC_EXTERNAL_URL` during commit status updates.
* `pipeline_name`(string): When set, this url will be used to override `BUILD_PIPELINE_NAME` during commit status updates.
* `labels`(string[]): Filter merge requests by label`[]`
* `paths` (string[]): Include merge request if one of the modified files match a path pattern (glob) or is included inside a directory defined here. Default: include all. 
* `ignore_paths` (string[]): Exclude merge request if one of the modified files matches a path pattern (glob) or is included inside a directory defined here. Default: exclude none. 
* `target_branch`(string): Filter merge requests by target_branch. Default is empty string.
* `source_branch`(string): Filter merge requests by source_branch. Default is empty string.
* `ssh_keys` (string[]): When set to a non-empty array, an ssh-agent will be started and the specified keys will be added to it.  This is only relevant for submodules with an ssh URL and passphrase encrypted keys are not supported.
* `recursive`: When set to `true`, will pull submodules by issuing a `git submodule update --init --recursive`.  Note that if your submodules are hosted on the same server, be sure to [use a relative path](https://www.gniibe.org/memo/software/git/using-submodule.html) to avoid ssh/https protocol clashing (as the MR is fetched via https, this resource would have no way to authenticate a git+ssh connection).

## Behavior

### `check`: Check for new merge requests

Checks if there are new merge requests or merge requests with new commits.

### `in`: Clone merge request source branch

`git clone`s the source branch of the respective merge request.

If you need to retrieve any information about the merge request in your tasks, the script writes the raw API response of the
[get single merge request call](https://docs.gitlab.com/ee/api/merge_requests.html#get-single-mr) to `.git/merge-request.json`. 
The name of the source branch is extracted to `.git/merge-request-source-branch` for convenience. 
The paths of all changed files are also saved to
`.git/resources/changed_files`.

### `out`: Update a merge request's merge status

Updates the merge request's `merge_status` which displays nicely in the GitLab UI and allows to only merge changes if they pass the test.

#### Parameters

* `repository`: The path of the repository of the merge request's source branch (required)
* `status`: The new status of the merge request (required, can be either `pending`, `running`, `success`, `failed`, or `canceled`)
* `labels`(string[]): The labels you want to add to your merge request
* `comment`: Add a comment for MR. Could be an object with `text`/`file` fields. If just the `file` or `text` is specified it is used to populate the field, if both `file` and `text` are specified then the file is substituted in to replace $FILE_CONTENT in the text.

## Example

```yaml
jobs:
- name: sample-merge-request
  plan:
  - get: merge-request
    trigger: true
  - put: merge-request
    params:
      repository: merge-request
      status: running
  - task: unit-test
    file: merge-request/ci/tasks/unit-test.yml
  on_failure:
    put: merge-request
    params:
      repository: merge-request
      status: failed
  on_success:
    put: merge-request
    params:
      repository: merge-request
      status: success
      labels: ['unit-test', 'stage']
      comment:
        file: out/commt.txt
        text: |
          Add new comment.
          $FILE_CONTENT
```
## Development

This resource uses genqlient for handling GraphQL APIs, by autogenerating Go
types from GraphQL queries. 

### Updating GraphQL queries or schema
The Gitlab API schema is included here to aid in
autogeneration but may need to be updated from time to time. The schema can be
downloaded by using the [Apollo](https://www.apollographql.com/docs/devtools/cli/) CLI.

* Install the Apollo CLI.
    * On Mac systems: `brew install apollo-cli`
    * On other sytems: `npm install -D apollo`
* Run the CLI to download the Gitlab schema:
    ```sh
    apollo client:download-schema \
      --endpoint=https://gitlab.com/api/graphql \
      gitlab_schema.graphql
    ```

With the schema downloaded, you can use genqlient to update or create new API
queries.

In `gitlab/graphql.go` there are several throwaway constants. Here we can
create a new constant (assigned to `_`, thus not used anywhere) with our raw
GraphQL query as a string. At the head of the string, add `# @genqlient` to let
genqlient know to generate this query.

A basic query might look like this:
```go
_ = `# @genqlient
query GetCurrentUser {
  currentUser{
    name
    publicEmail
  }
}`
```

Configuration commands can also be added as comments. A more complex example
might look like this: 
```go
_ = `# @genqlient
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
```

Once you have added all of the API queries and saved the file, run `go
generate` from the root of the repository to autogenerate the code.

Now, you can use the query as a Go function with the same name. So, our first
example can be called like this:
```go
resp := gitlab.GetCurrentUser()

fmt.Printf("Username: %s\n", resp.CurrentUser.Name)
```

See the
[genqlient
docs](https://github.com/Khan/genqlient/blob/main/docs/INTRODUCTION.md) for
more details and in-depth configuration.


### Testing
This project uses the Ginkgo test framework with Gomega matchers. All new
functionality should have accompanying test specs.

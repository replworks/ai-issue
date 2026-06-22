# TASKS.md

## MVP

### Foundation

* [X] Initialize project
* [X] Create CLI entrypoint
* [X] Establish project structure
* [X] Verify framework compliance

---

### Content Processing

* [X] Read issue content from clipboard
* [X] Validate content exists
* [X] Extract issue title
* [X] Extract issue body
* [X] Validate extracted issue draft

Acceptance Criteria:

* Markdown with a first H1, H2, or H3 heading produces a valid issue draft
* Markdown without content is rejected
* Missing title is rejected

---

### Repository Resolution

* [X] Determine target repository
* [X] Validate resolved repository

Acceptance Criteria:

* Repository can be resolved automatically
* Invalid repositories are rejected

---

### Publisher Resolution

* [X] Determine publishing user
* [X] Validate resolved publisher

Acceptance Criteria:

* Publisher information is available before publication

---

### Authentication

* [ ] Implement GitHub App device flow login
* [ ] Store and load GitHub App token locally
* [ ] Use stored token for publishing and diagnostics

Acceptance Criteria:

* `ai-issue login` obtains a GitHub App user access token
* Token is persisted locally
* Publishing works without manual PAT setup

---

### Issue Construction

* [X] Construct publishable issue
* [X] Preserve publisher traceability
* [X] Validate final publishable issue payload

Acceptance Criteria:

* Final issue contains title
* Final issue contains body
* Publisher information is preserved

---

### Preview Flow

* [X] Display publication preview
* [X] Display repository information
* [X] Display issue title
* [X] Request explicit approval

Acceptance Criteria:

* Publication cannot occur without approval

---

### Issue Publication

* [X] Create GitHub Issue
* [X] Return issue URL
* [X] Handle publication failures

Acceptance Criteria:

* Issue is created successfully
* Created issue URL is returned

---

### Diagnostics

* [X] Validate publishing prerequisites
* [X] Report diagnostic results

Acceptance Criteria:

* User can determine whether publication is possible before attempting publication

---

### Error Handling

* [X] Handle missing content
* [X] Handle invalid extracted issue drafts
* [X] Handle repository failures
* [X] Handle authentication failures
* [X] Handle publication failures

Acceptance Criteria:

* Failures produce actionable messages

---

### Testing

* [X] Unit tests
* [X] End-to-end publication flow test
* [X] Architecture compliance validation
* [X] Manual MVP flow validation

Acceptance Criteria:

* MVP flow works from content input to issue creation

---

### Release

* [X] Prepare release artifacts
* [X] Publish v0.1.0

Acceptance Criteria:

* End users can install and execute the application

### Additional Functions

* [x] Implement GitHub repository access validation in the `diagnose` command.
* [x] Add `--version` Support #19.
* [x] Fix Title Extraction for Inline Code in Markdown Headings.
* [x] Make Publisher Identity Configurable.
* [x] Add `--dry-run` Mode.

---

## Future

### Publication

* [ ] Support explicit repository selection
* [ ] Support non-interactive approval

---

### Input Sources

* [ ] Support file input
* [ ] Support stdin input

---

### Issue Customization

* [ ] Support custom metadata templates
* [ ] Support issue labels
* [ ] Support issue templates

---

### Multi-Identity

* [ ] Support multiple AI authors
* [ ] Support organization-wide configuration

---

### Platform Expansion

* [ ] Design hosted version

# TASKS.md

## MVP

### Foundation

* [X] Initialize project
* [X] Create CLI entrypoint
* [X] Establish project structure
* [ ] Verify framework compliance

---

### Content Processing

* [X] Read issue content from clipboard
* [X] Validate content exists
* [X] Extract issue title
* [X] Extract issue body
* [X] Validate extracted issue draft

Acceptance Criteria:

* Markdown with a first H1 heading produces a valid issue draft
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
* [ ] Architecture compliance validation
* [ ] Manual MVP flow validation

Acceptance Criteria:

* MVP flow works from content input to issue creation

---

### Release

* [ ] Prepare release artifacts
* [ ] Publish v0.1.0

Acceptance Criteria:

* End users can install and execute the application

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

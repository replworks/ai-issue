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
* [ ] Validate issue draft

Acceptance Criteria:

* Valid markdown produces a valid issue draft
* Invalid markdown is rejected
* Missing title is rejected

---

### Repository Resolution

* [X] Determine target repository
* [ ] Validate repository information

Acceptance Criteria:

* Repository can be resolved automatically
* Invalid repositories are rejected

---

### Publisher Resolution

* [X] Determine publishing user
* [ ] Validate publisher information

Acceptance Criteria:

* Publisher information is available before publication

---

### Issue Construction

* [X] Construct publishable issue
* [ ] Preserve publisher traceability
* [ ] Validate final issue payload

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
* [ ] Handle publication failures

Acceptance Criteria:

* Issue is created successfully
* Created issue URL is returned

---

### Diagnostics

* [X] Validate publishing prerequisites
* [ ] Report validation results

Acceptance Criteria:

* User can determine whether publication is possible before attempting publication

---

### Error Handling

* [X] Handle missing content
* [ ] Handle invalid issue drafts
* [X] Handle repository failures
* [X] Handle authentication failures
* [X] Handle publication failures

Acceptance Criteria:

* Failures produce actionable messages

---

### Testing

* [X] Unit tests
* [ ] End-to-end publication test
* [ ] Architecture validation
* [ ] Manual MVP validation

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

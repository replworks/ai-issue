# TASKS.md

## MVP

### Foundation

* [ ] Initialize project
* [ ] Create CLI entrypoint
* [ ] Establish project structure
* [ ] Verify framework compliance

---

### Content Processing

* [ ] Read issue content from clipboard
* [ ] Validate content exists
* [ ] Extract issue title
* [ ] Extract issue body
* [ ] Validate issue draft

Acceptance Criteria:

* Valid markdown produces a valid issue draft
* Invalid markdown is rejected
* Missing title is rejected

---

### Repository Resolution

* [ ] Determine target repository
* [ ] Validate repository information

Acceptance Criteria:

* Repository can be resolved automatically
* Invalid repositories are rejected

---

### Publisher Resolution

* [ ] Determine publishing user
* [ ] Validate publisher information

Acceptance Criteria:

* Publisher information is available before publication

---

### Issue Construction

* [ ] Construct publishable issue
* [ ] Preserve publisher traceability
* [ ] Validate final issue payload

Acceptance Criteria:

* Final issue contains title
* Final issue contains body
* Publisher information is preserved

---

### Preview Flow

* [ ] Display publication preview
* [ ] Display repository information
* [ ] Display issue title
* [ ] Request explicit approval

Acceptance Criteria:

* Publication cannot occur without approval

---

### Issue Publication

* [ ] Create GitHub Issue
* [ ] Return issue URL
* [ ] Handle publication failures

Acceptance Criteria:

* Issue is created successfully
* Created issue URL is returned

---

### Diagnostics

* [ ] Validate publishing prerequisites
* [ ] Report validation results

Acceptance Criteria:

* User can determine whether publication is possible before attempting publication

---

### Error Handling

* [ ] Handle missing content
* [ ] Handle invalid issue drafts
* [ ] Handle repository failures
* [ ] Handle authentication failures
* [ ] Handle publication failures

Acceptance Criteria:

* Failures produce actionable messages

---

### Testing

* [ ] Unit tests
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

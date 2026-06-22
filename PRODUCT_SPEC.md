# PRODUCT_SPEC.md

# AI Issue Publisher

## Purpose

AI Issue Publisher is a CLI tool that publishes AI-generated work items as GitHub Issues.

The product exists to ensure valuable AI-generated work is preserved inside project backlogs.

The product also ensures that AI-generated work and human-generated work remain distinguishable within GitHub.

---

# Problem

AI conversations frequently produce valuable work items.

Examples:

* bug reports
* feature requests
* refactoring proposals
* architecture ideas
* technical debt items
* implementation plans

These work items are often lost because publishing them to GitHub requires manual effort.

A second problem exists.

When AI-generated content is published manually, GitHub typically presents the publishing human as the issue author.

This creates ambiguity regarding authorship and responsibility.

---

# Product Goal

The product must allow users to quickly publish AI-generated work items to GitHub.

The product must preserve the distinction between:

* the creator of the content
* the publisher of the content

---

# Users

Primary users:

* software developers
* solo founders
* indie hackers
* AI-assisted development teams

---

# Primary Use Case

A user receives a useful work item from an AI conversation.

The user decides the work item should be preserved.

The user publishes the work item as a GitHub Issue.

The issue becomes part of the project backlog.

---

# Inputs

The product accepts markdown content.

Example:

```markdown
# Add timestamps to logging system

Current logs do not contain timestamps.
```

---

# Outputs

The product creates a GitHub Issue.

The created issue becomes visible in the target repository.

The product returns the created issue URL.

---

# Functional Requirements

## FR-001

The product shall create GitHub Issues from markdown content.

---

## FR-002

The product shall determine an issue title from the provided markdown.

---

## FR-003

The product shall determine an issue body from the provided markdown.

---

## FR-004

The product shall determine the target GitHub repository automatically.

---

## FR-005

The product shall identify the publishing user.

---

## FR-006

The product shall display a preview before issue creation.

---

## FR-007

The product shall require explicit confirmation before creating an issue.

---

## FR-008

The product shall create the GitHub Issue after confirmation.

---

## FR-009

The product shall return the created issue URL.

---

## FR-010

The product shall provide a diagnostic command that validates publishing prerequisites.

---

# Issue Construction Rules

## Title Rule

The first heading at the top of the markdown becomes the issue title.

Supported heading levels:

* H1
* H2
* H3

Example:

```markdown
# Add timestamps to logging system
```

becomes

```text
Add timestamps to logging system
```

The same rule applies to the first `##` or `###` heading when it appears at the top of the markdown.

---

## Body Rule

The issue body is generated from the remaining markdown content after title extraction.

---

## Publisher Rule

Publisher information must be preserved as part of the generated issue.

---

# User Flow

## Publish Issue

1. User copies AI-generated markdown.
2. User executes the publish command.
3. The product analyzes the content.
4. The product generates a preview.
5. The user confirms publication.
6. The product creates a GitHub Issue.
7. The product displays the created issue URL.

---

## Diagnostic Check

1. User executes the diagnostic command.
2. The product validates required prerequisites.
3. The product reports results.

---

# Error Conditions

Issue creation must fail when:

* no content is available
* repository cannot be determined
* issue title cannot be determined
* issue publication is not possible

The product must display a human-readable error message.

---

# Non-Goals

The following are outside the scope of this product:

* AI model integration
* automatic issue generation
* pull request generation
* commit generation
* Slack integration
* Jira integration
* project management functionality
* analytics
* reporting
* AI performance tracking

---

# Acceptance Criteria

The product is complete when:

1. A user can publish AI-generated markdown as a GitHub Issue.
2. The issue is created in the correct repository.
3. The user can review content before publication.
4. Publisher information is preserved.
5. The created issue URL is returned.

---

# Success Criteria

The product is successful when:

* valuable AI-generated work is not lost
* issue publication takes less than one minute
* publishing requires minimal manual effort
* AI-generated work can be preserved consistently

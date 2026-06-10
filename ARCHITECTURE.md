# ARCHITECTURE.md

# AI Issue Publisher Architecture

## Purpose

This document defines the internal structure of the product.

PRODUCT_SPEC.md defines product requirements.

ARCHITECTURE.md defines how the product operates internally.

FRAMEWORK.md defines implementation constraints.

This document must remain technology-agnostic.

Do not introduce implementation details.

Do not introduce programming language details.

Do not introduce library-specific behavior.

---

# Core Concept

The system exists to preserve a distinction between:

```text
Content Creation
```

and

```text
Content Publication
```

These are independent responsibilities.

The creator of content and the publisher of content must not be treated as the same entity.

---

# Primary Invariant

The following condition must always hold:

```text
Author != Publisher
```

Definitions:

```text
Author
    Entity that generated issue content.

Publisher
    Human that decided to publish issue content.
```

Any architecture that causes Author and Publisher to become the same conceptual entity is invalid.

---

# Architectural Goals

The architecture must:

* preserve AI authorship
* preserve publisher traceability
* minimize publication effort
* prevent loss of AI-generated work
* support repository-local operation

---

# System Flow

The system processes content through the following stages:

```text
Content Source
        ↓
Issue Extraction
        ↓
Issue Construction
        ↓
Issue Preview
        ↓
Issue Publication
        ↓
Issue URL
```

Each stage has a single responsibility.

---

# Content Source

Responsibilities:

* obtain issue content
* provide content for processing

Outputs:

```text
Raw Issue Content
```

The content source is not responsible for:

* title generation
* issue creation
* repository selection

---

# Issue Extraction

Responsibilities:

* determine issue title
* determine issue body
* validate issue content

Outputs:

```text
Issue Draft
```

Example:

```text
Title

Body
```

Issue Extraction is not responsible for publication.

---

# Publisher Resolution

Responsibilities:

* determine publishing identity
* provide publication context

Outputs:

```text
Publisher
```

Publisher Resolution is not responsible for issue authorship.

---

# Repository Resolution

Responsibilities:

* determine target repository

Outputs:

```text
Repository
```

Repository Resolution must occur before publication.

---

# Issue Construction

Responsibilities:

* combine issue draft
* combine publisher information
* construct final issue payload

Outputs:

```text
Publishable Issue
```

Issue Construction is responsible for preserving authorship and publisher traceability.

---

# Issue Preview

Responsibilities:

* display publication result before publication
* allow publication approval

Outputs:

```text
Approved Issue
```

Issue creation must not occur before approval.

---

# Issue Publication

Responsibilities:

* create issue
* return issue identifier
* return issue URL

Outputs:

```text
Created Issue
```

Issue Publication is the only stage allowed to create GitHub Issues.

---

# Responsibility Boundaries

## Content Source

Owns:

```text
Content
```

Does Not Own:

```text
Repository
Publisher
Issue Creation
```

---

## Issue Extraction

Owns:

```text
Title
Body
Validation
```

Does Not Own:

```text
Publication
```

---

## Repository Resolution

Owns:

```text
Repository Selection
```

Does Not Own:

```text
Issue Content
```

---

## Publisher Resolution

Owns:

```text
Publisher Identity
```

Does Not Own:

```text
Issue Authorship
```

---

## Issue Construction

Owns:

```text
Final Issue Payload
```

Does Not Own:

```text
Issue Creation
```

---

## Issue Publication

Owns:

```text
Issue Creation
```

Does Not Own:

```text
Issue Content Generation
```

---

# Architectural Rules

## Rule 1

Issue publication must occur only after issue construction completes.

---

## Rule 2

Issue publication must occur only after explicit approval.

---

## Rule 3

Publisher information must remain available after publication.

---

## Rule 4

Repository selection must complete before issue publication.

---

## Rule 5

Issue content generation and issue publication must remain separate responsibilities.

---

# Failure Boundaries

The following failures must stop publication:

```text
No Content

No Repository

Invalid Issue Draft

Publication Rejected

Publication Failure
```

No partial publication is allowed.

---

# Non-Goals

This architecture does not support:

* automatic issue generation
* automatic issue publication
* pull request generation
* commit generation
* project management workflows
* issue prioritization
* issue analytics

---

# Architectural Invariants

The following conditions must always remain true:

```text
Author != Publisher
```

```text
Issue Creation requires Approval
```

```text
Repository determined before Publication
```

```text
Issue Draft created before Publication
```

```text
Issue Publication does not generate Issue Content
```

```text
Publisher does not become Author
```

Any implementation that violates these invariants is architecturally incorrect.

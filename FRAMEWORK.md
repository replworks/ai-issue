# FRAMEWORK.md

# Go CLI Framework

## Purpose

This document defines implementation constraints.

PRODUCT_SPEC.md defines product requirements.

ARCHITECTURE.md defines system structure.

FRAMEWORK.md defines implementation rules.

All implementation decisions must follow this document.

Do not replace technologies.

Do not introduce alternative stacks.

Do not redesign implementation conventions.

---

# Language

Use:

```text
Go >= 1.24
```

Do not use:

```text
Node.js
TypeScript
Python
PHP
Java
Rust
```

---

# Architecture Style

Use:

```text
Single Binary CLI
```

Requirements:

* standalone executable
* no backend service
* no database
* no web server
* no daemon process

---

# Command Framework

Use:

```text
github.com/spf13/cobra
```

All commands must be implemented through Cobra.

---

# Configuration

Configuration format:

```text
YAML
```

Location:

```text
~/.config/<application>/config.yaml
```

Configuration must be optional.

Applications must start with sensible defaults whenever possible.

---

# Secrets

Use environment variables for secrets.

Never:

* hardcode secrets
* commit secrets
* store secrets in repository files

---

# Markdown Processing

When markdown parsing is required:

Use:

```text
AST-based parsing
```

Avoid:

```text
Regular-expression-only parsing
```

Structured formats should be parsed structurally rather than through string manipulation.

---

# External Integrations

External services must be isolated behind dedicated adapters.

Business logic must not depend directly on external SDKs or APIs.

Implementation should allow replacement of integration mechanisms without changing business logic.

---

# Directory Structure

Use:

```text
cmd/
    <application>/

internal/
```

Structure internal packages around responsibilities.

Example:

```text
internal/
    parser/
    publisher/
    repository/
```

Avoid generic package names:

```text
utils/
helpers/
common/
shared/
misc/
```

Packages should communicate intent through naming.

---

# Error Handling

Return explicit errors.

Do not silently recover from unexpected failures.

Error messages must be actionable.

Examples:

```text
Clipboard is empty.
```

```text
Repository could not be determined.
```

```text
Authentication failed.
```

Avoid:

```text
Unknown error.
```

---

# Logging

Prefer structured logging.

Logs must assist troubleshooting.

Do not log:

* secrets
* tokens
* credentials

---

# Testing

Use:

```text
Go standard testing package
```

Requirements:

* unit tests for business logic
* deterministic tests
* isolated tests

Tests must not require:

* network access
* external services
* manual interaction

External dependencies should be mocked.

---

# Dependency Management

Minimize dependencies.

Before introducing a dependency:

1. Verify standard library is insufficient.
2. Verify dependency solves a meaningful problem.
3. Verify dependency is actively maintained.

Avoid dependency duplication.

---

# Distribution

Distribution method:

```text
GitHub Releases
```

Produce standalone binaries for:

```text
macOS
Linux
Windows
```

Installation must not require compilation from source.

---

# Coding Principles

Prefer:

```text
Simple
Explicit
Readable
Deterministic
```

Avoid:

```text
Premature optimization
Speculative abstraction
Framework-like complexity
Hidden behavior
```

Code should be understandable without extensive context.

---

# Change Management

When architecture changes:

1. Update ARCHITECTURE.md.
2. Update implementation.

When requirements change:

1. Update PRODUCT_SPEC.md.
2. Update ARCHITECTURE.md if necessary.
3. Update implementation.

Architecture and implementation must not diverge.

---

# Framework Invariants

The following conditions must always remain true:

```text
Single Binary CLI
```

```text
Configuration is externalized
```

```text
Secrets are not stored in source code
```

```text
Business logic is isolated from external integrations
```

```text
Tests do not require external services
```

```text
Implementation favors simplicity over abstraction
```

Any implementation that violates these invariants is incorrect.

# AGENTS.md

## DOCUMENT_ORDER

1. AGENTS.md
2. .replworks/PRODUCT_SPEC.md
3. .replworks/TECH_STACK.md
4. .replworks/ARCHITECTURE.md
5. .replworks/TASKS.md
   Only these documents are authoritative.

---

## IGNORE

Ignore all files under:

```text
docs/
```

Never use files in docs/ as requirements.
Never implement features described only in docs/.

---

## SOURCE_OF_TRUTH

Product Requirements:

```text
PRODUCT_SPEC.md
```

Implementation Constraints:

```text
TECH_STACK.md
```

Architecture:

```text
ARCHITECTURE.md
```

Execution Plan:

```text
TASKS.md
```

If a conflict exists:

```text
PRODUCT_SPEC.md
>
TECH_STACK.md
>
ARCHITECTURE.md
>
TASKS.md
>
everything else
```

---

## DOCUMENT_RESPONSIBILITIES

PRODUCT_SPEC.md defines:

```text
What the product is.
What the product does.
```

TECH_STACK.md defines:

```text
How the product must be implemented.
```

ARCHITECTURE.md defines:

```text
How the product works.
```

TASKS.md defines:

```text
What should be implemented next.
```

Do not move responsibilities between documents.

---

## EXTERNAL_BOUNDARY

Define once. Referenced by TASK_EXECUTION and MOCK_RULES below.

```text
External boundary = any behavior not controlled by this codebase.
Examples:
third-party DOM
third-party API
browser runtime behavior
```

---

## IMPLEMENTATION_RULES

Implement only the selected task.
Do not implement:

```text
future work
roadmap items
optional features
assumptions
inferred requirements
```

Requirements must originate from:

```text
PRODUCT_SPEC.md
```

Implementation must follow:

```text
TECH_STACK.md
ARCHITECTURE.md
```

---

## TASK_EXECUTION

For every task:

1. Read PRODUCT_SPEC.md
2. Read TECH_STACK.md
3. Read ARCHITECTURE.md
4. Read task definition
5. If the task touches a domain not covered by verified knowledge in PRODUCT_SPEC.md or ARCHITECTURE.md: stop. Mark the relevant section UNVERIFIED. Do not implement against an UNVERIFIED section. Require explicit human confirmation before continuing.
6. Implement
7. Write unit tests for internal logic
8. If the task touches an EXTERNAL_BOUNDARY: write an E2E test against the live boundary. A mocked test alone does not satisfy this step.
9. Run all tests
10. Stop
    Do not start another task automatically.

---

## MOCK_RULES

Mock only observed behavior.

```text
Allowed sources:
recorded live response
documented spec
```

```text
Forbidden sources:
assumed behavior
guessed response
inferred event flow
```

If a mock's values cannot be traced to a recorded observation or a spec, do not write it.
Any code touching an EXTERNAL_BOUNDARY requires at least one live observation before it may be mocked.
Re-verify mocks when the external system's behavior may have changed.

---

## PRODUCT_SPEC_CHANGES

If implementation reveals missing product requirements, or a PRODUCT_SPEC.md section is marked UNVERIFIED:

```text
Stop.
Do not invent requirements.
```

Update PRODUCT_SPEC.md, and clear the UNVERIFIED mark only after human confirmation, before implementation continues.

---

## TECH_STACK_CHANGES

If implementation requires tech stack changes:

1. Update TECH_STACK.md
2. Update implementation
   Never allow tech stack and code to diverge.

---

## ARCHITECTURE_CHANGES

If implementation requires architecture changes, or an ARCHITECTURE.md section is marked UNVERIFIED:

1. Update ARCHITECTURE.md
2. Clear the UNVERIFIED mark only after human confirmation
3. Update implementation
   Never allow architecture and code to diverge.

---

## TASK_CHANGES

If implementation invalidates a task:
Update TASKS.md.

---

## DESIGN_RULES

Prefer:

```text
simple
explicit
minimal
```

Avoid:

```text
abstraction without use
premature optimization
speculative features
```

---

## FILE_CREATION

Do not create new top-level documents unless explicitly requested.
Prefer modifying existing files.

---

## SUCCESS_CRITERIA

Task is complete only when:

- product requirements satisfied
- architectural requirements satisfied
- tech stack constraints satisfied
- acceptance criteria satisfied
- no UNVERIFIED sections remain in scope for this task
- code runs
- unit tests pass
- E2E tests pass for any EXTERNAL_BOUNDARY code touched
  Then stop.

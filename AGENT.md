# AGENTS.md

## DOCUMENT_ORDER

1. AGENTS.md
2. PRODUCT_SPEC.md
3. ARCHITECTURE.md
4. FRAMEWORK.md
5. TASKS.md

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

Architecture:

```text
ARCHITECTURE.md
```

Implementation Constraints:

```text
FRAMEWORK.md
```

Execution Plan:

```text
TASKS.md
```

If a conflict exists:

```text
PRODUCT_SPEC.md
>
ARCHITECTURE.md
>
FRAMEWORK.md
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

ARCHITECTURE.md defines:

```text
How the product works.
```

FRAMEWORK.md defines:

```text
How the product must be implemented.
```

TASKS.md defines:

```text
What should be implemented next.
```

Do not move responsibilities between documents.

---

## IMPLEMENTATION_RULES

Implement only the selected task.

Do not implement:

* future work
* roadmap items
* optional features
* assumptions
* inferred requirements

Requirements must originate from:

```text
PRODUCT_SPEC.md
```

Implementation must follow:

```text
ARCHITECTURE.md
FRAMEWORK.md
```

---

## TASK_EXECUTION

For every task:

1. Read PRODUCT_SPEC.md
2. Read ARCHITECTURE.md
3. Read FRAMEWORK.md
4. Read task definition
5. Implement
6. Test
7. Stop

Do not start another task automatically.

---

## PRODUCT_CHANGES

If implementation reveals missing product requirements:

Stop.

Do not invent requirements.

Update:

```text
PRODUCT_SPEC.md
```

before implementation continues.

---

## ARCHITECTURE_CHANGES

If implementation requires architecture changes:

1. Update ARCHITECTURE.md
2. Update implementation

Never allow architecture and code to diverge.

---

## FRAMEWORK_CHANGES

If implementation requires framework changes:

1. Update FRAMEWORK.md
2. Update implementation

Never allow framework and code to diverge.

---

## TASK_CHANGES

If implementation invalidates a task:

Update TASKS.md.

---

## DESIGN_RULES

Prefer:

* simple
* explicit
* minimal

Avoid:

* abstraction without use
* premature optimization
* speculative features

---

## FILE_CREATION

Do not create new top-level documents unless explicitly requested.

Prefer modifying existing files.

---

## SUCCESS_CRITERIA

Task is complete only when:

* product requirements satisfied
* architectural requirements satisfied
* framework constraints satisfied
* acceptance criteria satisfied
* code runs
* tests pass

Then stop.

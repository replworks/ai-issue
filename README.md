# AI Issue Publisher

> Stop losing good AI ideas in chat history.
>
> Turn AI conversations into GitHub Issues with one command.

AI Issue Publisher converts AI-generated markdown from ChatGPT, Claude, Cursor, and other assistants into GitHub Issues.

The AI writes the idea. Humans decide whether it becomes work.

---

## Installation

### Homebrew (Recommended)

```bash
brew install replworks/tap/ai-issue
```

### Go

```bash
go install github.com/replworks/ai-issue/cmd/ai-issue@latest
```

Verify installation:

```bash
ai-issue diagnose
```

---

## Configuration

Create a GitHub Fine-grained Personal Access Token with:

```text
Issues: Read and write
```

Then configure:

```bash
export GITHUB_TOKEN=github_pat_xxxxxxxxxxxxxxxxx
```

Optional:

```bash
export AI_ISSUE_PUBLISHER=replworks-bot
```

---

## Usage

Copy AI-generated markdown to your clipboard and run:

```bash
ai-issue
```

Preview only:

```bash
ai-issue --dry-run
```

Diagnostics:

```bash
ai-issue diagnose
```

---

## Example

### AI Output

```markdown
# Add timestamps to logging system

Current logs do not contain timestamps, making debugging difficult.

Acceptance Criteria

- Include UTC timestamps
- Preserve current log format
- Add tests
```

### Publish

```bash
ai-issue
```

### Result

```text
✅ Issue created successfully!
https://github.com/owner/repository/issues/42
```

---

## Core Philosophy

### Author ≠ Publisher

AI Issue Publisher is built around a simple principle:

- AI authors the content.
- Humans review the content.
- Humans decide whether to publish it.

Publishing is always an explicit human decision.

### Dedicated AI Identity

Issues are created under a dedicated AI identity.

By default:

```text
@ai-backlog-bot
```

Override locally:

```bash
export AI_ISSUE_PUBLISHER=replworks-bot
```

This makes AI-generated issues immediately identifiable while preserving human accountability.

---

## Troubleshooting

### GITHUB_TOKEN is required

```bash
export GITHUB_TOKEN=github_pat_xxxxxxxxxxxxxxxxx
```

### Repository could not be determined

Run the command inside a Git repository.

### Clipboard is empty

Copy AI-generated markdown before running the command.

### Resource not accessible by personal access token

Verify:

- Fine-grained PAT is being used
- `Issues: Read and write` permission is granted
- Repository access is granted
- Organization approval is completed (if required)
- Token expiration is 366 days or less

### Need more details?

```bash
ai-issue diagnose
```

---

## Development

Run tests:

```bash
go test ./...
```

Build:

```bash
make build
```

Release:

```bash
goreleaser release --clean
```

---

## License

MIT

---

Built for developers who use AI every day and want good ideas to reach the backlog.

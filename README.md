# AI Issue Publisher

> Stop losing good AI ideas in chat history. Save them to GitHub with one command.
>
> Turn valuable AI conversations into GitHub Issues — with clear ownership and accountability.

AI Issue Publisher is a CLI tool that converts AI-generated ideas from ChatGPT, Claude, Cursor, and other AI assistants into GitHub Issues.

Instead of losing useful suggestions in chat history, you can quickly save them to your project backlog while clearly labeling them as **AI-generated drafts**.

The AI creates the content. The human decides whether it should become work.

---

## Why AI Issue Publisher?

AI tools generate a large number of potentially useful ideas:

* Feature suggestions
* Bug reports
* Refactoring proposals
* Product improvements
* Technical debt items

Most of these ideas disappear in chat history and never reach the project backlog.

AI Issue Publisher bridges that gap by providing a fast workflow for publishing AI-generated content directly to GitHub while preserving clear responsibility boundaries.

---

## Core Philosophy

### Author ≠ Publisher

AI Issue Publisher is built around a simple principle:

* AI **authors** the content.
* Humans **review** the content.
* Humans **choose** whether to publish it.

Publishing is an explicit decision made by a person.

This distinction helps teams avoid confusion about ownership and ensures AI-generated work is always identifiable.

### Dedicated AI Identity

Issues are created using a dedicated GitHub account (for example, `ai-backlog-bot`).

This makes it immediately obvious that:

* The issue originated from AI output.
* A human approved publication.
* The content should be treated as a draft until reviewed.

---

## Features

* 📋 Automatic clipboard reading
* 📂 Automatic Git repository detection
* 📝 Smart title extraction from Markdown (`# Heading`)
* 👀 Preview before publishing
* ✅ Explicit confirmation step
* 🤖 Dedicated AI bot account support
* ⚡ Fast workflow (typically under one minute)

---

## Installation

### Option 1: Install via Go

```bash
go install github.com/replworks/ai-issue/cmd/ai-issue@latest
```

### Option 2: Build from Source

```bash
git clone https://github.com/replworks/ai-issue.git

cd ai-issue

go mod tidy

go build -o ai-issue ./cmd/ai-issue
```

Optional global installation:

```bash
sudo mv ai-issue /usr/local/bin/
```

Verify installation:

```bash
ai-issue --help
ai-issue diagnose
```

---

## Configuration

### Create a Dedicated Bot Account

Recommended:

1. Create a GitHub account such as `ai-backlog-bot`.
2. Generate a Personal Access Token (PAT).
3. Grant repository access.
4. Configure the token locally.

### Set GitHub Token

```bash
export GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

For persistence:

**Zsh**

```bash
echo 'export GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx' >> ~/.zshrc
```

**Bash**

```bash
echo 'export GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx' >> ~/.bashrc
```

---

## Usage

### Typical Workflow

1. Copy an AI response to your clipboard.
2. Navigate to your Git repository.
3. Run:

```bash
ai-issue
```

1. Review the generated preview.
2. Confirm publication.

That's it.

---

## Commands

```bash
ai-issue
```

Publish the markdown from your clipboard as a new GitHub Issue.

```bash
ai-issue diagnose
```

Check Git availability, repository resolution, publisher validation, clipboard availability, and token configuration.

```bash
ai-issue --help
```

Display command help.

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

### Command

```bash
ai-issue
```

### Preview

```text
Repository: yourname/project

Title: Add timestamps to logging system

Body preview:
Current logs do not contain timestamps, making debugging difficult.

**Publisher:** @ai-backlog-bot

This issue will be created by ai-backlog-bot account.

Create Issue? (Y/n):
```

### Result

```text
✅ Issue created successfully!
https://github.com/yourname/project/issues/42
```

---

## Troubleshooting

### GITHUB_TOKEN is required

Configure the environment variable:

```bash
export GITHUB_TOKEN=...
```

### Repository could not be determined

Run the command inside a Git repository.

### Permission denied

Verify that the bot account has access to the target repository.

### Clipboard is empty

Copy the AI-generated content before running the command.

### Need more details?

Run:

```bash
ai-issue diagnose
```

---

## Project Structure

```text
ai-issue/
├── cmd/
│   └── ai-issue/
├── internal/
│   ├── adapter/
│   ├── cli/
│   ├── compliance/
│   ├── construction/
│   ├── domain/
│   ├── extraction/
│   ├── preview/
│   ├── publisher/
│   └── repository/
├── .goreleaser.yml
├── ARCHITECTURE.md
├── FRAMEWORK.md
├── LICENSE
├── PRODUCT_SPEC.md
├── README.md
└── TASKS.md
```

---

## Development

### Run Tests

```bash
go test ./...
```

### Release

This project uses GoReleaser for GitHub Releases.

Install:

```bash
go install github.com/goreleaser/goreleaser@latest
```

Snapshot release:

```bash
goreleaser release --snapshot --clean
```

Production release:

```bash
goreleaser release
```

### Supported Platforms

* macOS (arm64, amd64)
* Linux (arm64, amd64)
* Windows (amd64)

---

## License

MIT License

---

Built for AI-native developers who believe that AI can generate ideas, but humans remain responsible for deciding what gets shipped.

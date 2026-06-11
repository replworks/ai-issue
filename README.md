# AI Issue Publisher

**A CLI tool that registers AI-generated ideas as GitHub Issues while keeping them distinct from human accountability.**

A tool designed to help you **quickly and easily** save great ideas from conversations with AI (ChatGPT, Claude, Cursor, etc.) into your GitHub backlog, with a clear indication that "this is an AI-generated draft."

---

## Core Philosophy

- **Author ≠ Publisher**
- AI **generates** the content; the human **decides** whether to register it.
- Issues appear as `opened by ai-backlog-bot` in GitHub, allowing you to identify AI-generated content at a glance.

---

## Features

- Automatic clipboard reading (Cmd+C → Run immediately)
- Automatic Git repository detection
- Automatic Markdown title extraction
- Preview + Confirmation step
- Issue creation via a dedicated AI Bot account
- Complete in under 1 minute

---

## Installation Guide

### 1. Prepare a GitHub Bot Account (Required)

1. Create a GitHub account named `ai-backlog-bot` (or a name of your choice).
2. Generate a **Personal Access Token** for this account.
   - Required Scope: `repo` (full repository access)
   - Both Classic Tokens and Fine-grained Tokens are supported.
3. Register the token as an environment variable in your terminal:

`export GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`

> **Tip:** Add this line to your `~/.zshrc` or `~/.bashrc` to make it persistent.

### 2. Download and Install

```bash
# 1. Download and unzip the file
unzip ai-issue-updated.zip -d ai-issue
cd ai-issue

# 2. Install dependencies
go mod tidy

# 3. Build the binary
go build -o ai-issue ./cmd/ai-issue

# 4. Install globally (Recommended)
sudo mv ai-issue /usr/local/bin/
```

Verify the installation:

```bash
ai-issue --help
ai-issue diagnose
```

## Usage

### Typical Workflow

1. Ask your AI question and copy the entire response (Cmd + C).
2. Run the command in your terminal:

```bash
ai-issue
```

1. Review the preview and press `Y` or `Enter` to create the issue.

### Available Commands

```bash
ai-issue                # Publish AI Issue (default)
ai-issue diagnose       # Verify settings, Token, and Git Repository
ai-issue --help         # Show help menu
```

## Example

Copying an AI response:

```bash
# Add timestamps to logging system

Current logs do not contain timestamps, making it hard to debug...
```

Execution result:

```bash
Repository: yourname/project
Title: Add timestamps to logging system
Author: ai-backlog-bot

Create this issue? (Y/n) Y

✅ Issue created successfully!
URL: https://github.com/yourname/project/issues/42
```

## Troubleshooting

- "GITHUB_TOKEN is required" → Run `export GITHUB_TOKEN=...`
- Repository not found → Ensure you are running the command inside a Git repository folder.
- Permission denied → Ensure the Bot account has permission to access the target repository.
- Clipboard empty → Make sure you have copied the AI response to your clipboard first.

Try running `ai-issue diagnose` to troubleshoot your environment.

## Project Structure

```text
ai-issue/
├── cmd/ai-issue/
├── internal/
│   ├── adapter/
│   ├── cli/
│   ├── domain/
│   ├── extraction/
│   ├── preview/
│   ├── publisher/
│   └── repository/
├── PRODUCT_SPEC.md
├── ARCHITECTURE.md
├── FRAMEWORK.md
├── README.md
└── go.mod
```

## Running Tests

```bash
go test ./internal/... -v
```

## License

MIT License

# AI Issue Publisher

**A CLI tool to publish AI-generated ideas as GitHub Issues, clearly separating AI authorship from human responsibility.**

This tool allows you to **effortlessly save valuable ideas** generated during conversations with AI (ChatGPT, Claude, Cursor, etc.) to your GitHub backlog while making it explicitly clear that "this is an AI-generated draft."

---

## Core Philosophy

- **Author ≠ Publisher**
- AI **generates the content**, while humans **decide whether to publish it**.
- Issues are displayed as `opened by ai-backlog-bot` on GitHub, allowing everyone to identify AI-generated issues at a glance.

---

## Features

- **Automatic Clipboard Reading** — Simply copy the AI response and run `ai-issue` immediately.
- **Automatic Git Repository Detection**
- **Automatic Markdown Title Extraction** (Uses the first `# Heading`).
- **Preview & Confirmation Step** to prevent accidental creation.
- **Dedicated AI Bot Account** used for issue creation.
- **Blazing Fast** — Complete the entire process in 30 seconds to 1 minute.

---

## Installation

### 1. Prepare GitHub Token (Required)

Generate a **Personal Access Token** (with `repo` scope) for your `ai-backlog-bot` account.

```bash
export GITHUB_TOKEN=ghp_xxxxxxxxxxxxxxxxxxxx
```

### 2. Setup

```bash
# Extract the downloaded zip file
unzip ai-issue.zip
cd ai-issue

# Build the binary
go build -o ai-issue ./cmd/ai-issue

# Install globally (Optional)
sudo mv ai-issue /usr/local/bin/
```

---

## Usage

### Basic Flow

1. **Copy** the AI response to your clipboard (Cmd + C / Ctrl + C).
2. Run the command in your terminal:

```bash
ai-issue
```

1. Review the preview and press `Y` → Issue created successfully!

### Other Commands

```bash
ai-issue diagnose     # Validate configuration and permissions
ai-issue --help       # Show help menu
```

---

## Example

**After copying an AI response like this:**

```markdown
# Add timestamps to logging system

Current logs do not contain timestamps...
```

**Execution Result:**

```bash
Repository: company/backend
Title: Add timestamps to logging system
Author: ai-backlog-bot

Create this issue? (Y/n) Y

✅ Issue created: [https://github.com/company/backend/issues/123](https://github.com/company/backend/issues/123)
```

---

## Project Structure

```text
ai-issue/
├── cmd/ai-issue/
├── internal/
│   ├── adapter/
│   ├── cli/
│   ├── domain/
│   ├── extraction/
│   ├── publisher/
│   └── repository/
├── PRODUCT_SPEC.md
├── ARCHITECTURE.md
├── FRAMEWORK.md
└── README.md
```

---

## Documentation

- **PRODUCT_SPEC.md** — Product Requirements
- **ARCHITECTURE.md** — Architecture & Invariants
- **FRAMEWORK.md** — Implementation Rules

---

## License

MIT License

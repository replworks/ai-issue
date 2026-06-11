package repository

import (
	"os/exec"
	"strings"

	"ai-issue/internal/extraction"
)

func ResolveRepository() (string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	output, err := cmd.Output()
	if err != nil {
		return "", NewError("repository", "Not a git repository or no origin remote found.")
	}

	repoURL := strings.TrimSpace(string(output))
	// Simple extraction, e.g., git@github.com:owner/repo.git -> owner/repo
	repo := extractRepoName(repoURL)
	return repo, nil
}

func extractRepoName(url string) string {
	repo := strings.TrimSpace(url)
	repo = strings.TrimSuffix(repo, ".git")

	if strings.HasPrefix(repo, "git@") {
		if idx := strings.Index(repo, ":"); idx != -1 {
			return repo[idx+1:]
		}
	}

	if strings.HasPrefix(repo, "https://") || strings.HasPrefix(repo, "http://") {
		parts := strings.Split(repo, "/")
		if len(parts) >= 2 {
			return strings.Join(parts[len(parts)-2:], "/")
		}
	}

	if idx := strings.LastIndex(repo, ":"); idx != -1 {
		return repo[idx+1:]
	}

	return repo
}

func NewError(kind, msg string) error {
	return extraction.NewError(kind, msg) // reuse
}

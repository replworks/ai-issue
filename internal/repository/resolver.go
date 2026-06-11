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
	url = strings.TrimSpace(url)
	url = strings.TrimSuffix(url, ".git")

	// HTTPS
	if strings.Contains(url, "github.com/") {
		parts := strings.Split(url, "github.com/")
		if len(parts) > 1 {
			return parts[len(parts)-1]
		}
	}

	// SSH
	if idx := strings.LastIndex(url, ":"); idx != -1 {
		return url[idx+1:]
	}

	return url
}

func NewError(kind, msg string) error {
	return extraction.NewError(kind, msg) // reuse
}

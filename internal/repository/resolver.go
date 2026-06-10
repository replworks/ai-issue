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
	// Very basic extraction for MVP
	if idx := strings.LastIndex(url, ":"); idx != -1 {
		repo := url[idx+1:]
		repo = strings.TrimSuffix(repo, ".git")
		return repo
	}
	return url
}

func NewError(kind, msg string) error {
	return extraction.NewError(kind, msg) // reuse
}

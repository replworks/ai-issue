package repository

import (
	"os/exec"
	"strings"

	"github.com/replworks/ai-issue/internal/extraction"
)

func ResolveRepository() (string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	output, err := cmd.Output()
	if err != nil {
		return "", NewError("repository", "Not a git repository or no origin remote found.")
	}

	repoURL := strings.TrimSpace(string(output))
	repo := extractRepoName(repoURL)
	if err := ValidateRepository(repo); err != nil {
		return "", err
	}
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

func ValidateRepository(repo string) error {
	repo = strings.TrimSpace(repo)
	if repo == "" {
		return NewError("repository", "Repository could not be determined.")
	}
	parts := strings.Split(repo, "/")
	if len(parts) != 2 || strings.TrimSpace(parts[0]) == "" || strings.TrimSpace(parts[1]) == "" {
		return NewError("repository", "Repository could not be determined.")
	}
	return nil
}

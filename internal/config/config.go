package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const defaultPublisher = "ai-backlog-bot"
const defaultGitHubAppClientID = "Iv23lisWVphLHQXlTQql"

func PublisherIdentity() string {
	return normalizePublisher(os.Getenv("AI_ISSUE_PUBLISHER"))
}

func GitHubAppClientID() string {
	if v := strings.TrimSpace(os.Getenv("AI_ISSUE_GITHUB_APP_CLIENT_ID")); v != "" {
		return v
	}
	return defaultGitHubAppClientID
}

func TokenPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("failed to determine config directory: %w", err)
	}
	return filepath.Join(dir, "ai-issue", "token"), nil
}

func LoadToken() (string, error) {
	path, err := TokenPath()
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("GitHub App token not found. Run `ai-issue login` first")
		}
		return "", fmt.Errorf("failed to read GitHub App token: %w", err)
	}
	token := strings.TrimSpace(string(data))
	if token == "" {
		return "", fmt.Errorf("GitHub App token not found. Run `ai-issue login` first")
	}
	return token, nil
}

func SaveToken(token string) error {
	token = strings.TrimSpace(token)
	if token == "" {
		return fmt.Errorf("token is required")
	}
	path, err := TokenPath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}
	if err := os.WriteFile(path, []byte(token+"\n"), 0o600); err != nil {
		return fmt.Errorf("failed to store GitHub App token: %w", err)
	}
	return nil
}

func normalizePublisher(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return defaultPublisher
	}
	return strings.TrimPrefix(value, "@")
}

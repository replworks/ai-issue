package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/replworks/ai-issue/internal/extraction"
)

type Client struct {
	Token      string
	HTTPClient *http.Client
	BaseURL    string
}

type APIError struct {
	Message string `json:"message"`
}

func NewClient() (*Client, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, extraction.NewError("auth", "GITHUB_TOKEN environment variable is required. Set it with your publisher token.")
	}
	return &Client{
		Token:      token,
		HTTPClient: &http.Client{},
		BaseURL:    "https://api.github.com",
	}, nil
}

func (c *Client) CreateIssue(repo, title, body string) (string, error) {
	if c == nil {
		return "", fmt.Errorf("GitHub client is not initialized")
	}
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{}
	}
	if c.BaseURL == "" {
		c.BaseURL = "https://api.github.com"
	}

	url := fmt.Sprintf("%s/repos/%s/issues", c.BaseURL, repo)

	payload := map[string]interface{}{
		"title": title,
		"body":  body,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to encode GitHub issue payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create GitHub issue request: %w", err)
	}
	req.Header.Set("Authorization", "token "+c.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to publish GitHub issue: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	responseBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("GitHub issue publication failed\nRepository: %s\nStatus: %s\nReason:\n%s", repo, resp.Status, githubErrorMessage(responseBody))
	}

	// In real implementation, parse response for HTML URL
	return fmt.Sprintf("https://github.com/%s/issues", repo), nil
}

func (c *Client) CheckRepositoryAccess(repo string) error {
	if c == nil {
		return fmt.Errorf("GitHub client is not initialized")
	}
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{}
	}
	if c.BaseURL == "" {
		c.BaseURL = "https://api.github.com"
	}

	url := fmt.Sprintf("%s/repos/%s", c.BaseURL, repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create GitHub repository access request: %w", err)
	}
	req.Header.Set("Authorization", "token "+c.Token)
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to verify repository access: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return &RepositoryAccessError{
			Status:  http.StatusText(resp.StatusCode),
			Message: githubErrorMessage(body),
		}
	}

	return nil
}

type RepositoryAccessError struct {
	Status  string
	Message string
}

func (e *RepositoryAccessError) Error() string {
	if e == nil {
		return "repository access validation failed"
	}
	return fmt.Sprintf("repository access validation failed: %s", e.Status)
}

func githubErrorMessage(body []byte) string {
	var payload APIError
	if err := json.Unmarshal(body, &payload); err == nil && strings.TrimSpace(payload.Message) != "" {
		return payload.Message
	}

	trimmed := strings.TrimSpace(string(body))
	if trimmed == "" {
		return "Unknown GitHub API error."
	}
	return trimmed
}

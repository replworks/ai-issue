package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"ai-issue/internal/extraction"
)

type Client struct {
	Token      string
	HTTPClient *http.Client
	BaseURL    string
}

func NewClient() (*Client, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, extraction.NewError("auth", "GITHUB_TOKEN environment variable is required. Set it with your ai-backlog-bot token.")
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

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("GitHub issue publication failed: %s", resp.Status)
	}

	// In real implementation, parse response for HTML URL
	return fmt.Sprintf("https://github.com/%s/issues", repo), nil
}

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
	token string
}

func NewClient() (*Client, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, extraction.NewError("auth", "GITHUB_TOKEN environment variable is required. Set it with your ai-backlog-bot token.")
	}
	return &Client{token: token}, nil
}

func (c *Client) CreateIssue(repo, title, body string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/issues", repo)

	payload := map[string]interface{}{
		"title": title,
		"body":  body,
	}

	jsonData, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "token "+c.token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	// In real implementation, parse response for HTML URL
	return fmt.Sprintf("https://github.com/%s/issues", repo), nil
}

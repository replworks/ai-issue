package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
		return nil, extraction.NewError("auth", "GITHUB_TOKEN environment variable is required.")
	}
	return &Client{token: token}, nil
}

func (c *Client) CreateIssue(repo, title, body string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/issues", repo)

	payload := map[string]interface{}{
		"title": title,
		"body":  body,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "token "+c.token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("GitHub API error (%d): %s", resp.StatusCode, string(bodyBytes))
	}

	// Parse response to get real HTML URL
	var result struct {
		HTMLURL string `json:"html_url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		// fallback
		return fmt.Sprintf("https://github.com/%s/issues", repo), nil
	}

	return result.HTMLURL, nil
}
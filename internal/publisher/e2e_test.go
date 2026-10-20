package publisher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/replworks/ai-issue/internal/adapter/github"
	"github.com/replworks/ai-issue/internal/construction"
	"github.com/replworks/ai-issue/internal/domain"
)

// TestE2ELivePublicationPreservesAISource exercises the real GitHub issue
// creation boundary. It is opt-in because it creates a real issue.
// Run with: AI_ISSUE_E2E_PUBLISH=1 AI_ISSUE_E2E_REPO=owner/repo go test ./internal/publisher -run TestE2ELivePublicationPreservesAISource -v
func TestE2ELivePublicationPreservesAISource(t *testing.T) {
	if os.Getenv("AI_ISSUE_E2E_PUBLISH") == "" {
		t.Skip("set AI_ISSUE_E2E_PUBLISH=1 to create a live GitHub issue")
	}
	client, err := github.NewClient()
	if err != nil {
		t.Skipf("no local token: %v", err)
	}
	repo := os.Getenv("AI_ISSUE_E2E_REPO")
	if repo == "" {
		repo = "replworks/ai-issue"
	}

	draft := &domain.IssueDraft{
		Title: "AI Issue live E2E " + time.Now().UTC().Format("20060102-150405"),
		Body:  "This issue verifies AI source attribution at the GitHub API boundary.",
	}
	issue, err := construction.BuildPublishableIssue(draft, repo, "Claude")
	if err != nil {
		t.Fatalf("build issue: %v", err)
	}
	url, err := NewService(client).Publish(issue)
	if err != nil {
		t.Fatalf("create live issue: %v", err)
	}

	issuePath := strings.TrimPrefix(strings.TrimPrefix(url, "https://github.com/"), "http://github.com/")
	parts := strings.Split(strings.Trim(issuePath, "/"), "/")
	if len(parts) != 4 || parts[2] != "issues" {
		t.Fatalf("unexpected issue URL: %s", url)
	}
	req, err := http.NewRequest(http.MethodGet, client.BaseURL+"/repos/"+parts[0]+"/"+parts[1]+"/issues/"+parts[3], nil)
	if err != nil {
		t.Fatalf("create verification request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+client.Token)
	req.Header.Set("Accept", "application/vnd.github+json")
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		t.Fatalf("get live issue: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("get live issue status = %s", resp.Status)
	}
	var got struct {
		Body string `json:"body"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("decode live issue: %v", err)
	}
	if !strings.Contains(got.Body, "AI Source: Claude") {
		t.Fatalf("live issue body = %q, want AI source metadata", got.Body)
	}
	fmt.Println("created live issue:", url)
}

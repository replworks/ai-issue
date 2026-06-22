package publisher

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/replworks/ai-issue/internal/adapter/github"
	"github.com/replworks/ai-issue/internal/construction"
	"github.com/replworks/ai-issue/internal/domain"
)

func TestEndToEndPublicationFlow(t *testing.T) {
	var gotBody map[string]interface{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
			t.Fatalf("failed to decode body: %v", err)
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	ghClient := &github.Client{
		Token:      "token-123",
		HTTPClient: server.Client(),
		BaseURL:    server.URL,
	}

	draft := &domain.IssueDraft{
		Title: "Add timestamps to logging",
		Body:  "Current logs do not contain timestamps.",
	}

	issue, err := construction.BuildPublishableIssue(draft, "company/backend", "ai-backlog-bot")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	svc := NewService(ghClient, "ai-backlog-bot")
	url, err := svc.Publish(issue)
	if err != nil {
		t.Fatalf("unexpected publish error: %v", err)
	}
	if !strings.Contains(url, "company/backend") {
		t.Fatalf("url = %q, want repository path", url)
	}
	if body, _ := gotBody["body"].(string); strings.Contains(body, "AI Generated") || strings.Contains(body, "Publisher:") {
		t.Fatalf("body = %q, want clean issue body", body)
	}
	labels, _ := gotBody["labels"].([]interface{})
	if len(labels) != 1 || labels[0] != "ai-generated" {
		t.Fatalf("labels = %v, want [ai-generated]", labels)
	}
}

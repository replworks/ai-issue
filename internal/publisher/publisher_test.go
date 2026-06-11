package publisher

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ai-issue/internal/adapter/github"
	"ai-issue/internal/construction"
	"ai-issue/internal/domain"
)

func TestEndToEndPublicationFlow(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
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
}

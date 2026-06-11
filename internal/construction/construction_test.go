package construction

import (
	"strings"
	"testing"

	"github.com/replworks/ai-issue/internal/domain"
)

func TestBuildPublishableIssue(t *testing.T) {
	draft := &domain.IssueDraft{
		Title: "Add timestamps to logging",
		Body:  "Current logs do not contain timestamps.",
	}

	issue, err := BuildPublishableIssue(draft, "company/backend", "ai-backlog-bot")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if issue == nil {
		t.Fatal("expected issue, got nil")
	}
	if issue.Title != draft.Title {
		t.Fatalf("title = %q, want %q", issue.Title, draft.Title)
	}
	if issue.Repository != "company/backend" {
		t.Fatalf("repository = %q, want %q", issue.Repository, "company/backend")
	}
	if issue.Publisher != "ai-backlog-bot" {
		t.Fatalf("publisher = %q, want %q", issue.Publisher, "ai-backlog-bot")
	}
	if issue.Body == draft.Body {
		t.Fatal("expected publisher traceability to be preserved in body")
	}
}

func TestBuildPublishableIssueTrimsPublisherPrefix(t *testing.T) {
	draft := &domain.IssueDraft{
		Title: "Add timestamps to logging",
		Body:  "Current logs do not contain timestamps.",
	}

	issue, err := BuildPublishableIssue(draft, "company/backend", "@project-ai-bot")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if issue.Publisher != "project-ai-bot" {
		t.Fatalf("publisher = %q, want %q", issue.Publisher, "project-ai-bot")
	}
	if !contains(issue.Body, "**Publisher:** @project-ai-bot") {
		t.Fatalf("body = %q, want publisher footer", issue.Body)
	}
}

func TestValidatePublishableIssue(t *testing.T) {
	tests := []struct {
		name    string
		issue   *domain.PublishableIssue
		wantErr bool
	}{
		{
			name: "valid",
			issue: &domain.PublishableIssue{
				Title:      "Issue",
				Body:       "Body",
				Repository: "company/backend",
				Publisher:  "ai-backlog-bot",
			},
			wantErr: false,
		},
		{name: "nil", issue: nil, wantErr: true},
		{
			name: "missing title",
			issue: &domain.PublishableIssue{
				Body:       "Body",
				Repository: "company/backend",
				Publisher:  "ai-backlog-bot",
			},
			wantErr: true,
		},
		{
			name: "missing body",
			issue: &domain.PublishableIssue{
				Title:      "Issue",
				Repository: "company/backend",
				Publisher:  "ai-backlog-bot",
			},
			wantErr: true,
		},
		{
			name: "missing repository",
			issue: &domain.PublishableIssue{
				Title:     "Issue",
				Body:      "Body",
				Publisher: "ai-backlog-bot",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePublishableIssue(tt.issue)
			if tt.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

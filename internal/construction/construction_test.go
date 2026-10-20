package construction

import (
	"testing"

	"github.com/replworks/ai-issue/internal/domain"
)

func TestBuildPublishableIssue(t *testing.T) {
	draft := &domain.IssueDraft{
		Title: "Add timestamps to logging",
		Body:  "Current logs do not contain timestamps.",
	}

	issue, err := BuildPublishableIssue(draft, "company/backend", "Claude")
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
	if issue.Body != draft.Body+"\n\nAI Source: Claude" {
		t.Fatalf("body = %q, want AI source metadata", issue.Body)
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
			},
			wantErr: false,
		},
		{name: "nil", issue: nil, wantErr: true},
		{
			name: "missing title",
			issue: &domain.PublishableIssue{
				Body:       "Body",
				Repository: "company/backend",
			},
			wantErr: true,
		},
		{
			name: "missing body",
			issue: &domain.PublishableIssue{
				Title:      "Issue",
				Repository: "company/backend",
			},
			wantErr: true,
		},
		{
			name: "missing repository",
			issue: &domain.PublishableIssue{
				Title: "Issue",
				Body:  "Body",
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

func TestBuildPublishableIssueRequiresAISource(t *testing.T) {
	_, err := BuildPublishableIssue(
		&domain.IssueDraft{Title: "Issue", Body: "Body"},
		"company/backend",
		"",
	)
	if err == nil {
		t.Fatal("expected missing AI source error")
	}
}

func TestBuildPublishableIssueReplacesExistingAISourceMetadata(t *testing.T) {
	issue, err := BuildPublishableIssue(
		&domain.IssueDraft{Title: "Issue", Body: "Body\n\nAI Source: Old"},
		"company/backend",
		"Claude",
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if issue.Body != "Body\n\nAI Source: Claude" {
		t.Fatalf("body = %q, want one current AI source line", issue.Body)
	}
}

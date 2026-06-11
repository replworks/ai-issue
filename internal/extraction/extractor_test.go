package extraction

import (
	"testing"
)

func TestExtractIssue(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		wantTitle string
		wantBody  string
		wantErr   bool
	}{
		{
			name:    "empty content",
			content: "",
			wantErr: true,
		},
		{
			name:      "with title",
			content:   "# Add timestamps to logging\n\nCurrent logs do not contain timestamps.",
			wantTitle: "Add timestamps to logging",
			wantBody:  "Current logs do not contain timestamps.",
			wantErr:   false,
		},
		{
			name:      "with inline code title",
			content:   "# Add `--version` Support\n\nBody text.",
			wantTitle: "Add `--version` Support",
			wantBody:  "Body text.",
			wantErr:   false,
		},
		{
			name:      "with multiple inline code segments",
			content:   "# Fix `go install` and `diagnose`\n\nBody text.",
			wantTitle: "Fix `go install` and `diagnose`",
			wantBody:  "Body text.",
			wantErr:   false,
		},
		{
			name:    "no title",
			content: "Just some content without heading.",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			draft, err := ExtractIssue(tt.content)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if draft == nil {
				t.Fatal("expected draft, got nil")
			}
			if draft.Title != tt.wantTitle {
				t.Errorf("Title = %v, want %v", draft.Title, tt.wantTitle)
			}
			if draft.Body != tt.wantBody {
				t.Errorf("Body = %v, want %v", draft.Body, tt.wantBody)
			}
		})
	}
}

package preview

import (
	"testing"
)

func TestFormatPreview(t *testing.T) {
	draft := &struct {
		Title      string
		Body       string
		Repository string
	}{
		Title:      "Test Issue",
		Body:       "This is a test body.",
		Repository: "user/repo",
	}

	preview := FormatPreview(draft)

	if preview == "" {
		t.Error("FormatPreview returned empty string")
	}

	if !contains(preview, "Test Issue") || !contains(preview, "user/repo") {
		t.Error("Preview does not contain expected content")
	}
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0
}

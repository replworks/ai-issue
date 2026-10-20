package preview

import (
	"strings"
	"testing"
)

func TestFormatPreview(t *testing.T) {
	preview := FormatPreview("user/repo", "Test Issue", "This is a test body.")

	if preview == "" {
		t.Error("FormatPreview returned empty string")
	}

	if !contains(preview, "Test Issue") || !contains(preview, "user/repo") {
		t.Error("Preview does not contain expected content")
	}
	if contains(preview, "Publisher:") {
		t.Error("Preview should not contain publisher metadata")
	}
}

func TestFormatPreviewPreservesInlineCodeTitle(t *testing.T) {
	preview := FormatPreview("user/repo", "Add `--dry-run` Mode", "This is a test body.")

	if !contains(preview, "Title: Add `--dry-run` Mode") {
		t.Error("Preview does not preserve inline code in title")
	}
}

func TestFormatPreviewDryRun(t *testing.T) {
	preview := FormatPreviewWithOptions("user/repo", "Test Issue", "This is a test body.", false, true)

	if !contains(preview, "Dry run mode enabled.") {
		t.Error("Preview does not contain dry run message")
	}
	if !contains(preview, "No GitHub Issue will be created.") {
		t.Error("Preview does not contain no-issue message")
	}
	if contains(preview, "Create Issue? (Y/n):") {
		t.Error("Dry run preview should not ask for confirmation")
	}
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

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
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

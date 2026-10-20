package preview

import (
	"fmt"
	"strings"
)

func FormatPreview(repo, title, body string) string {
	return FormatPreviewWithOptions(repo, title, body, true, false)
}

func FormatPreviewWithOptions(repo, title, body string, prompt bool, dryRun bool) string {
	var b strings.Builder
	b.WriteString("\n=== AI Issue Preview ===\n")
	fmt.Fprintf(&b, "Repository: %s\n", repo)
	fmt.Fprintf(&b, "Title: %s\n", title)
	b.WriteString("\nBody preview:\n")
	if len(body) > 300 {
		b.WriteString(body[:300] + "...\n")
	} else {
		b.WriteString(body + "\n")
	}
	if dryRun {
		b.WriteString("\nDry run mode enabled.\nNo GitHub Issue will be created.\n")
		return b.String()
	}
	if prompt {
		b.WriteString("\nCreate Issue? (Y/n): ")
	}
	return b.String()
}

func ShowPreview(repo, title, body string) {
	fmt.Print(FormatPreview(repo, title, body))
}

func ShowDryRunPreview(repo, title, body string) {
	fmt.Print(FormatPreviewWithOptions(repo, title, body, false, true))
}

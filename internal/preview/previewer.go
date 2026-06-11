package preview

import (
	"fmt"
	"strings"
)

func FormatPreview(repo, title, body string) string {
	var b strings.Builder
	b.WriteString("\n=== AI Issue Preview ===\n")
	b.WriteString(fmt.Sprintf("Repository: %s\n", repo))
	b.WriteString(fmt.Sprintf("Title: %s\n", title))
	b.WriteString("\nBody preview:\n")
	if len(body) > 300 {
		b.WriteString(body[:300] + "...\n")
	} else {
		b.WriteString(body + "\n")
	}
	b.WriteString("\nThis issue will be created by ai-backlog-bot account.\n")
	b.WriteString("\nCreate Issue? (Y/n): ")
	return b.String()
}

func ShowPreview(repo, title, body string) {
	fmt.Print(FormatPreview(repo, title, body))
}

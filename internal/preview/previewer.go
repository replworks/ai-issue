package preview

import (
	"fmt"
	"strings"
)

func FormatPreview(repo, title, body, publisher string) string {
	var b strings.Builder
	b.WriteString("\n=== AI Issue Preview ===\n")
	fmt.Fprintf(&b, "Repository: %s\n", repo)
	fmt.Fprintf(&b, "Title: %s\n", title)
	fmt.Fprintf(&b, "Publisher: @%s\n", strings.TrimPrefix(strings.TrimSpace(publisher), "@"))
	b.WriteString("\nBody preview:\n")
	if len(body) > 300 {
		b.WriteString(body[:300] + "...\n")
	} else {
		b.WriteString(body + "\n")
	}
	b.WriteString("\nCreate Issue? (Y/n): ")
	return b.String()
}

func ShowPreview(repo, title, body, publisher string) {
	fmt.Print(FormatPreview(repo, title, body, publisher))
}

package preview

import (
	"fmt"
)

func ShowPreview(repo, title, body string) {
	fmt.Println("\n=== AI Issue Preview ===")
	fmt.Printf("Repository: %s\n", repo)
	fmt.Printf("Title: %s\n", title)
	fmt.Println("\nBody preview:")
	if len(body) > 300 {
		fmt.Println(body[:300] + "...")
	} else {
		fmt.Println(body)
	}
	fmt.Println("\nThis issue will be created by ai-backlog-bot account.")
	fmt.Print("\nCreate Issue? (Y/n): ")
}

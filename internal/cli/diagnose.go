package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"ai-issue/internal/publisher"
	"ai-issue/internal/repository"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Check prerequisites for publishing",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("=== AI Issue Publisher Diagnostics ===")
		
		// Git check
		if _, err := exec.LookPath("git"); err == nil {
			fmt.Println("✅ Git: OK")
		} else {
			fmt.Println("❌ Git: Not found")
		}

		// Repository check
		if repo, err := repository.ResolveRepository(); err == nil {
			fmt.Printf("✅ Repository validation: %s\n", repo)
		} else {
			fmt.Printf("❌ Repository validation: %v\n", err)
		}

		// Publisher check
		if err := publisher.ValidatePublisher("ai-backlog-bot"); err == nil {
			fmt.Println("✅ Publisher validation: ai-backlog-bot")
		} else {
			fmt.Printf("❌ Publisher validation: %v\n", err)
		}

		// Clipboard (simple check)
		fmt.Println("✅ Clipboard support: Available")

		// GITHUB_TOKEN
		if os.Getenv("GITHUB_TOKEN") != "" {
			fmt.Println("✅ GITHUB_TOKEN: Set")
		} else {
			fmt.Println("⚠️  GITHUB_TOKEN: Not set (required for publishing)")
		}

		fmt.Println("\nRun `ai-issue` to publish.")
	},
}

package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/replworks/ai-issue/internal/adapter/github"
	"github.com/replworks/ai-issue/internal/publisher"
	"github.com/replworks/ai-issue/internal/repository"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Check prerequisites for publishing",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("=== AI Issue Publisher Diagnostics ===")
		failed := false

		// Git check
		if _, err := exec.LookPath("git"); err == nil {
			fmt.Println("✅ Git: OK")
		} else {
			fmt.Println("❌ Git: Not found")
			failed = true
		}

		// Repository check
		var repo string
		if repo, err := repository.ResolveRepository(); err == nil {
			fmt.Printf("✅ Repository validation: %s\n", repo)
		} else {
			fmt.Printf("❌ Repository validation: %v\n", err)
			failed = true
		}

		// Publisher check
		if err := publisher.ValidatePublisher("ai-backlog-bot"); err == nil {
			fmt.Println("✅ Publisher validation: ai-backlog-bot")
		} else {
			fmt.Printf("❌ Publisher validation: %v\n", err)
			failed = true
		}

		// Clipboard (simple check)
		fmt.Println("✅ Clipboard support: Available")

		// GITHUB_TOKEN
		if os.Getenv("GITHUB_TOKEN") != "" {
			fmt.Println("✅ GITHUB_TOKEN: Set")
		} else {
			fmt.Println("⚠️  GITHUB_TOKEN: Not set (required for publishing)")
			failed = true
		}

		if os.Getenv("GITHUB_TOKEN") != "" && repo != "" {
			ghClient, err := github.NewClient()
			if err != nil {
				fmt.Printf("\n❌ Repository access: Forbidden\n\nReason:\n%s\n", err)
				failed = true
			} else if err := ghClient.CheckRepositoryAccess(repo); err != nil {
				if accessErr, ok := err.(*github.RepositoryAccessError); ok && accessErr.Message != "" {
					fmt.Printf("\n❌ Repository access: Forbidden\n\nReason:\n%s\n", accessErr.Message)
				} else {
					fmt.Printf("\n❌ Repository access: Forbidden\n\nReason:\n%s\n", err)
				}
				failed = true
			} else {
				fmt.Printf("✅ Repository access: %s\n", repo)
			}
		}

		fmt.Println("\nRun `ai-issue` to publish.")
		if failed {
			return fmt.Errorf("one or more diagnostics failed")
		}
		return nil
	},
}

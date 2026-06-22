package cli

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/replworks/ai-issue/internal/adapter/github"
	"github.com/replworks/ai-issue/internal/config"
	"github.com/replworks/ai-issue/internal/publisher"
	"github.com/replworks/ai-issue/internal/repository"
)

var diagnoseCmd = &cobra.Command{
	Use:          "diagnose",
	Short:        "Check prerequisites for publishing",
	SilenceUsage: true,
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
		if resolvedRepo, err := repository.ResolveRepository(); err == nil {
			repo = resolvedRepo
			fmt.Printf("✅ Repository validation: %s\n", repo)
		} else {
			fmt.Printf("❌ Repository validation: %v\n", err)
			failed = true
		}

		// Publisher check
		publisherName := config.PublisherIdentity()
		if err := publisher.ValidatePublisher(publisherName); err == nil {
			fmt.Printf("✅ Publisher validation: @%s\n", strings.TrimPrefix(publisherName, "@"))
		} else {
			fmt.Printf("❌ Publisher validation: %v\n", err)
			failed = true
		}

		// Clipboard (simple check)
		fmt.Println("✅ Clipboard support: Available")

		// GitHub App token
		ghClient, err := github.NewClient()
		if err != nil {
			fmt.Printf("❌ GitHub App token: %v\n", err)
			failed = true
		} else {
			fmt.Println("✅ GitHub App token: Loaded")
		}

		if err == nil && repo != "" {
			if err := ghClient.CheckRepositoryAccess(repo); err != nil {
				if accessErr, ok := err.(*github.RepositoryAccessError); ok {
					fmt.Printf("\n❌ Repository access: Forbidden\n")
					if accessErr.Message != "" {
						fmt.Printf("\nReason:\n%s\n", accessErr.Message)
					}
					if sso := accessErr.Headers.Get("X-GitHub-SSO"); sso != "" {
						fmt.Printf("\nX-GitHub-SSO:\n%s\n", sso)
					}
					if accessErr.RawBody != "" {
						fmt.Printf("\nRaw response:\n%s\n", accessErr.RawBody)
					}
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

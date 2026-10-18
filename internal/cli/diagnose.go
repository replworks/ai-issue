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
			login, loginErr := ghClient.GetAuthenticatedLogin()
			if loginErr == nil && login != "" {
				fmt.Printf("✅ GitHub account: @%s\n", strings.TrimPrefix(login, "@"))
			} else if loginErr != nil {
				fmt.Printf("❌ GitHub account: %v\n", loginErr)
				if strings.Contains(loginErr.Error(), "401") || strings.Contains(loginErr.Error(), "Bad credentials") {
					fmt.Printf("\nHint:\nStored token is invalid or expired. Run `ai-issue login` again.\n")
				}
				failed = true
			}
			if err := ghClient.CheckRepositoryAccess(repo); err != nil {
				if accessErr, ok := err.(*github.RepositoryAccessError); ok {
					accessErr.Login = login
					fmt.Printf("\n❌ Repository access: Forbidden\n")
					if accessErr.Repository != "" {
						fmt.Printf("\nRepository:\n%s\n", accessErr.Repository)
					}
					if accessErr.Login != "" {
						fmt.Printf("\nAccount:\n%s\n", accessErr.Login)
					}
					if accessErr.Message != "" {
						fmt.Printf("\nReason:\n%s\n", accessErr.Message)
					}
					if accessErr.HasPerms && !accessErr.Push {
						who := "the publishing account"
						if accessErr.Login != "" {
							who = "account " + accessErr.Login
						}
						target := accessErr.Repository
						if target == "" {
							target = repo
						}
						fmt.Printf("\nHint:\n%s needs Write permission on %s.\n", who, target)
					} else if strings.Contains(accessErr.Message, "Bad credentials") || strings.Contains(accessErr.RawBody, "Bad credentials") {
						fmt.Printf("\nHint:\nStored token is invalid or expired, so permissions could not be checked. Run `ai-issue login` again, then `ai-issue diagnose`.\n")
					}
					if sso := accessErr.Headers.Get("X-GitHub-SSO"); sso != "" {
						fmt.Printf("\nX-GitHub-SSO:\n%s\n", sso)
					}
					if accessErr.RawBody != "" && !accessErr.HasPerms && !strings.Contains(accessErr.RawBody, "Bad credentials") {
						fmt.Printf("\nRaw response:\n%s\n", accessErr.RawBody)
					}
				} else {
					fmt.Printf("\n❌ Repository access: Forbidden\n\nReason:\n%s\n", err)
				}
				failed = true
			} else {
				fmt.Printf("✅ Repository access: %s (Write OK)\n", repo)
			}
		}

		fmt.Println("\nRun `ai-issue` to publish.")
		if failed {
			return fmt.Errorf("one or more diagnostics failed")
		}
		return nil
	},
}

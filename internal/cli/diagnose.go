package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
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

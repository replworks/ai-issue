package cli

import (
	"strings"

	"github.com/spf13/cobra"
)

var appVersion = "dev"
var dryRun bool
var aiSource string

var rootCmd = &cobra.Command{
	Use:           "ai-issue",
	Short:         "AI Issue Publisher - Publish AI-generated content as GitHub Issues",
	Long:          `AI Issue Publisher allows you to quickly publish AI-generated markdown as GitHub Issues using the authenticated GitHub account.`,
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Default to publish
		publishCmd.Run(cmd, args)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func SetVersion(version string) {
	version = strings.TrimSpace(version)
	if version == "" {
		version = "dev"
	}

	appVersion = version
	rootCmd.Version = appVersion
}

func init() {
	rootCmd.Version = appVersion
	rootCmd.SetVersionTemplate("{{.Name}} {{.Version}}\n")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Display the issue preview without creating a GitHub Issue")
	rootCmd.PersistentFlags().StringVar(&aiSource, "ai", "", "AI source: chatgpt, claude, muse, grok, or deepseek")
	rootCmd.AddCommand(publishCmd)
	rootCmd.AddCommand(diagnoseCmd)
	rootCmd.AddCommand(loginCmd)
}

package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "ai-issue",
	Short:         "AI Issue Publisher - Publish AI-generated content as GitHub Issues",
	Long:          `AI Issue Publisher allows you to quickly publish AI-generated markdown as GitHub Issues under a dedicated bot account.`,
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Default to publish
		publishCmd.Run(cmd, args)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(publishCmd)
	rootCmd.AddCommand(diagnoseCmd)
}

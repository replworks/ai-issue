package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"github.com/replworks/ai-issue/internal/adapter/github"
	"github.com/replworks/ai-issue/internal/aisource"
	"github.com/replworks/ai-issue/internal/construction"
	"github.com/replworks/ai-issue/internal/extraction"
	"github.com/replworks/ai-issue/internal/preview"
	"github.com/replworks/ai-issue/internal/publisher"
	"github.com/replworks/ai-issue/internal/repository"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish AI-generated markdown as GitHub Issue",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runPublish(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func runPublish() error {
	// 1. Get content from clipboard
	content, err := clipboard.ReadAll()
	if err != nil || strings.TrimSpace(content) == "" {
		return extraction.NewError("content", "Clipboard is empty. Copy the AI response first.")
	}

	// 2. Select the AI source
	selectedSource, err := aisource.Select(aiSource, os.Stdin, os.Stdout)
	if err != nil {
		return err
	}

	// 3. Extract
	draft, err := extraction.ExtractIssue(content)
	if err != nil {
		return err
	}
	if err := extraction.ValidateIssueDraft(draft); err != nil {
		return err
	}

	// 4. Resolve repo
	repo, err := repository.ResolveRepository()
	if err != nil {
		return err
	}
	if err := repository.ValidateRepository(repo); err != nil {
		return err
	}

	publishable, err := construction.BuildPublishableIssue(draft, repo, selectedSource)
	if err != nil {
		return err
	}

	// 5. Preview
	if dryRun {
		preview.ShowDryRunPreview(repo, publishable.Title, publishable.Body)
		return nil
	}
	preview.ShowPreview(repo, publishable.Title, publishable.Body)

	// 6. Confirm
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input != "y" && input != "" {
		fmt.Println("Publication cancelled.")
		return nil
	}

	// 7. Publish
	ghClient, err := github.NewClient()
	if err != nil {
		return err
	}

	pubService := publisher.NewService(ghClient)

	url, err := pubService.Publish(publishable)
	if err != nil {
		return err
	}

	fmt.Printf("\n✅ Issue created successfully!\nURL: %s\n", url)
	return nil
}

package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"ai-issue/internal/domain"
	"ai-issue/internal/extraction"
	"ai-issue/internal/preview"
	"ai-issue/internal/publisher"
	"ai-issue/internal/repository"
	"ai-issue/internal/adapter/github"
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

	// 2. Extract
	draft, err := extraction.ExtractIssue(content)
	if err != nil {
		return err
	}
	if err := extraction.ValidateIssueDraft(draft); err != nil {
		return err
	}

	// 3. Resolve repo
	repo, err := repository.ResolveRepository()
	if err != nil {
		return err
	}
	if err := repository.ValidateRepository(repo); err != nil {
		return err
	}

	// 4. Preview
	preview.ShowPreview(repo, draft.Title, draft.Body)

	// 5. Confirm
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input != "y" && input != "" {
		fmt.Println("Publication cancelled.")
		return nil
	}

	// 6. Publish
	ghClient, err := github.NewClient()
	if err != nil {
		return err
	}

	pubService := publisher.NewService(ghClient, "ai-backlog-bot")
	if err := publisher.ValidatePublisher("ai-backlog-bot"); err != nil {
		return err
	}

	publishable := &domain.PublishableIssue{
		Title:      draft.Title,
		Body:       draft.Body,
		Repository: repo,
	}

	url, err := pubService.Publish(publishable)
	if err != nil {
		return err
	}

	fmt.Printf("\n✅ Issue created successfully!\nURL: %s\n", url)
	return nil
}

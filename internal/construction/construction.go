package construction

import (
	"strings"

	"github.com/replworks/ai-issue/internal/domain"
)

func BuildPublishableIssue(draft *domain.IssueDraft, repositoryName, aiSource string) (*domain.PublishableIssue, error) {
	aiSource = strings.TrimSpace(aiSource)
	if aiSource == "" {
		return nil, domainError("AI source could not be determined.")
	}
	issue := &domain.PublishableIssue{
		Title:      "",
		Body:       "",
		Repository: strings.TrimSpace(repositoryName),
	}

	if draft != nil {
		issue.Title = strings.TrimSpace(draft.Title)
		issue.Body = appendAISource(draft.Body, aiSource)
	}

	if err := ValidatePublishableIssue(issue); err != nil {
		return nil, err
	}

	return issue, nil
}

func appendAISource(body, aiSource string) string {
	lines := strings.Split(strings.TrimSpace(body), "\n")
	filtered := make([]string, 0, len(lines))
	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "AI Source:") {
			continue
		}
		filtered = append(filtered, line)
	}
	cleanBody := strings.TrimSpace(strings.Join(filtered, "\n"))
	return cleanBody + "\n\nAI Source: " + aiSource
}

func ValidatePublishableIssue(issue *domain.PublishableIssue) error {
	if issue == nil {
		return domainError("issue payload could not be created.")
	}
	if strings.TrimSpace(issue.Title) == "" {
		return domainError("issue title could not be determined.")
	}
	if strings.TrimSpace(issue.Body) == "" {
		return domainError("issue body could not be determined.")
	}
	if strings.TrimSpace(issue.Repository) == "" {
		return domainError("repository could not be determined.")
	}
	return nil
}

func domainError(message string) error {
	return &constructionError{message: message}
}

type constructionError struct {
	message string
}

func (e *constructionError) Error() string {
	return e.message
}

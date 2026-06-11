package construction

import (
	"strings"

	"ai-issue/internal/domain"
)

func BuildPublishableIssue(draft *domain.IssueDraft, repositoryName, publisherName string) (*domain.PublishableIssue, error) {
	issue := &domain.PublishableIssue{
		Title:      "",
		Body:       "",
		Repository: strings.TrimSpace(repositoryName),
		Publisher:  strings.TrimSpace(publisherName),
	}

	if draft != nil {
		issue.Title = strings.TrimSpace(draft.Title)
		issue.Body = strings.TrimSpace(draft.Body)
	}

	if strings.TrimSpace(publisherName) != "" {
		issue.Body = strings.TrimSpace(issue.Body + "\n\n**Publisher:** @" + strings.TrimSpace(publisherName))
	}

	if err := ValidatePublishableIssue(issue); err != nil {
		return nil, err
	}

	return issue, nil
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
	if strings.TrimSpace(issue.Publisher) == "" {
		return domainError("publisher could not be determined.")
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

package publisher

import (
	"fmt"
	"strings"

	"github.com/replworks/ai-issue/internal/adapter/github"
	"github.com/replworks/ai-issue/internal/domain"
)

type Service struct {
	ghClient *github.Client
	botName  string
}

func NewService(ghClient *github.Client, botName string) *Service {
	return &Service{ghClient: ghClient, botName: botName}
}

func (s *Service) Publish(issue *domain.PublishableIssue) (string, error) {
	if err := ValidatePublisher(s.botName); err != nil {
		return "", err
	}
	if issue == nil {
		return "", fmt.Errorf("publishable issue is required")
	}

	// Add AI marker
	fullBody := fmt.Sprintf("**AI Generated** • Published by @%s\n\n%s", s.botName, issue.Body)

	url, err := s.ghClient.CreateIssue(issue.Repository, issue.Title, fullBody)
	if err != nil {
		return "", err
	}

	return url, nil
}

func ValidatePublisher(username string) error {
	if strings.TrimSpace(strings.TrimPrefix(username, "@")) == "" {
		return fmt.Errorf("publisher information is required")
	}
	return nil
}

package publisher

import (
	"ai-issue/internal/adapter/github"
	"ai-issue/internal/domain"
	"fmt"
)

type Service struct {
	ghClient *github.Client
	botName  string
}

func NewService(ghClient *github.Client, botName string) *Service {
	return &Service{ghClient: ghClient, botName: botName}
}

func (s *Service) Publish(issue *domain.PublishableIssue) (string, error) {
	// Add AI marker
	fullBody := fmt.Sprintf("**AI Generated** • Published by @%s\n\n%s", s.botName, issue.Body)

	url, err := s.ghClient.CreateIssue(issue.Repository, issue.Title, fullBody)
	if err != nil {
		return "", err
	}

	return url, nil
}

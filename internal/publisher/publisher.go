package publisher

import (
	"fmt"

	"github.com/replworks/ai-issue/internal/adapter/github"
	"github.com/replworks/ai-issue/internal/domain"
)

type Service struct {
	ghClient *github.Client
}

func NewService(ghClient *github.Client) *Service {
	return &Service{ghClient: ghClient}
}

func (s *Service) Publish(issue *domain.PublishableIssue) (string, error) {
	if issue == nil {
		return "", fmt.Errorf("publishable issue is required")
	}

	url, err := s.ghClient.CreateIssue(issue.Repository, issue.Title, issue.Body, []string{"ai-generated"})
	if err != nil {
		return "", err
	}

	return url, nil
}

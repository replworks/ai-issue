package config

import (
	"os"
	"strings"
)

const defaultPublisher = "ai-backlog-bot"

func PublisherIdentity() string {
	return normalizePublisher(os.Getenv("AI_ISSUE_PUBLISHER"))
}

func normalizePublisher(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return defaultPublisher
	}
	return strings.TrimPrefix(value, "@")
}

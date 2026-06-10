package extraction

import (
	"bytes"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"ai-issue/internal/domain"
)

func ExtractIssue(content string) (*domain.IssueDraft, error) {
	if strings.TrimSpace(content) == "" {
		return nil, ErrEmptyContent
	}

	// Simple title extraction from first # heading
	title, body := extractTitleAndBody(content)

	return &domain.IssueDraft{
		Title: title,
		Body:  body,
	}, nil
}

func extractTitleAndBody(md string) (string, string) {
	lines := strings.Split(md, "\n")
	title := "Untitled AI Issue"
	bodyStart := 0

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			title = strings.TrimPrefix(line, "# ")
			bodyStart = i + 1
			break
		}
	}

	body := strings.Join(lines[bodyStart:], "\n")
	return title, strings.TrimSpace(body)
}

var ErrEmptyContent = NewError("content", "Clipboard is empty. Copy AI response first.")

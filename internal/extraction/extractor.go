package extraction

import (
	"strings"

	"ai-issue/internal/domain"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func ExtractIssue(content string) (*domain.IssueDraft, error) {
	if strings.TrimSpace(content) == "" {
		return nil, ErrEmptyContent
	}

	title, body := extractTitleAndBody(content)
	if strings.TrimSpace(title) == "" {
		return nil, NewError("draft", "Issue title could not be determined.")
	}

	return &domain.IssueDraft{
		Title: title,
		Body:  body,
	}, nil
}

func extractTitleAndBody(md string) (string, string) {
	source := text.NewReader([]byte(md))
	doc := goldmark.DefaultParser().Parse(source)

	title := ""
	body := md

	ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering || title != "" {
			return ast.WalkContinue, nil
		}
		heading, ok := node.(*ast.Heading)
		if !ok || heading.Level != 1 {
			return ast.WalkContinue, nil
		}

		var buf strings.Builder
		for c := heading.FirstChild(); c != nil; c = c.NextSibling() {
			if textNode, ok := c.(*ast.Text); ok {
				buf.WriteString(string(textNode.Segment.Value(source.Source())))
			}
		}
		title = strings.TrimSpace(buf.String())
		if title == "" {
			title = "Untitled AI Issue"
		}

		lines := strings.Split(md, "\n")
		for i, line := range lines {
			if strings.TrimSpace(line) == strings.TrimSpace("# "+title) {
				body = strings.TrimSpace(strings.Join(lines[i+1:], "\n"))
				break
			}
		}
		return ast.WalkStop, nil
	})

	if title == "" {
		title = "Untitled AI Issue"
		body = strings.TrimSpace(md)
	}

	return title, body
}

var ErrEmptyContent = NewError("content", "Clipboard is empty. Copy AI response first.")

package aisource

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Source struct {
	ID   string
	Name string
}

var supportedSources = []Source{
	{ID: "chatgpt", Name: "ChatGPT"},
	{ID: "claude", Name: "Claude"},
	{ID: "muse", Name: "Muse"},
	{ID: "grok", Name: "Grok"},
	{ID: "deepseek", Name: "DeepSeek"},
}

func Supported() []Source {
	return append([]Source(nil), supportedSources...)
}

func Resolve(value string) (string, error) {
	value = strings.ToLower(strings.TrimSpace(value))
	for _, source := range supportedSources {
		if value == source.ID || value == strings.ToLower(source.Name) {
			return source.Name, nil
		}
	}
	return "", fmt.Errorf("unsupported AI source %q; choose chatgpt, claude, muse, grok, or deepseek", value)
}

func Select(value string, in io.Reader, out io.Writer) (string, error) {
	if strings.TrimSpace(value) != "" {
		return Resolve(value)
	}

	if _, err := fmt.Fprintln(out, "Select the AI source:"); err != nil {
		return "", fmt.Errorf("failed to display AI source options: %w", err)
	}
	for i, source := range supportedSources {
		if _, err := fmt.Fprintf(out, "%d) %s\n", i+1, source.Name); err != nil {
			return "", fmt.Errorf("failed to display AI source options: %w", err)
		}
	}
	if _, err := fmt.Fprint(out, "Choice: "); err != nil {
		return "", fmt.Errorf("failed to display AI source prompt: %w", err)
	}

	var choice string
	if _, err := fmt.Fscanln(in, &choice); err != nil {
		return "", fmt.Errorf("AI source selection is required")
	}
	index, err := strconv.Atoi(strings.TrimSpace(choice))
	if err != nil || index < 1 || index > len(supportedSources) {
		return "", fmt.Errorf("invalid AI source choice %q; choose a number from 1 to %d", choice, len(supportedSources))
	}
	return supportedSources[index-1].Name, nil
}

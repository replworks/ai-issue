package aisource

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const maxSourceNameLength = 40

type Source struct {
	ID   string
	Name string
}

var supportedSources = []Source{
	{ID: "chatgpt", Name: "ChatGPT"},
	{ID: "claude", Name: "Claude"},
	{ID: "gemini", Name: "Gemini"},
	{ID: "muse", Name: "Muse"},
	{ID: "grok", Name: "Grok"},
	{ID: "deepseek", Name: "DeepSeek"},
}

func Supported() []Source {
	return append([]Source(nil), supportedSources...)
}

func Resolve(value string) (string, error) {
	name := strings.TrimSpace(value)
	if name == "" {
		return "", fmt.Errorf("AI source is required")
	}

	for _, source := range supportedSources {
		if strings.EqualFold(name, source.ID) || strings.EqualFold(name, source.Name) {
			return source.Name, nil
		}
	}

	if err := validateCustomName(name); err != nil {
		return "", err
	}
	return name, nil
}

func validateCustomName(name string) error {
	if utf8.RuneCountInString(name) > maxSourceNameLength {
		return fmt.Errorf("invalid AI source name %q; enter %d characters or fewer", name, maxSourceNameLength)
	}
	for _, r := range name {
		if unicode.IsControl(r) {
			return fmt.Errorf("invalid AI source name %q; enter a single line without control characters", name)
		}
	}
	return nil
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
	if _, err := fmt.Fprintf(out, "%d) Other\n", len(supportedSources)+1); err != nil {
		return "", fmt.Errorf("failed to display AI source options: %w", err)
	}
	if _, err := fmt.Fprint(out, "Choice: "); err != nil {
		return "", fmt.Errorf("failed to display AI source prompt: %w", err)
	}

	choice, err := readLine(in)
	if err != nil {
		return "", fmt.Errorf("AI source selection is required")
	}
	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(supportedSources)+1 {
		return "", fmt.Errorf("invalid AI source choice %q; choose a number from 1 to %d", choice, len(supportedSources)+1)
	}
	if index == len(supportedSources)+1 {
		if _, err := fmt.Fprint(out, "AI name: "); err != nil {
			return "", fmt.Errorf("failed to display AI source prompt: %w", err)
		}
		name, err := readLine(in)
		if err != nil {
			return "", fmt.Errorf("AI source name is required")
		}
		return Resolve(name)
	}
	return supportedSources[index-1].Name, nil
}

func readLine(reader io.Reader) (string, error) {
	var line strings.Builder
	buf := make([]byte, 1)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			if buf[0] == '\n' {
				return strings.TrimSpace(line.String()), nil
			}
			if buf[0] != '\r' {
				line.WriteByte(buf[0])
			}
		}
		if err != nil {
			value := strings.TrimSpace(line.String())
			if value == "" {
				return "", err
			}
			return value, nil
		}
	}
}

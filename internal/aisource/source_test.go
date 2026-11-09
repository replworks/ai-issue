package aisource

import (
	"strings"
	"testing"
)

func TestResolveSupportedSources(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  string
	}{
		{input: "chatgpt", want: "ChatGPT"},
		{input: "Claude", want: "Claude"},
		{input: "gemini", want: "Gemini"},
		{input: " muse ", want: "Muse"},
		{input: "grok", want: "Grok"},
		{input: "deepseek", want: "DeepSeek"},
	} {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Resolve(tt.input)
			if err != nil {
				t.Fatalf("Resolve(%q) returned error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Fatalf("Resolve(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestResolveCustomSource(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  string
	}{
		{input: "Perplexity", want: "Perplexity"},
		{input: "  Cursor  ", want: "Cursor"},
		{input: "Claude 3.5 Sonnet", want: "Claude 3.5 Sonnet"},
	} {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Resolve(tt.input)
			if err != nil {
				t.Fatalf("Resolve(%q) returned error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Fatalf("Resolve(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestResolveRejectsInvalidSource(t *testing.T) {
	for name, input := range map[string]string{
		"empty":     "   ",
		"multiline": "Perplexity\nGemini",
		"control":   "Perplexity\tGemini",
		"too long":  strings.Repeat("a", maxSourceNameLength+1),
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := Resolve(input); err == nil {
				t.Fatalf("Resolve(%q) expected error", input)
			}
		})
	}
}

func TestSelectInteractive(t *testing.T) {
	var out strings.Builder
	got, err := Select("", strings.NewReader("3\n"), &out)
	if err != nil {
		t.Fatalf("Select returned error: %v", err)
	}
	if got != "Gemini" {
		t.Fatalf("Select = %q, want Gemini", got)
	}
	if !strings.Contains(out.String(), "1) ChatGPT") ||
		!strings.Contains(out.String(), "3) Gemini") ||
		!strings.Contains(out.String(), "6) DeepSeek") ||
		!strings.Contains(out.String(), "7) Other") {
		t.Fatalf("selection prompt = %q, want all supported sources and Other", out.String())
	}
}

func TestSelectOther(t *testing.T) {
	var out strings.Builder
	got, err := Select("", strings.NewReader("7\nPerplexity\n"), &out)
	if err != nil {
		t.Fatalf("Select returned error: %v", err)
	}
	if got != "Perplexity" {
		t.Fatalf("Select = %q, want Perplexity", got)
	}
	if !strings.Contains(out.String(), "AI name: ") {
		t.Fatalf("selection prompt = %q, want custom name prompt", out.String())
	}
}

func TestSelectOtherCanonicalizesSupportedName(t *testing.T) {
	got, err := Select("", strings.NewReader("7\nclaude\n"), &strings.Builder{})
	if err != nil {
		t.Fatalf("Select returned error: %v", err)
	}
	if got != "Claude" {
		t.Fatalf("Select = %q, want Claude", got)
	}
}

func TestSelectRejectsInvalidChoice(t *testing.T) {
	if _, err := Select("", strings.NewReader("8\n"), &strings.Builder{}); err == nil {
		t.Fatal("expected invalid choice error")
	}
}

func TestSelectRequiresChoice(t *testing.T) {
	if _, err := Select("", strings.NewReader(""), &strings.Builder{}); err == nil {
		t.Fatal("expected required selection error")
	}
}

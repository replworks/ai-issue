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

func TestSelectInteractive(t *testing.T) {
	var out strings.Builder
	got, err := Select("", strings.NewReader("2\n"), &out)
	if err != nil {
		t.Fatalf("Select returned error: %v", err)
	}
	if got != "Claude" {
		t.Fatalf("Select = %q, want Claude", got)
	}
	if !strings.Contains(out.String(), "1) ChatGPT") || !strings.Contains(out.String(), "5) DeepSeek") {
		t.Fatalf("selection prompt = %q, want all supported sources", out.String())
	}
}

func TestSelectRejectsUnsupportedSource(t *testing.T) {
	if _, err := Select("cursor", strings.NewReader(""), &strings.Builder{}); err == nil {
		t.Fatal("expected unsupported source error")
	}
}

package config

import "testing"

func TestNormalizePublisher(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "default", input: "", want: defaultPublisher},
		{name: "plain", input: "replworks-bot", want: "replworks-bot"},
		{name: "with at", input: "@project-ai-bot", want: "project-ai-bot"},
		{name: "trim spaces", input: "  @engineering-bot  ", want: "engineering-bot"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizePublisher(tt.input); got != tt.want {
				t.Fatalf("normalizePublisher(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

package repository

import (
	"testing"
)

func TestExtractRepoName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "HTTPS URL",
			input:    "https://github.com/company/backend.git",
			expected: "company/backend",
		},
		{
			name:     "SSH URL",
			input:    "git@github.com:company/backend.git",
			expected: "company/backend",
		},
		{
			name:     "no .git suffix",
			input:    "https://github.com/user/repo",
			expected: "user/repo",
		},
		{
			name:     "empty",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractRepoName(tt.input)
			if got != tt.expected {
				t.Errorf("extractRepoName(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestValidateRepository(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "valid", input: "company/backend", wantErr: false},
		{name: "empty", input: "", wantErr: true},
		{name: "missing owner", input: "/backend", wantErr: true},
		{name: "missing repo", input: "company/", wantErr: true},
		{name: "invalid format", input: "company-backend", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRepository(tt.input)
			if tt.wantErr && err == nil {
				t.Fatalf("expected error for %q", tt.input)
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error for %q: %v", tt.input, err)
			}
		})
	}
}

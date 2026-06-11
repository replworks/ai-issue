package github

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateIssueSuccess(t *testing.T) {
	var gotMethod string
	var gotPath string
	var gotAuth string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"html_url":"https://github.com/company/backend/issues/1"}`))
	}))
	defer server.Close()

	client := &Client{
		Token:      "token-123",
		HTTPClient: server.Client(),
		BaseURL:    server.URL,
	}

	url, err := client.CreateIssue("company/backend", "Title", "Body")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if url != "https://github.com/company/backend/issues" {
		t.Fatalf("url = %q, want %q", url, "https://github.com/company/backend/issues")
	}
	if gotMethod != http.MethodPost {
		t.Fatalf("method = %q, want %q", gotMethod, http.MethodPost)
	}
	if gotPath != "/repos/company/backend/issues" {
		t.Fatalf("path = %q, want %q", gotPath, "/repos/company/backend/issues")
	}
	if gotAuth != "token token-123" {
		t.Fatalf("auth = %q, want %q", gotAuth, "token token-123")
	}
}

func TestCreateIssueFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	client := &Client{
		Token:      "token-123",
		HTTPClient: server.Client(),
		BaseURL:    server.URL,
	}

	_, err := client.CreateIssue("company/backend", "Title", "Body")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "GitHub issue publication failed") {
		t.Fatalf("error = %v, want publication failure message", err)
	}
}

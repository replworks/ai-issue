package github

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateIssueSuccess(t *testing.T) {
	var gotMethod string
	var gotPath string
	var gotAuth string
	var gotLabels []string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")
		var payload map[string]interface{}
		_ = json.NewDecoder(r.Body).Decode(&payload)
		if labels, ok := payload["labels"].([]interface{}); ok {
			gotLabels = make([]string, 0, len(labels))
			for _, label := range labels {
				gotLabels = append(gotLabels, label.(string))
			}
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"html_url":"https://github.com/company/backend/issues/1"}`))
	}))
	defer server.Close()

	client := &Client{
		Token:      "token-123",
		HTTPClient: server.Client(),
		BaseURL:    server.URL,
	}

	url, err := client.CreateIssue("company/backend", "Title", "Body", []string{"ai-generated"})
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
	if gotAuth != "Bearer token-123" {
		t.Fatalf("auth = %q, want %q", gotAuth, "Bearer token-123")
	}
	if len(gotLabels) != 1 || gotLabels[0] != "ai-generated" {
		t.Fatalf("labels = %v, want [ai-generated]", gotLabels)
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

	_, err := client.CreateIssue("company/backend", "Title", "Body", nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "GitHub issue publication failed") {
		t.Fatalf("error = %v, want publication failure message", err)
	}
}

func TestCheckRepositoryAccessSuccess(t *testing.T) {
	var gotMethod string
	var gotPath string
	var gotAuth string
	var gotAccept string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")
		gotAccept = r.Header.Get("Accept")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"full_name":"company/backend"}`))
	}))
	defer server.Close()

	client := &Client{
		Token:      "token-123",
		HTTPClient: server.Client(),
		BaseURL:    server.URL,
	}

	if err := client.CheckRepositoryAccess("company/backend"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Fatalf("method = %q, want %q", gotMethod, http.MethodGet)
	}
	if gotPath != "/repos/company/backend" {
		t.Fatalf("path = %q, want %q", gotPath, "/repos/company/backend")
	}
	if gotAuth != "Bearer token-123" {
		t.Fatalf("auth = %q, want %q", gotAuth, "Bearer token-123")
	}
	if gotAccept != "application/vnd.github+json" {
		t.Fatalf("accept = %q, want %q", gotAccept, "application/vnd.github+json")
	}
}

func TestCheckRepositoryAccessFailureParsesMessage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"message":"fine-grained tokens are not allowed for this org"}`))
	}))
	defer server.Close()

	client := &Client{
		Token:      "token-123",
		HTTPClient: server.Client(),
		BaseURL:    server.URL,
	}

	err := client.CheckRepositoryAccess("company/backend")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	accessErr, ok := err.(*RepositoryAccessError)
	if !ok {
		t.Fatalf("error type = %T, want *RepositoryAccessError", err)
	}
	if accessErr.Status != http.StatusText(http.StatusForbidden) {
		t.Fatalf("status = %q, want %q", accessErr.Status, http.StatusText(http.StatusForbidden))
	}
	if accessErr.Message != "fine-grained tokens are not allowed for this org" {
		t.Fatalf("message = %q, want %q", accessErr.Message, "fine-grained tokens are not allowed for this org")
	}
	if accessErr.RawBody == "" {
		t.Fatal("expected raw body to be preserved")
	}
	if accessErr.Headers == nil {
		t.Fatal("expected response headers to be preserved")
	}
}

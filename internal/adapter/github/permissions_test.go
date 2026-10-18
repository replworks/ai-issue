package github

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAuthenticatedLoginSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/user" {
			t.Fatalf("path = %q, want %q", r.URL.Path, "/user")
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"login":"replworks-bot"}`))
	}))
	defer server.Close()

	client := &Client{Token: "token-123", HTTPClient: server.Client(), BaseURL: server.URL}
	login, err := client.GetAuthenticatedLogin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if login != "replworks-bot" {
		t.Fatalf("login = %q, want %q", login, "replworks-bot")
	}
}

func TestCheckRepositoryAccessPullOnlyFails(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"full_name":"replworks/img-proxy","permissions":{"pull":true,"push":false,"admin":false}}`))
	}))
	defer server.Close()

	client := &Client{Token: "token-123", HTTPClient: server.Client(), BaseURL: server.URL}
	err := client.CheckRepositoryAccess("replworks/img-proxy")
	if err == nil {
		t.Fatal("expected error for pull-only access, got nil")
	}
	accessErr, ok := err.(*RepositoryAccessError)
	if !ok {
		t.Fatalf("error type = %T, want *RepositoryAccessError", err)
	}
	if !accessErr.HasPerms || accessErr.Push {
		t.Fatalf("expected HasPerms=true Push=false, got %+v", accessErr)
	}
	if !strings.Contains(accessErr.Message, "Write permission") {
		t.Fatalf("message = %q, want Write permission hint", accessErr.Message)
	}
	if accessErr.Repository != "replworks/img-proxy" {
		t.Fatalf("repository = %q, want replworks/img-proxy", accessErr.Repository)
	}
}

func TestCheckRepositoryAccessPushOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"full_name":"company/backend","permissions":{"pull":true,"push":true,"admin":false}}`))
	}))
	defer server.Close()

	client := &Client{Token: "token-123", HTTPClient: server.Client(), BaseURL: server.URL}
	if err := client.CheckRepositoryAccess("company/backend"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCreateIssueLabelPermissionErrorIncludesAccountAndRepo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/user":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"login":"replworks-bot"}`))
		case r.Method == http.MethodPost && r.URL.Path == "/repos/replworks/img-proxy/issues":
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(`{"message":"You do not have permission to create labels on this repository."}`))
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := &Client{Token: "token-123", HTTPClient: server.Client(), BaseURL: server.URL}
	_, err := client.CreateIssue("replworks/img-proxy", "Title", "Body", []string{"ai-generated"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	msg := err.Error()
	if !strings.Contains(msg, "replworks-bot") {
		t.Fatalf("error should include account login, got: %v", err)
	}
	if !strings.Contains(msg, "replworks/img-proxy") {
		t.Fatalf("error should include repository, got: %v", err)
	}
	if !strings.Contains(msg, "Write permission") {
		t.Fatalf("error should include Write permission hint, got: %v", err)
	}
}

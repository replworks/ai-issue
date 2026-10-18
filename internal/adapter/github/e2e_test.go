package github

import (
	"os"
	"testing"
)

// Live-boundary E2E check for issue #50. Skipped by default so
// `go test ./...` stays offline/deterministic (TECH_STACK Testing).
// Run with: AI_ISSUE_E2E=1 AI_ISSUE_E2E_REPO=replworks/ai-issue go test ./internal/adapter/github/ -run TestE2E -v
func TestE2ELiveRepositoryPermissions(t *testing.T) {
	if os.Getenv("AI_ISSUE_E2E") == "" {
		t.Skip("set AI_ISSUE_E2E=1 to run live GitHub API check")
	}
	client, err := NewClient()
	if err != nil {
		t.Skipf("no local token: %v", err)
	}
	repo := os.Getenv("AI_ISSUE_E2E_REPO")
	if repo == "" {
		repo = "replworks/ai-issue"
	}
	login, err := client.GetAuthenticatedLogin()
	if err != nil {
		t.Fatalf("live GET /user failed: %v", err)
	}
	t.Logf("authenticated as %s", login)
	if err := client.CheckRepositoryAccess(repo); err != nil {
		t.Logf("live access check for %s returned: %v", repo, err)
	} else {
		t.Logf("live access check for %s: Write OK", repo)
	}
}

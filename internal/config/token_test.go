package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveAndLoadToken(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", dir)

	if err := SaveToken("token-123"); err != nil {
		t.Fatalf("SaveToken error: %v", err)
	}

	got, err := LoadToken()
	if err != nil {
		t.Fatalf("LoadToken error: %v", err)
	}
	if got != "token-123" {
		t.Fatalf("LoadToken = %q, want %q", got, "token-123")
	}

	path, err := TokenPath()
	if err != nil {
		t.Fatalf("TokenPath error: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("token file missing: %v", err)
	}
	if filepath.Dir(path) == "" {
		t.Fatal("expected token path")
	}
}

package compliance

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestFrameworkCompliance(t *testing.T) {
	root := repoRoot(t)
	assertFileExists(t, root, "cmd/ai-issue/main.go")
	assertFileExists(t, root, "internal/cli/root.go")
	assertFileExists(t, root, "internal/cli/publish.go")
	assertFileExists(t, root, "internal/cli/diagnose.go")

	assertContains(t, root, "go.mod", "module ai-issue")
	assertContains(t, root, "go.mod", "go 1.24")
	assertContains(t, root, "internal/cli/root.go", "github.com/spf13/cobra")
	assertContains(t, root, "internal/extraction/extractor.go", "github.com/yuin/goldmark")
}

func TestArchitectureCompliance(t *testing.T) {
	root := repoRoot(t)
	assertContains(t, root, "PRODUCT_SPEC.md", "The first H1 heading becomes the issue title.")
	assertContains(t, root, "ARCHITECTURE.md", "Author != Publisher")
	assertContains(t, root, "FRAMEWORK.md", "AST-based parsing")
	assertContains(t, root, "FRAMEWORK.md", "Single Binary CLI")
}

func repoRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("unable to determine test file path")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
}

func assertFileExists(t *testing.T, root, rel string) {
	t.Helper()
	if _, err := os.Stat(filepath.Join(root, rel)); err != nil {
		t.Fatalf("expected file %s to exist: %v", rel, err)
	}
}

func assertContains(t *testing.T, root, rel, want string) {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, rel))
	if err != nil {
		t.Fatalf("failed reading %s: %v", rel, err)
	}
	if !strings.Contains(string(data), want) {
		t.Fatalf("%s does not contain %q", rel, want)
	}
}

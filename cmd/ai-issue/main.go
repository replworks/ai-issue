package main

import (
	"os"

	"github.com/replworks/ai-issue/internal/cli"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	_, _ = commit, date
	cli.SetVersion(version)
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}

package cli

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/replworks/ai-issue/internal/adapter/github"
	"github.com/replworks/ai-issue/internal/config"
)

var loginCmd = &cobra.Command{
	Use:          "login",
	Short:        "Authenticate with GitHub App device flow",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runLogin(cmd.Context())
	},
}

func runLogin(ctx context.Context) error {
	clientID := config.GitHubAppClientID()
	device, err := github.RequestDeviceCode(clientID)
	if err != nil {
		return err
	}

	fmt.Println("Open this URL in your browser:")
	fmt.Println(device.VerificationURI)
	fmt.Printf("Enter this code: %s\n", device.UserCode)
	if err := openBrowser(device.VerificationURI); err == nil {
		fmt.Println("Browser opened automatically.")
	} else {
		fmt.Printf("Could not open browser automatically: %v\n", err)
	}

	interval := time.Duration(device.Interval) * time.Second
	if interval <= 0 {
		interval = 5 * time.Second
	}
	deadline := time.Now().Add(time.Duration(device.ExpiresIn) * time.Second)
	if device.ExpiresIn <= 0 {
		deadline = time.Now().Add(15 * time.Minute)
	}

	for time.Now().Before(deadline) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(interval):
		}

		token, err := github.ExchangeDeviceCode(clientID, device.DeviceCode)
		if err != nil {
			if strings.Contains(err.Error(), "authorization_pending") || strings.Contains(err.Error(), "slow_down") {
				continue
			}
			if strings.Contains(err.Error(), "access_denied") {
				return fmt.Errorf("login was denied in GitHub")
			}
			if strings.Contains(err.Error(), "expired_token") {
				return fmt.Errorf("device code expired. Run `ai-issue login` again")
			}
			return err
		}
		if strings.TrimSpace(token.AccessToken) == "" {
			continue
		}
		if err := config.SaveToken(token.AccessToken); err != nil {
			return err
		}
		fmt.Println("Login successful. Token saved locally.")
		return nil
	}

	return fmt.Errorf("device code expired. Run `ai-issue login` again")
}

func openBrowser(target string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", target).Run()
	case "windows":
		return exec.Command("cmd", "/c", "start", "", target).Run()
	default:
		return exec.Command("xdg-open", target).Run()
	}
}

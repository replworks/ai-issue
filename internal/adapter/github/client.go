package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/replworks/ai-issue/internal/config"
	"github.com/replworks/ai-issue/internal/extraction"
)

type Client struct {
	Token      string
	HTTPClient *http.Client
	BaseURL    string
}

var (
	deviceCodeEndpoint  = "https://github.com/login/device/code"
	accessTokenEndpoint = "https://github.com/login/oauth/access_token"
)

func SetDeviceFlowEndpointsForTest(device, access string) {
	deviceCodeEndpoint = device
	accessTokenEndpoint = access
}

func RestoreDeviceFlowEndpoints(device, access string) {
	deviceCodeEndpoint = device
	accessTokenEndpoint = access
}

func DeviceFlowEndpointsForTest() (string, string) {
	return deviceCodeEndpoint, accessTokenEndpoint
}

type APIError struct {
	Message string `json:"message"`
}

func NewClient() (*Client, error) {
	token, err := config.LoadToken()
	if err != nil {
		return nil, extraction.NewError("auth", err.Error())
	}
	if token == "" {
		return nil, extraction.NewError("auth", "GitHub App token is required. Run `ai-issue login` first.")
	}
	return &Client{
		Token:      token,
		HTTPClient: &http.Client{},
		BaseURL:    "https://api.github.com",
	}, nil
}

type DeviceCodeResponse struct {
	DeviceCode      string
	UserCode        string
	VerificationURI string
	ExpiresIn       int
	Interval        int
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	Error       string `json:"error"`
}

func RequestDeviceCode(clientID string) (*DeviceCodeResponse, error) {
	form := url.Values{}
	form.Set("client_id", clientID)

	req, err := http.NewRequest(http.MethodPost, deviceCodeEndpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create device code request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request device code: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("device code request failed: %s", githubErrorMessage(body))
	}

	var decoded struct {
		DeviceCode      string `json:"device_code"`
		UserCode        string `json:"user_code"`
		VerificationURI string `json:"verification_uri"`
		ExpiresIn       int    `json:"expires_in"`
		Interval        int    `json:"interval"`
	}
	if err := json.Unmarshal(body, &decoded); err != nil {
		return nil, fmt.Errorf("failed to decode device code response: %w", err)
	}
	return &DeviceCodeResponse{
		DeviceCode:      decoded.DeviceCode,
		UserCode:        decoded.UserCode,
		VerificationURI: decoded.VerificationURI,
		ExpiresIn:       decoded.ExpiresIn,
		Interval:        decoded.Interval,
	}, nil
}

func ExchangeDeviceCode(clientID, deviceCode string) (*AccessTokenResponse, error) {
	form := url.Values{}
	form.Set("client_id", clientID)
	form.Set("device_code", deviceCode)
	form.Set("grant_type", "urn:ietf:params:oauth:grant-type:device_code")

	req, err := http.NewRequest(http.MethodPost, accessTokenEndpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create access token request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange device code: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("access token request failed: %s", githubErrorMessage(body))
	}

	var decoded AccessTokenResponse
	if err := json.Unmarshal(body, &decoded); err != nil {
		return nil, fmt.Errorf("failed to decode access token response: %w", err)
	}
	return &decoded, nil
}

func (c *Client) CreateIssue(repo, title, body string, labels []string) (string, error) {
	if c == nil {
		return "", fmt.Errorf("GitHub client is not initialized")
	}
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{}
	}
	if c.BaseURL == "" {
		c.BaseURL = "https://api.github.com"
	}

	url := fmt.Sprintf("%s/repos/%s/issues", c.BaseURL, repo)

	payload := map[string]interface{}{
		"title": title,
		"body":  body,
	}
	if len(labels) > 0 {
		payload["labels"] = labels
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to encode GitHub issue payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create GitHub issue request: %w", err)
	}
	req.Header.Set("Authorization", "token "+c.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to publish GitHub issue: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	responseBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("GitHub issue publication failed\nRepository: %s\nStatus: %s\nReason:\n%s", repo, resp.Status, githubErrorMessage(responseBody))
	}

	// In real implementation, parse response for HTML URL
	return fmt.Sprintf("https://github.com/%s/issues", repo), nil
}

func (c *Client) CheckRepositoryAccess(repo string) error {
	if c == nil {
		return fmt.Errorf("GitHub client is not initialized")
	}
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{}
	}
	if c.BaseURL == "" {
		c.BaseURL = "https://api.github.com"
	}

	url := fmt.Sprintf("%s/repos/%s", c.BaseURL, repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create GitHub repository access request: %w", err)
	}
	req.Header.Set("Authorization", "token "+c.Token)
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to verify repository access: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return &RepositoryAccessError{
			Status:  http.StatusText(resp.StatusCode),
			Message: githubErrorMessage(body),
		}
	}

	return nil
}

type RepositoryAccessError struct {
	Status  string
	Message string
}

func (e *RepositoryAccessError) Error() string {
	if e == nil {
		return "repository access validation failed"
	}
	return fmt.Sprintf("repository access validation failed: %s", e.Status)
}

func githubErrorMessage(body []byte) string {
	var payload APIError
	if err := json.Unmarshal(body, &payload); err == nil && strings.TrimSpace(payload.Message) != "" {
		return payload.Message
	}

	trimmed := strings.TrimSpace(string(body))
	if trimmed == "" {
		return "Unknown GitHub API error."
	}
	return trimmed
}

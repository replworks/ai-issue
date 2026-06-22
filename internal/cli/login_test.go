package cli

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/replworks/ai-issue/internal/adapter/github"
	"github.com/replworks/ai-issue/internal/config"
)

func TestRunLoginSavesToken(t *testing.T) {
	oldDevice, oldAccess := github.DeviceFlowEndpointsForTest()
	defer github.RestoreDeviceFlowEndpoints(oldDevice, oldAccess)

	deviceCalls := 0
	accessCalls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/login/device/code":
			deviceCalls++
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"device_code":"device-123","user_code":"WDJB-MJHT","verification_uri":"https://github.com/login/device","expires_in":900,"interval":1}`))
		case "/login/oauth/access_token":
			accessCalls++
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"token-123","token_type":"bearer","scope":""}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer server.Close()

	github.SetDeviceFlowEndpointsForTest(server.URL+"/login/device/code", server.URL+"/login/oauth/access_token")

	t.Setenv("XDG_CONFIG_HOME", t.TempDir())

	if err := runLogin(t.Context()); err != nil {
		t.Fatalf("runLogin error: %v", err)
	}

	if deviceCalls != 1 {
		t.Fatalf("deviceCalls = %d, want 1", deviceCalls)
	}
	if accessCalls == 0 {
		t.Fatal("expected access token polling")
	}

	token, err := config.LoadToken()
	if err != nil {
		t.Fatalf("LoadToken error: %v", err)
	}
	if token != "token-123" {
		t.Fatalf("token = %q, want %q", token, "token-123")
	}
}

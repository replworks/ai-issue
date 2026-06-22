package github

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestAndExchangeDeviceCode(t *testing.T) {
	oldDevice := deviceCodeEndpoint
	oldAccess := accessTokenEndpoint
	defer func() {
		deviceCodeEndpoint = oldDevice
		accessTokenEndpoint = oldAccess
	}()

	var gotDevicePath string
	var gotAccessPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/login/device/code":
			gotDevicePath = r.URL.Path
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"device_code":"device-123","user_code":"WDJB-MJHT","verification_uri":"https://github.com/login/device","expires_in":900,"interval":5}`))
		case "/login/oauth/access_token":
			gotAccessPath = r.URL.Path
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"token-123","token_type":"bearer","scope":""}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer server.Close()

	deviceCodeEndpoint = server.URL + "/login/device/code"
	accessTokenEndpoint = server.URL + "/login/oauth/access_token"

	device, err := RequestDeviceCode("client-123")
	if err != nil {
		t.Fatalf("RequestDeviceCode error: %v", err)
	}
	if device.DeviceCode != "device-123" || device.UserCode != "WDJB-MJHT" {
		t.Fatalf("unexpected device response: %#v", device)
	}
	if gotDevicePath != "/login/device/code" {
		t.Fatalf("gotDevicePath = %q", gotDevicePath)
	}

	token, err := ExchangeDeviceCode("client-123", device.DeviceCode)
	if err != nil {
		t.Fatalf("ExchangeDeviceCode error: %v", err)
	}
	if token.AccessToken != "token-123" {
		t.Fatalf("token = %#v", token)
	}
	if gotAccessPath != "/login/oauth/access_token" {
		t.Fatalf("gotAccessPath = %q", gotAccessPath)
	}
}

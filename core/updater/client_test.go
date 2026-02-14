package updater

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPProviderCheck(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		var req UpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if req.AppID != "daf-wails-template" {
			t.Fatalf("unexpected appId: %s", req.AppID)
		}
		if req.Channel != "stable" {
			t.Fatalf("unexpected channel: %s", req.Channel)
		}

		info := UpdateInfo{
			ProtocolVersion: "1",
			LatestVersion:   "1.1.0",
			DownloadURL:     "https://cdn.example.com/app.zip",
			Checksum:        "abc",
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(info); err != nil {
			t.Fatalf("encode response: %v", err)
		}
	}))
	defer server.Close()

	provider := &HTTPProvider{Endpoint: server.URL}
	info, err := provider.Check(routineContext(), UpdateRequest{
		AppID:    "daf-wails-template",
		Version:  "1.0.0",
		Platform: "windows",
		Arch:     "amd64",
		Channel:  "stable",
	})
	if err != nil {
		t.Fatalf("check update: %v", err)
	}
	if info.LatestVersion != "1.1.0" {
		t.Fatalf("unexpected latest version: %s", info.LatestVersion)
	}
}

package updater

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type UpdateProvider interface {
	Check(ctx context.Context, req UpdateRequest) (UpdateInfo, error)
}

type HTTPProvider struct {
	Endpoint string
	Client   *http.Client
}

func (p *HTTPProvider) Check(ctx context.Context, req UpdateRequest) (UpdateInfo, error) {
	if p == nil {
		return UpdateInfo{}, fmt.Errorf("provider is nil")
	}
	if strings.TrimSpace(p.Endpoint) == "" {
		return UpdateInfo{}, fmt.Errorf("provider endpoint is empty")
	}

	logger := slog.Default()
	start := time.Now()
	logger.Info("updater.check.start",
		"endpoint", p.Endpoint,
		"appId", req.AppID,
		"version", req.Version,
		"channel", req.Channel,
		"platform", req.Platform,
		"arch", req.Arch,
	)

	payload, err := json.Marshal(req)
	if err != nil {
		return UpdateInfo{}, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, p.Endpoint, bytes.NewReader(payload))
	if err != nil {
		logger.Error("updater.check.request_error", "endpoint", p.Endpoint, "error", err)
		return UpdateInfo{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	client := p.Client
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(request)
	if err != nil {
		logger.Error("updater.check.request_failed", "endpoint", p.Endpoint, "error", err, "duration", time.Since(start))
		return UpdateInfo{}, err
	}
	defer resp.Body.Close()

	logger.Info("updater.check.response",
		"endpoint", p.Endpoint,
		"status", resp.StatusCode,
		"duration", time.Since(start),
	)

	if resp.StatusCode != http.StatusOK {
		message, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		text := strings.TrimSpace(string(message))
		if text == "" {
			text = resp.Status
		}
		logger.Error("updater.check.failed", "endpoint", p.Endpoint, "status", resp.StatusCode, "error", text)
		return UpdateInfo{}, fmt.Errorf("update check failed: %s", text)
	}

	var info UpdateInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		logger.Error("updater.check.decode_failed", "endpoint", p.Endpoint, "error", err)
		return UpdateInfo{}, err
	}

	logger.Info("updater.check.success",
		"endpoint", p.Endpoint,
		"latestVersion", info.LatestVersion,
		"force", info.Force,
		"duration", time.Since(start),
	)

	return info, nil
}

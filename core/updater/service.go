package updater

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

var ErrPolicyDenied = errors.New("update denied by policy")

type Updater struct {
	Provider   UpdateProvider
	Policy     UpdatePolicy
	Downloader Downloader
	Verifier   Verifier
	Installer  Installer
}

func (u *Updater) Run(ctx context.Context, req UpdateRequest) (InstallResult, error) {
	logger := slog.Default()
	start := time.Now()
	logger.Info("updater.run.start",
		"appId", req.AppID,
		"version", req.Version,
		"channel", req.Channel,
		"platform", req.Platform,
		"arch", req.Arch,
	)

	if u == nil {
		logger.Error("updater.run.failed", "step", "init", "error", "updater is nil")
		return InstallResult{}, fmt.Errorf("updater is nil")
	}
	if u.Provider == nil {
		logger.Error("updater.run.failed", "step", "provider", "error", "provider is nil")
		return InstallResult{}, fmt.Errorf("provider is nil")
	}
	if u.Policy == nil {
		logger.Error("updater.run.failed", "step", "policy", "error", "policy is nil")
		return InstallResult{}, fmt.Errorf("policy is nil")
	}
	if u.Downloader == nil {
		logger.Error("updater.run.failed", "step", "downloader", "error", "downloader is nil")
		return InstallResult{}, fmt.Errorf("downloader is nil")
	}
	if u.Verifier == nil {
		logger.Error("updater.run.failed", "step", "verifier", "error", "verifier is nil")
		return InstallResult{}, fmt.Errorf("verifier is nil")
	}
	if u.Installer == nil {
		logger.Error("updater.run.failed", "step", "installer", "error", "installer is nil")
		return InstallResult{}, fmt.Errorf("installer is nil")
	}

	info, err := u.Provider.Check(ctx, req)
	if err != nil {
		logger.Error("updater.run.failed", "step", "check", "error", err, "duration", time.Since(start))
		return InstallResult{}, err
	}
	logger.Info("updater.run.check_ok", "latestVersion", info.LatestVersion, "force", info.Force)

	decision, err := u.Policy.Decide(ctx, info)
	if err != nil {
		logger.Error("updater.run.failed", "step", "policy", "error", err, "duration", time.Since(start))
		return InstallResult{}, err
	}
	if !decision.Allowed {
		logger.Warn("updater.run.denied", "reason", decision.Reason)
		if decision.Reason == "" {
			return InstallResult{}, ErrPolicyDenied
		}
		return InstallResult{}, fmt.Errorf("%w: %s", ErrPolicyDenied, decision.Reason)
	}
	logger.Info("updater.run.policy_ok")

	download, err := u.Downloader.Download(ctx, info)
	if err != nil {
		logger.Error("updater.run.failed", "step", "download", "error", err, "duration", time.Since(start))
		return InstallResult{}, err
	}
	logger.Info("updater.run.download_ok", "path", download.Path, "size", download.Size)

	if err := u.Verifier.Verify(ctx, info, download); err != nil {
		logger.Error("updater.run.failed", "step", "verify", "error", err, "duration", time.Since(start))
		return InstallResult{}, err
	}
	logger.Info("updater.run.verify_ok")

	result, err := u.Installer.Install(ctx, info, download)
	if err != nil {
		logger.Error("updater.run.failed", "step", "install", "error", err, "duration", time.Since(start))
		return InstallResult{}, err
	}
	logger.Info("updater.run.success", "path", result.Path, "duration", time.Since(start))

	return result, nil
}

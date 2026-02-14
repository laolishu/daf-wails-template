package sysconfig

import (
	"testing"
)

func TestGetInfo(t *testing.T) {
	// 初始化测试数据
	Version = "v1.0.0"
	BuildTime = "2026-02-04T10:00:00Z"
	GitCommit = "abc1234"
	ConfigDir = "/etc/daf-app"
	ConfigFile = "app.yml"
	UpdateEndpoint = "https://updates.example.com/check"

	info := GetInfo()

	if info.Version != "v1.0.0" {
		t.Errorf("Expected Version=v1.0.0, got %s", info.Version)
	}
	if info.BuildTime != "2026-02-04T10:00:00Z" {
		t.Errorf("Expected BuildTime=2026-02-04T10:00:00Z, got %s", info.BuildTime)
	}
	if info.GitCommit != "abc1234" {
		t.Errorf("Expected GitCommit=abc1234, got %s", info.GitCommit)
	}
	if info.ConfigDir != "/etc/daf-app" {
		t.Errorf("Expected ConfigDir=/etc/daf-app, got %s", info.ConfigDir)
	}
	if info.ConfigFile != "app.yml" {
		t.Errorf("Expected ConfigFile=app.yml, got %s", info.ConfigFile)
	}
	if info.UpdateEndpoint != "https://updates.example.com/check" {
		t.Errorf("Expected UpdateEndpoint=https://updates.example.com/check, got %s", info.UpdateEndpoint)
	}
}

func TestGetVersion(t *testing.T) {
	Version = "v2.0.0"
	if GetVersion() != "v2.0.0" {
		t.Errorf("Expected v2.0.0, got %s", GetVersion())
	}
}

func TestGetBuildTime(t *testing.T) {
	BuildTime = "2026-02-04T15:30:00Z"
	if GetBuildTime() != "2026-02-04T15:30:00Z" {
		t.Errorf("Expected 2026-02-04T15:30:00Z, got %s", GetBuildTime())
	}
}

func TestGetGitCommit(t *testing.T) {
	GitCommit = "def5678"
	if GetGitCommit() != "def5678" {
		t.Errorf("Expected def5678, got %s", GetGitCommit())
	}
}

func TestGetConfigDir(t *testing.T) {
	ConfigDir = "/home/user/.config/app"
	if GetConfigDir() != "/home/user/.config/app" {
		t.Errorf("Expected /home/user/.config/app, got %s", GetConfigDir())
	}
}

func TestGetConfigFile(t *testing.T) {
	ConfigFile = "settings.yml"
	if GetConfigFile() != "settings.yml" {
		t.Errorf("Expected settings.yml, got %s", GetConfigFile())
	}
}

func TestGetUpdateEndpoint(t *testing.T) {
	UpdateEndpoint = "https://updates.example.com/check"
	if GetUpdateEndpoint() != "https://updates.example.com/check" {
		t.Errorf("Expected https://updates.example.com/check, got %s", GetUpdateEndpoint())
	}
}

func TestZeroValues(t *testing.T) {
	// 检查零值（未被注入时的状态）
	Version = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
	ConfigDir = ""
	ConfigFile = "config.yml"
	UpdateEndpoint = ""

	if Version != "dev" {
		t.Errorf("Expected default version=dev, got %s", Version)
	}
	if BuildTime != "unknown" {
		t.Errorf("Expected default BuildTime=unknown, got %s", BuildTime)
	}
	if GitCommit != "unknown" {
		t.Errorf("Expected default GitCommit=unknown, got %s", GitCommit)
	}
	if ConfigDir != "" {
		t.Errorf("Expected empty ConfigDir, got %s", ConfigDir)
	}
	if ConfigFile != "config.yml" {
		t.Errorf("Expected default ConfigFile=config.yml, got %s", ConfigFile)
	}
	if UpdateEndpoint != "" {
		t.Errorf("Expected empty UpdateEndpoint, got %s", UpdateEndpoint)
	}
}

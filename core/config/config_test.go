package config

import (
	"os"
	"path/filepath"
	"testing"

	"daf-wails-template/core/sysconfig"
)

func TestLoadCreatesDefaultConfig(t *testing.T) {
	tempDir := t.TempDir()
	restoreSysconfig(t, tempDir, "config.yml")

	manager, err := NewManager()
	if err != nil {
		t.Fatalf("NewManager() error = %v", err)
	}

	if err := manager.Load(); err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if manager.GetLogDir() != DefaultLogDir {
		t.Fatalf("GetLogDir() = %q, want %q", manager.GetLogDir(), DefaultLogDir)
	}
	if manager.GetLogLevel() != DefaultLogLevel {
		t.Fatalf("GetLogLevel() = %q, want %q", manager.GetLogLevel(), DefaultLogLevel)
	}
	if manager.GetLogRetentionDays() != DefaultLogRetentionDays {
		t.Fatalf("GetLogRetentionDays() = %d, want %d", manager.GetLogRetentionDays(), DefaultLogRetentionDays)
	}
	if manager.GetLanguage() != DefaultLanguage {
		t.Fatalf("GetLanguage() = %q, want %q", manager.GetLanguage(), DefaultLanguage)
	}
	if manager.GetWindowWidth() != DefaultWindowWidth {
		t.Fatalf("GetWindowWidth() = %d, want %d", manager.GetWindowWidth(), DefaultWindowWidth)
	}
	if manager.GetWindowHeight() != DefaultWindowHeight {
		t.Fatalf("GetWindowHeight() = %d, want %d", manager.GetWindowHeight(), DefaultWindowHeight)
	}
	if manager.GetWindowTitle() != DefaultWindowTitle {
		t.Fatalf("GetWindowTitle() = %q, want %q", manager.GetWindowTitle(), DefaultWindowTitle)
	}

	configPath := filepath.Join(tempDir, "config.yml")
	if _, err := os.Stat(configPath); err != nil {
		t.Fatalf("config file not created: %v", err)
	}
}

func TestLoadExistingConfig(t *testing.T) {
	tempDir := t.TempDir()
	restoreSysconfig(t, tempDir, "config.yml")

	configPath := filepath.Join(tempDir, "config.yml")
	content := "log:\n  dir: \"custom-logs\"\n  level: \"warn\"\n  retention_days: 14\ni18n:\n  language: \"en-US\"\nwindow:\n  width: 1280\n  height: 720\n  title: \"Custom App\"\n"
	if err := os.WriteFile(configPath, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	manager, err := NewManager()
	if err != nil {
		t.Fatalf("NewManager() error = %v", err)
	}

	if err := manager.Load(); err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if manager.GetLogDir() != "custom-logs" {
		t.Fatalf("GetLogDir() = %q, want %q", manager.GetLogDir(), "custom-logs")
	}
	if manager.GetLogLevel() != "warn" {
		t.Fatalf("GetLogLevel() = %q, want %q", manager.GetLogLevel(), "warn")
	}
	if manager.GetLogRetentionDays() != 14 {
		t.Fatalf("GetLogRetentionDays() = %d, want %d", manager.GetLogRetentionDays(), 14)
	}
	if manager.GetLanguage() != "en-US" {
		t.Fatalf("GetLanguage() = %q, want %q", manager.GetLanguage(), "en-US")
	}
	if manager.GetWindowWidth() != 1280 {
		t.Fatalf("GetWindowWidth() = %d, want %d", manager.GetWindowWidth(), 1280)
	}
	if manager.GetWindowHeight() != 720 {
		t.Fatalf("GetWindowHeight() = %d, want %d", manager.GetWindowHeight(), 720)
	}
	if manager.GetWindowTitle() != "Custom App" {
		t.Fatalf("GetWindowTitle() = %q, want %q", manager.GetWindowTitle(), "Custom App")
	}
}

func TestSaveWritesConfig(t *testing.T) {
	tempDir := t.TempDir()
	restoreSysconfig(t, tempDir, "config.yml")

	manager, err := NewManager()
	if err != nil {
		t.Fatalf("NewManager() error = %v", err)
	}

	manager.SetLogDir("runtime-logs")
	manager.SetLogLevel("debug")
	manager.SetLogRetentionDays(3)
	manager.SetLanguage("ja-JP")
	manager.SetWindowWidth(800)
	manager.SetWindowHeight(600)
	manager.SetWindowTitle("Test App")

	if err := manager.Save(); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	reloaded, err := NewManager()
	if err != nil {
		t.Fatalf("NewManager() error = %v", err)
	}

	if err := reloaded.Load(); err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if reloaded.GetLogDir() != "runtime-logs" {
		t.Fatalf("GetLogDir() = %q, want %q", reloaded.GetLogDir(), "runtime-logs")
	}
	if reloaded.GetLogLevel() != "debug" {
		t.Fatalf("GetLogLevel() = %q, want %q", reloaded.GetLogLevel(), "debug")
	}
	if reloaded.GetLogRetentionDays() != 3 {
		t.Fatalf("GetLogRetentionDays() = %d, want %d", reloaded.GetLogRetentionDays(), 3)
	}
	if reloaded.GetLanguage() != "ja-JP" {
		t.Fatalf("GetLanguage() = %q, want %q", reloaded.GetLanguage(), "ja-JP")
	}
	if reloaded.GetWindowWidth() != 800 {
		t.Fatalf("GetWindowWidth() = %d, want %d", reloaded.GetWindowWidth(), 800)
	}
	if reloaded.GetWindowHeight() != 600 {
		t.Fatalf("GetWindowHeight() = %d, want %d", reloaded.GetWindowHeight(), 600)
	}
	if reloaded.GetWindowTitle() != "Test App" {
		t.Fatalf("GetWindowTitle() = %q, want %q", reloaded.GetWindowTitle(), "Test App")
	}
}

func restoreSysconfig(t *testing.T, configDir, configFile string) {
	origDir := sysconfig.ConfigDir
	origFile := sysconfig.ConfigFile

	sysconfig.ConfigDir = configDir
	sysconfig.ConfigFile = configFile

	t.Cleanup(func() {
		sysconfig.ConfigDir = origDir
		sysconfig.ConfigFile = origFile
	})
}

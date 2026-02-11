package logger

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCleanupOldFiles(t *testing.T) {
	dir := t.TempDir()
	now := time.Now().In(time.Local)

	createLogFile(t, dir, now)
	createLogFile(t, dir, now.AddDate(0, 0, -1))
	createLogFile(t, dir, now.AddDate(0, 0, -6))
	createLogFile(t, dir, now.AddDate(0, 0, -7))
	createLogFile(t, dir, now.AddDate(0, 0, -10))

	if err := cleanupOldFiles(dir, 7); err != nil {
		t.Fatalf("cleanupOldFiles() error = %v", err)
	}

	assertExists(t, filepath.Join(dir, now.Format("2006-01-02")+".log"))
	assertExists(t, filepath.Join(dir, now.AddDate(0, 0, -1).Format("2006-01-02")+".log"))
	assertExists(t, filepath.Join(dir, now.AddDate(0, 0, -6).Format("2006-01-02")+".log"))
	assertNotExists(t, filepath.Join(dir, now.AddDate(0, 0, -7).Format("2006-01-02")+".log"))
	assertNotExists(t, filepath.Join(dir, now.AddDate(0, 0, -10).Format("2006-01-02")+".log"))
}

func TestParseLevel(t *testing.T) {
	cases := map[string]string{
		"debug": "DEBUG",
		"info":  "INFO",
		"warn":  "WARN",
		"error": "ERROR",
		"":      "INFO",
	}

	for input, expected := range cases {
		level := parseLevel(input)
		if level.String() != expected {
			t.Fatalf("parseLevel(%q) = %s, want %s", input, level.String(), expected)
		}
	}
}

func createLogFile(t *testing.T, dir string, date time.Time) {
	t.Helper()
	name := filepath.Join(dir, date.Format("2006-01-02")+".log")
	if err := os.WriteFile(name, []byte("test"), 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
}

func assertExists(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected %s to exist: %v", path, err)
	}
}

func assertNotExists(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); err == nil {
		t.Fatalf("expected %s to be removed", path)
	}
}

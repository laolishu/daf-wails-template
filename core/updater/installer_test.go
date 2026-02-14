package updater

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestBasicInstallerDownload(t *testing.T) {
	payload := []byte("download-content")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(payload)
	}))
	defer server.Close()

	installer := &BasicInstaller{TempDir: t.TempDir()}
	result, err := installer.Download(routineContext(), UpdateInfo{DownloadURL: server.URL})
	if err != nil {
		t.Fatalf("download: %v", err)
	}

	data, err := os.ReadFile(result.Path)
	if err != nil {
		t.Fatalf("read file: %v", err)
	}
	if !bytes.Equal(data, payload) {
		t.Fatalf("downloaded content mismatch")
	}
	if result.Checksum == "" {
		t.Fatalf("checksum empty")
	}
}

func TestBasicInstallerInstall(t *testing.T) {
	zipPath := filepath.Join(t.TempDir(), "payload.zip")
	zipFile, err := os.Create(zipPath)
	if err != nil {
		t.Fatalf("create zip: %v", err)
	}

	writer := zip.NewWriter(zipFile)
	entry, err := writer.Create("app/readme.txt")
	if err != nil {
		t.Fatalf("create entry: %v", err)
	}
	if _, err := entry.Write([]byte("ok")); err != nil {
		t.Fatalf("write entry: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("close zip: %v", err)
	}
	if err := zipFile.Close(); err != nil {
		t.Fatalf("close file: %v", err)
	}

	content, err := os.ReadFile(zipPath)
	if err != nil {
		t.Fatalf("read zip: %v", err)
	}
	hash := sha256.Sum256(content)
	checksum := hex.EncodeToString(hash[:])

	installer := &BasicInstaller{TempDir: t.TempDir()}
	result, err := installer.Install(routineContext(), UpdateInfo{Checksum: checksum}, DownloadResult{Path: zipPath, Checksum: checksum})
	if err != nil {
		t.Fatalf("install: %v", err)
	}

	extracted := filepath.Join(result.Path, "app", "readme.txt")
	data, err := os.ReadFile(extracted)
	if err != nil {
		t.Fatalf("read extracted: %v", err)
	}
	if string(data) != "ok" {
		t.Fatalf("unexpected content: %s", string(data))
	}
}

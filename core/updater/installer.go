package updater

import (
	"archive/zip"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Installer interface {
	Install(ctx context.Context, info UpdateInfo, file DownloadResult) (InstallResult, error)
}

type BasicInstaller struct {
	TempDir string
	Client  *http.Client
}

func (b *BasicInstaller) Download(ctx context.Context, info UpdateInfo) (DownloadResult, error) {
	if b == nil {
		return DownloadResult{}, fmt.Errorf("installer is nil")
	}
	if strings.TrimSpace(info.DownloadURL) == "" {
		return DownloadResult{}, fmt.Errorf("download url is empty")
	}

	logger := slog.Default()
	start := time.Now()
	logger.Info("updater.download.start", "url", info.DownloadURL)

	tempRoot := b.tempRoot()
	if err := os.MkdirAll(tempRoot, 0o755); err != nil {
		return DownloadResult{}, err
	}

	client := b.Client
	if client == nil {
		client = http.DefaultClient
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, info.DownloadURL, nil)
	if err != nil {
		return DownloadResult{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("updater.download.request_failed", "url", info.DownloadURL, "error", err, "duration", time.Since(start))
		return DownloadResult{}, err
	}
	defer resp.Body.Close()

	logger.Info("updater.download.response", "url", info.DownloadURL, "status", resp.StatusCode, "duration", time.Since(start))

	if resp.StatusCode != http.StatusOK {
		logger.Error("updater.download.failed", "url", info.DownloadURL, "status", resp.StatusCode)
		return DownloadResult{}, fmt.Errorf("download failed: %s", resp.Status)
	}

	file, err := os.CreateTemp(tempRoot, "update-*.zip")
	if err != nil {
		return DownloadResult{}, err
	}
	defer file.Close()

	hasher := sha256.New()
	writer := io.MultiWriter(file, hasher)
	size, err := io.Copy(writer, resp.Body)
	if err != nil {
		logger.Error("updater.download.write_failed", "url", info.DownloadURL, "error", err, "duration", time.Since(start))
		return DownloadResult{}, err
	}
	result := DownloadResult{
		Path:     file.Name(),
		Size:     size,
		Checksum: hex.EncodeToString(hasher.Sum(nil)),
	}

	logger.Info("updater.download.success", "url", info.DownloadURL, "path", result.Path, "size", result.Size, "duration", time.Since(start))

	return result, nil
}

func (b *BasicInstaller) Install(ctx context.Context, info UpdateInfo, file DownloadResult) (InstallResult, error) {
	if b == nil {
		return InstallResult{}, fmt.Errorf("installer is nil")
	}
	if strings.TrimSpace(file.Path) == "" {
		return InstallResult{}, fmt.Errorf("download path is empty")
	}

	logger := slog.Default()
	start := time.Now()
	logger.Info("updater.install.start", "path", file.Path)

	if err := verifyChecksum(info, file); err != nil {
		logger.Error("updater.install.verify_failed", "path", file.Path, "error", err)
		return InstallResult{}, err
	}
	logger.Info("updater.install.verify_ok", "path", file.Path)

	dest, err := os.MkdirTemp(b.tempRoot(), "update-install-")
	if err != nil {
		return InstallResult{}, err
	}

	if err := unzip(file.Path, dest); err != nil {
		logger.Error("updater.install.unzip_failed", "path", file.Path, "dest", dest, "error", err, "duration", time.Since(start))
		return InstallResult{}, err
	}

	result := InstallResult{Path: dest}
	logger.Info("updater.install.success", "path", result.Path, "duration", time.Since(start))

	return result, nil
}

func (b *BasicInstaller) tempRoot() string {
	if strings.TrimSpace(b.TempDir) == "" {
		return os.TempDir()
	}
	return b.TempDir
}

func verifyChecksum(info UpdateInfo, file DownloadResult) error {
	expected := strings.TrimSpace(info.Checksum)
	if expected == "" {
		return fmt.Errorf("checksum is empty")
	}

	actual := strings.TrimSpace(file.Checksum)
	if actual == "" {
		var err error
		actual, err = computeFileSHA256(file.Path)
		if err != nil {
			return err
		}
	}

	if strings.ToLower(expected) != strings.ToLower(actual) {
		return fmt.Errorf("checksum mismatch")
	}
	return nil
}

func unzip(source, destination string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if err := extractZipEntry(file, destination); err != nil {
			return err
		}
	}
	return nil
}

func extractZipEntry(file *zip.File, destination string) error {
	cleanName := filepath.Clean(file.Name)
	if filepath.IsAbs(cleanName) || cleanName == "." || cleanName == ".." || strings.HasPrefix(cleanName, ".."+string(os.PathSeparator)) {
		return fmt.Errorf("invalid zip entry: %s", file.Name)
	}

	path := filepath.Join(destination, cleanName)
	if file.FileInfo().IsDir() {
		return os.MkdirAll(path, file.Mode())
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	input, err := file.Open()
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer output.Close()

	if _, err := io.Copy(output, input); err != nil {
		return err
	}
	return nil
}

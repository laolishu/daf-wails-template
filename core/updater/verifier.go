package updater

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"
)

type Verifier interface {
	Verify(ctx context.Context, info UpdateInfo, file DownloadResult) error
}

type SHA256Verifier struct{}

func (SHA256Verifier) Verify(ctx context.Context, info UpdateInfo, file DownloadResult) error {
	logger := slog.Default()
	start := time.Now()
	logger.Info("updater.verify.start", "path", file.Path)

	expected := strings.TrimSpace(info.Checksum)
	if expected == "" {
		logger.Error("updater.verify.failed", "path", file.Path, "error", "checksum is empty")
		return fmt.Errorf("checksum is empty")
	}
	if strings.TrimSpace(file.Path) == "" {
		logger.Error("updater.verify.failed", "path", file.Path, "error", "download path is empty")
		return fmt.Errorf("download path is empty")
	}

	actual := strings.TrimSpace(file.Checksum)
	if actual == "" {
		var err error
		actual, err = computeFileSHA256(file.Path)
		if err != nil {
			logger.Error("updater.verify.failed", "path", file.Path, "error", err)
			return err
		}
	}

	normalizedExpected := strings.ToLower(expected)
	normalizedActual := strings.ToLower(actual)
	if normalizedExpected != normalizedActual {
		logger.Error("updater.verify.mismatch", "path", file.Path, "expected", normalizedExpected, "actual", normalizedActual)
		return fmt.Errorf("checksum mismatch")
	}
	logger.Info("updater.verify.success", "path", file.Path, "duration", time.Since(start))
	return nil
}

func computeFileSHA256(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

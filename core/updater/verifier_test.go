package updater

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
)

func TestSHA256Verifier(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "payload.bin")
	payload := []byte("hello-updater")
	if err := os.WriteFile(path, payload, 0o644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	hash := sha256.Sum256(payload)
	checksum := hex.EncodeToString(hash[:])

	verifier := SHA256Verifier{}
	if err := verifier.Verify(routineContext(), UpdateInfo{Checksum: checksum}, DownloadResult{Path: path}); err != nil {
		t.Fatalf("verify: %v", err)
	}
}

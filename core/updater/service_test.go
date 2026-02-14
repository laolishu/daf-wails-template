package updater

import (
	"context"
	"testing"
)

type stubProvider struct {
	info UpdateInfo
}

func (s stubProvider) Check(ctx context.Context, req UpdateRequest) (UpdateInfo, error) {
	return s.info, nil
}

type stubPolicy struct{}

func (stubPolicy) Decide(ctx context.Context, info UpdateInfo) (PolicyDecision, error) {
	return PolicyDecision{Allowed: true}, nil
}

type stubDownloader struct {
	result DownloadResult
	calls  *[]string
}

func (s stubDownloader) Download(ctx context.Context, info UpdateInfo) (DownloadResult, error) {
	if s.calls != nil {
		*s.calls = append(*s.calls, "download")
	}
	return s.result, nil
}

type stubVerifier struct {
	calls *[]string
}

func (s stubVerifier) Verify(ctx context.Context, info UpdateInfo, file DownloadResult) error {
	if s.calls != nil {
		*s.calls = append(*s.calls, "verify")
	}
	return nil
}

type stubInstaller struct {
	result InstallResult
	calls  *[]string
}

func (s stubInstaller) Install(ctx context.Context, info UpdateInfo, file DownloadResult) (InstallResult, error) {
	if s.calls != nil {
		*s.calls = append(*s.calls, "install")
	}
	return s.result, nil
}

func TestUpdaterRunOrder(t *testing.T) {
	calls := make([]string, 0, 3)
	updater := &Updater{
		Provider:   stubProvider{info: UpdateInfo{Checksum: "abc"}},
		Policy:     stubPolicy{},
		Downloader: stubDownloader{result: DownloadResult{Path: "file", Checksum: "abc"}, calls: &calls},
		Verifier:   stubVerifier{calls: &calls},
		Installer:  stubInstaller{result: InstallResult{Path: "path"}, calls: &calls},
	}

	if _, err := updater.Run(context.Background(), UpdateRequest{}); err != nil {
		t.Fatalf("run: %v", err)
	}

	if len(calls) != 3 {
		t.Fatalf("unexpected call count: %d", len(calls))
	}
	if calls[0] != "download" || calls[1] != "verify" || calls[2] != "install" {
		t.Fatalf("unexpected call order: %v", calls)
	}
}

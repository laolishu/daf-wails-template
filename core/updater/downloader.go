package updater

import "context"

type Downloader interface {
	Download(ctx context.Context, info UpdateInfo) (DownloadResult, error)
}

# Remote Update Module

## Purpose
The remote updater lives in `core/updater` and provides interface-driven update checks, downloads, verification, and extraction:
- `UpdateProvider` handles update checks (default `HTTPProvider`).
- `Updater` orchestrates check, policy, download, verify, install.
- `BasicInstaller` downloads a ZIP and extracts it to a temp directory; it does not replace the running process.

## Usage
### 1. Configure update endpoint
Inject `UpdateEndpoint` via sysconfig.

### 2. Check for updates
`App.CheckForUpdate` builds an `UpdateRequest` and calls `HTTPProvider.Check`:
- Method: `POST`
- Content-Type: `application/json`
- Fields: `appId`, `version`, `platform`, `arch`, `channel`, `language` (optional)

### 3. Download and install
`App.DownloadUpdate` uses `Updater` to download and extract, returning the extraction path.

## Key Fields and APIs
- sysconfig: `UpdateEndpoint`
- Models: `core/updater/model.go` `UpdateRequest`/`UpdateInfo`
- Flow: `core/updater/service.go`, `core/updater/installer.go`

## Related Docs
- [System config setup](../sysconfig-setup.md)
- [Updater service](../../core/updater/service.go)
- [HTTP provider](../../core/updater/client.go)
- [Installer](../../core/updater/installer.go)

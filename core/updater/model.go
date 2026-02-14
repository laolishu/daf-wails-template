package updater

type UpdateInfo struct {
	ProtocolVersion     string `json:"protocolVersion"`
	LatestVersion       string `json:"latestVersion"`
	Force               bool   `json:"force"`
	Channel             string `json:"channel"`
	DownloadURL         string `json:"downloadUrl"`
	Checksum            string `json:"checksum"`
	ReleaseNotes        string `json:"releaseNotes"`
	MinSupportedVersion string `json:"minSupportedVersion"`
}

// UpdateRequest 升级检查请求参数
type UpdateRequest struct {
	AppID    string `json:"appId"`              // AppID 应用标识符
	Version  string `json:"version"`            // Version 当前版本号
	Platform string `json:"platform"`           // Platform 运行平台（如 windows、darwin、linux）
	Arch     string `json:"arch"`               // Arch 架构（如 amd64、arm64）
	Channel  string `json:"channel"`            // Channel 升级通道（如 stable、beta）
	Language string `json:"language,omitempty"` // Language 客户端语言（如 zh-CN、en-US，可选）
}

type DownloadResult struct {
	Path     string
	Size     int64
	Checksum string
}

type InstallResult struct {
	Path string
}

type PolicyDecision struct {
	Allowed bool
	Reason  string
}

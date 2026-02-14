/*
 * @Descripttion:
 * @version:
 * @Author: lfzxs@qq.com
 * @Date: 2026-02-06 10:37:58
 * @LastEditors: lfzxs@qq.com
 * @LastEditTime: 2026-02-14 11:02:19
 */
package sysconfig

// 编译时通过 ldflags 注入的系统变量
// 使用示例: go build -ldflags "-X core/sysconfig.Version=v1.0.0 -X core/sysconfig.BuildTime=2026-02-04T10:00:00Z"

var (
	// Version 应用版本号（如 v1.0.0）
	Version = "dev"

	// BuildTime 构建时间（ISO 8601 格式）
	BuildTime = "unknown"

	// GitCommit Git 提交哈希（简称，如 abc1234）
	GitCommit = "unknown"

	// ConfigDir 配置文件目录
	ConfigDir = "config"

	// ConfigFile 配置文件名称
	ConfigFile = "config.yml"

	// UpdateEndpoint 升级检测服务地址（HTTP URL）
	// UpdateEndpoint = "https://api.example.com/update/check"
	UpdateEndpoint = "http://127.0.0.1:4523/m2/7838125-7586793-default/418614329"
)

// Info 返回所有系统配置信息的结构体
type Info struct {
	Version        string
	BuildTime      string
	GitCommit      string
	ConfigDir      string
	ConfigFile     string
	UpdateEndpoint string
}

// GetInfo 返回系统配置信息快照
func GetInfo() Info {
	return Info{
		Version:        Version,
		BuildTime:      BuildTime,
		GitCommit:      GitCommit,
		ConfigDir:      ConfigDir,
		ConfigFile:     ConfigFile,
		UpdateEndpoint: UpdateEndpoint,
	}
}

// GetVersion 返回应用版本号
func GetVersion() string {
	return Version
}

// GetBuildTime 返回构建时间
func GetBuildTime() string {
	return BuildTime
}

// GetGitCommit 返回 Git 提交哈希
func GetGitCommit() string {
	return GitCommit
}

// GetConfigDir 返回配置文件目录
func GetConfigDir() string {
	return ConfigDir
}

// GetConfigFile 返回配置文件名称
func GetConfigFile() string {
	return ConfigFile
}

// GetUpdateEndpoint 返回升级检测服务地址
func GetUpdateEndpoint() string {
	return UpdateEndpoint
}

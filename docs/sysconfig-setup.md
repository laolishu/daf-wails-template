# 系统配置模块（sysconfig）设置指南

系统配置模块用于管理编译时注入的全局系统变量，如版本号、构建时间、Git 提交哈希和配置目录。

## 概述

`core/sysconfig` 模块提供了一种标准化的方式来管理编译时确定的系统级配置。这些值通过 Go 编译时 ldflags 机制注入，运行时只读，与应用配置（core/config）相区分。

### 标准系统变量

- **Version**：应用版本号（如 `v1.0.0`）
- **BuildTime**：构建时间（ISO 8601 格式，如 `2026-02-04T10:00:00Z`）
- **GitCommit**：Git 提交哈希（简称，如 `abc1234`）
- **ConfigDir**：配置文件目录（如 `/etc/app` 或 `/home/user/.config/app`）
- **ConfigFile**：配置文件名称（如 `config.yml`），默认值为 `config.yml`

## 编译时注入

### 基本语法

使用 Go 的 `-ldflags` 标志在编译时注入系统变量：

```bash
go build -ldflags "-X core/sysconfig.Version=v1.0.0 -X core/sysconfig.BuildTime=2026-02-04T10:00:00Z -X core/sysconfig.GitCommit=abc1234 -X core/sysconfig.ConfigDir=/etc/app" -o app
```

### Makefile 示例

创建 `Makefile` 来自动化构建过程：

```makefile
VERSION ?= dev
BUILD_TIME := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT := $(shell git rev-parse --short HEAD || echo 'unknown')
CONFIG_DIR ?= /etc/daf-app

BUILD_FLAGS := -ldflags "-X core/sysconfig.Version=$(VERSION) -X core/sysconfig.BuildTime=$(BUILD_TIME) -X core/sysconfig.GitCommit=$(GIT_COMMIT) -X core/sysconfig.ConfigDir=$(CONFIG_DIR)"

.PHONY: build
build:
	go build $(BUILD_FLAGS) -o bin/app ./cmd/desktop

.PHONY: build-release
build-release:
	go build $(BUILD_FLAGS) -ldflags "-s -w" -o bin/app ./cmd/desktop
```

### Shell 脚本示例

使用 shell 脚本进行构建：

```bash
#!/bin/bash

VERSION=${1:-v1.0.0}
BUILD_TIME=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT=$(git rev-parse --short HEAD || echo 'unknown')
CONFIG_DIR=${2:-/etc/daf-app}

go build \
  -ldflags "-X core/sysconfig.Version=$VERSION -X core/sysconfig.BuildTime=$BUILD_TIME -X core/sysconfig.GitCommit=$GIT_COMMIT -X core/sysconfig.ConfigDir=$CONFIG_DIR" \
  -o app ./cmd/desktop
```

运行脚本：
```bash
./build.sh v1.0.0 /etc/app
```

## 使用方式

### 获取所有系统配置

```go
import "core/sysconfig"

func main() {
	info := sysconfig.GetInfo()
	fmt.Printf("Version: %s\n", info.Version)
	fmt.Printf("Build Time: %s\n", info.BuildTime)
	fmt.Printf("Git Commit: %s\n", info.GitCommit)
	fmt.Printf("Config Dir: %s\n", info.ConfigDir)
	fmt.Printf("Config File: %s\n", info.ConfigFile)
}
```

### 获取单个变量

```go
version := sysconfig.GetVersion()
buildTime := sysconfig.GetBuildTime()
gitCommit := sysconfig.GetGitCommit()
configDir := sysconfig.GetConfigDir()
configFile := sysconfig.GetConfigFile()
```

### 获取完整配置文件路径

```go
import (
	"path/filepath"
	"core/sysconfig"
)

// 组合配置目录和文件名
ConfigPath := filepath.Join(sysconfig.GetConfigDir(), sysconfig.GetConfigFile())
fmt.Printf("Full config path: %s\n", configPath)
// 输出：/etc/daf-app/config.yml
```

## 零值检查

系统变量的默认值如下：

| 变量 | 默认值 |
|------|--------|
| Version | `dev` |
| BuildTime | `unknown` |
| GitCommit | `unknown` |
| ConfigDir | （空字符串） |
| ConfigFile | `config.yml` |

在生产环境中，确保通过 ldflags 为所有变量注入合适的值。

## 与应用配置的区别

| 特性 | sysconfig | core/config |
|------|-----------|-----------|
| 注入时机 | 编译时 | 运行时 |
| 可修改性 | 只读 | 可修改 |
| 用途 | 系统级信息（版本、构建） | 应用行为配置（日志级别、插件开关） |
| 来源 | ldflags | 配置文件、环境变量 |

## CI/CD 集成示例（GitHub Actions）

```yaml
name: Build and Release

on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.20
      
      - name: Get version
        id: version
        run: |
          VERSION=$(git describe --tags --always)
          echo "version=$VERSION" >> $GITHUB_OUTPUT
      
      - name: Build
        run: |
          BUILD_TIME=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
          GIT_COMMIT=$(git rev-parse --short HEAD)
          go build \
            -ldflags "-X core/sysconfig.Version=${{ steps.version.outputs.version }} -X core/sysconfig.BuildTime=$BUILD_TIME -X core/sysconfig.GitCommit=$GIT_COMMIT -X core/sysconfig.ConfigDir=/etc/app" \
            -o bin/app ./cmd/desktop
```

## 最佳实践

1. **始终注入版本和构建时间**：便于诊断和问题追踪
2. **使用 Git 提交哈希**：快速定位代码版本
3. **配置目录应遵循平台约定**：
   - Linux: `~/.config/app` 或 `/etc/app`
   - macOS: `~/Library/Application Support/app`
   - Windows: `%APPDATA%\app`
4. **在测试时使用合理的默认值**：开发过程中可用 `dev` 和 `unknown`
5. **文档化变量含义**：在 Makefile 或构建脚本中添加注释

## 故障排除

### 变量未被注入

检查 ldflags 语法，确保使用 `-X` 而不是 `-x`：

```bash
# ❌ 错误
go build -ldflags "-x core/sysconfig.Version=v1.0.0"

# ✅ 正确
go build -ldflags "-X core/sysconfig.Version=v1.0.0"
```

### 包路径错误

确保包路径与 Go module 名称一致：

```bash
# 如果 go.mod 中是 module example.com/app
# 则应使用完整路径
go build -ldflags "-X example.com/app/core/sysconfig.Version=v1.0.0"
```

### 时间格式问题

使用 `-u` 确保 UTC 时间：

```bash
# ✅ 正确（UTC）
date -u +'%Y-%m-%dT%H:%M:%SZ'

# ❌ 可能有时区问题
date +'%Y-%m-%dT%H:%M:%S'
```

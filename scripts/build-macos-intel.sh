#!/bin/bash

# DAF-Wails macOS 构建脚本（Intel/AMD64）
# 支持自动注入系统配置（版本、构建时间、Git 提交、配置目录）

set -e

# 参数配置（可通过环境变量或命令行参数覆盖）
VERSION="${1:-${VERSION:-dev}}"
BUILD_TIME="${BUILD_TIME:-$(date -u +'%Y-%m-%dT%H:%M:%SZ')}"
GIT_COMMIT="${GIT_COMMIT:-$(git rev-parse --short HEAD 2>/dev/null || echo 'unknown')}"
CONFIG_DIR="${CONFIG_DIR:-~/Library/Application Support/daf-app}"

echo "======================================"
echo "DAF-Wails Build Script (macOS Intel)"
echo "======================================"
echo "Version:     $VERSION"
echo "Build Time:  $BUILD_TIME"
echo "Git Commit:  $GIT_COMMIT"
echo "Config Dir:  $CONFIG_DIR"
echo "======================================"

cd "$(dirname "$0")/.."

# 构建 ldflags
LDFLAGS="-X {{.ModulePath}}/core/sysconfig.Version=$VERSION"
LDFLAGS="$LDFLAGS -X {{.ModulePath}}/core/sysconfig.BuildTime=$BUILD_TIME"
LDFLAGS="$LDFLAGS -X {{.ModulePath}}/core/sysconfig.GitCommit=$GIT_COMMIT"
LDFLAGS="$LDFLAGS -X {{.ModulePath}}/core/sysconfig.ConfigDir=$CONFIG_DIR"

echo "Building application for macOS (Intel)..."
wails build --clean --platform darwin -ldflags "$LDFLAGS"

echo ""
echo "Build completed successfully!"
echo "Binary location: build/bin/"

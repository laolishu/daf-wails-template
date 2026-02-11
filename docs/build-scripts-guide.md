# 构建脚本使用指南

DAF-Wails 提供了一套完整的跨平台构建脚本，自动注入系统配置信息（版本号、构建时间、Git 提交、配置目录）到最终的二进制文件中。

## 可用脚本

### Linux/Unix 通用构建
```bash
./scripts/build.sh [VERSION]
```

### Windows 构建
```bat
scripts\build-windows.bat
```

### macOS 构建
```bash
# Universal Binary (ARM64 + Intel)
./scripts/build-macos.sh [VERSION]

# ARM64 专用
./scripts/build-macos-arm.sh [VERSION]

# Intel 专用
./scripts/build-macos-intel.sh [VERSION]
```

## 参数说明

所有构建脚本支持以下参数配置：

| 参数 | 说明 | 默认值 | 传递方式 |
|------|------|--------|----------|
| VERSION | 应用版本号 | `dev` | 命令行参数或环境变量 |
| BUILD_TIME | 构建时间（ISO 8601） | 当前 UTC 时间 | 环境变量 |
| GIT_COMMIT | Git 提交哈希（简称） | 当前分支最新提交 | 环境变量 |
| CONFIG_DIR | 配置文件目录 | 平台默认路径 | 环境变量 |

### 平台默认配置目录

- **Linux/Unix**: `/etc/daf-app`
- **macOS**: `~/Library/Application Support/daf-app`
- **Windows**: `C:\ProgramData\daf-app`

## 使用示例

### 开发构建（使用默认值）

```bash
# Linux/macOS
./scripts/build.sh

# Windows
scripts\build-windows.bat
```

构建后的二进制将包含：
- Version: `dev`
- BuildTime: 当前 UTC 时间
- GitCommit: 当前提交哈希
- ConfigDir: 平台默认值

### 发布构建（指定版本）

```bash
# 方式 1：命令行参数
./scripts/build.sh v1.0.0

# 方式 2：环境变量
VERSION=v1.0.0 ./scripts/build.sh
```

### 完全自定义

```bash
VERSION=v1.2.3 \
BUILD_TIME=2026-02-04T10:00:00Z \
GIT_COMMIT=abc1234 \
CONFIG_DIR=/opt/myapp \
./scripts/build.sh
```

### Windows 示例

```bat
set VERSION=v1.0.0
set CONFIG_DIR=C:\MyApp
scripts\build-windows.bat
```

## 验证注入结果

构建完成后，可以通过以下方式验证系统配置是否正确注入：

### 方法 1：运行应用并查看日志

应用启动时会通过 `sysconfig.GetInfo()` 获取系统配置信息。

### 方法 2：使用 strings 命令（Unix/Linux/macOS）

```bash
strings build/bin/your-app | grep -E "v[0-9]+\.[0-9]+\.[0-9]+"
```

### 方法 3：PowerShell（Windows）

```powershell
Select-String -Path build\bin\your-app.exe -Pattern "v\d+\.\d+\.\d+"
```

## CI/CD 集成

### GitHub Actions 示例

```yaml
name: Build Release

on:
  push:
    tags:
      - 'v*'

jobs:
  build-linux:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.21
      
      - name: Set up Wails
        run: go install github.com/wailsapp/wails/v2/cmd/wails@latest
      
      - name: Build
        run: |
          VERSION=${GITHUB_REF#refs/tags/}
          ./scripts/build.sh $VERSION
      
      - name: Upload artifact
        uses: actions/upload-artifact@v2
        with:
          name: linux-build
          path: build/bin/

  build-windows:
    runs-on: windows-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.21
      
      - name: Set up Wails
        run: go install github.com/wailsapp/wails/v2/cmd/wails@latest
      
      - name: Build
        run: |
          $env:VERSION = $env:GITHUB_REF -replace 'refs/tags/', ''
          scripts\build-windows.bat
      
      - name: Upload artifact
        uses: actions/upload-artifact@v2
        with:
          name: windows-build
          path: build\bin\

  build-macos:
    runs-on: macos-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.21
      
      - name: Set up Wails
        run: go install github.com/wailsapp/wails/v2/cmd/wails@latest
      
      - name: Build
        run: |
          VERSION=${GITHUB_REF#refs/tags/}
          ./scripts/build-macos.sh $VERSION
      
      - name: Upload artifact
        uses: actions/upload-artifact@v2
        with:
          name: macos-build
          path: build/bin/
```

## 故障排除

### 问题：版本信息未注入

**原因**：ldflags 语法错误或模块路径不匹配。

**解决**：检查 `go.mod` 中的模块路径是否与脚本中的 `{{.ModulePath}}` 一致。

### 问题：BUILD_TIME 格式错误

**原因**：`date` 命令参数在不同系统上可能有差异。

**解决**：
- macOS/Linux：使用 `date -u +'%Y-%m-%dT%H:%M:%SZ'`
- Windows（PowerShell）：使用 `Get-Date -Format 'yyyy-MM-ddTHH:mm:ssZ' -AsUTC`

### 问题：Git 提交哈希为 'unknown'

**原因**：不在 Git 仓库中或 Git 未安装。

**解决**：
1. 确保项目是 Git 仓库
2. 确保已安装 Git 并可在命令行访问
3. 或手动设置：`GIT_COMMIT=abc1234 ./scripts/build.sh`

### 问题：Windows 脚本无法执行

**原因**：PowerShell 执行策略限制。

**解决**：
```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

## 最佳实践

1. **开发阶段**：直接运行脚本，使用默认值（`dev` 版本）
2. **测试阶段**：指定版本号，如 `v1.0.0-beta.1`
3. **发布阶段**：使用语义化版本（Semantic Versioning），如 `v1.0.0`
4. **CI/CD**：从 Git 标签自动提取版本号
5. **配置目录**：生产环境应明确指定配置目录，避免使用默认值
6. **文档化**：在 README 中说明版本号规范与构建流程

## 相关文档

- [系统配置模块文档](./sysconfig-setup.md)
- [Wails 官方文档](https://wails.io/)
- [Go ldflags 详解](https://pkg.go.dev/cmd/link)

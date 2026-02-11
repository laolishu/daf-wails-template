# 添加应用配置模块（core/config）

## 背景
当前项目已有 `core/sysconfig`，用于管理编译时注入的**系统配置**（版本、构建时间等）。但缺少**应用配置模块**，用于管理运行时可修改的用户配置，如日志级别、日志目录、国际化语言等。

应用配置（core/config）与系统配置（core/sysconfig）的区别：
- **系统配置**：编译时确定、运行时只读、系统级信息（版本、构建、ConfigDir）
- **应用配置**：运行时加载、用户可修改、应用行为配置（日志级别、插件开关等）

本提案创建 `core/config` 模块，使用 `viper` 库实现配置的加载、保存与监听，并基于 `sysconfig.ConfigDir` 和 `sysconfig.ConfigFile` 确定配置文件路径。

## 目标
1. 创建 `core/config` 模块，提供统一的应用配置管理接口
2. 使用 `viper` 库实现配置文件的读写与监听
3. 定义核心应用配置项：
   - 日志配置：LogDir（日志目录）、LogLevel（日志级别）
   - 国际化配置：Language（语言）
4. 配置文件路径基于 `sysconfig.GetConfigDir()` 和 `sysconfig.GetConfigFile()`
5. 支持配置文件不存在时使用默认值并自动创建
6. 支持配置文件热重载（可选）

## 非目标
- 本提案不实现 logger 模块，仅定义日志相关的配置项
- 本提案不实现 i18n 模块，仅定义语言配置项
- 不包含配置加密或高级验证机制
- 不实现配置的远程同步或云端管理

## 设计概述

### 模块位置
- `core/config/config.go`：配置管理核心逻辑
- `core/config/config_test.go`：单元测试
- `core/config/defaults.go`：默认配置值定义

### 配置文件格式
使用 YAML 格式（与 `sysconfig.ConfigFile` 默认值 `config.yml` 一致）。

示例配置文件内容：
```yaml
# 日志配置
log:
  dir: "logs"           # 日志目录（相对路径或绝对路径）
  level: "info"         # 日志级别：debug / info / warn / error

# 国际化配置
i18n:
  language: "zh-CN"     # 语言：zh-CN / en-US / ja-JP 等
```

### 关键接口
```go
package config

import "github.com/spf13/viper"

// Manager 应用配置管理器
type Manager struct {
    v *viper.Viper
}

// NewManager 创建配置管理器
func NewManager() (*Manager, error)

// Load 加载配置文件
func (m *Manager) Load() error

// Save 保存配置文件
func (m *Manager) Save() error

// GetLogDir 获取日志目录
func (m *Manager) GetLogDir() string

// GetLogLevel 获取日志级别
func (m *Manager) GetLogLevel() string

// GetLanguage 获取语言
func (m *Manager) GetLanguage() string

// SetLogDir 设置日志目录
func (m *Manager) SetLogDir(dir string)

// SetLogLevel 设置日志级别
func (m *Manager) SetLogLevel(level string)

// SetLanguage 设置语言
func (m *Manager) SetLanguage(lang string)
```

### 配置文件路径解析
1. 从 `sysconfig.GetConfigDir()` 获取配置目录
2. 从 `sysconfig.GetConfigFile()` 获取配置文件名
3. 组合得到完整路径（使用 `filepath.Join`）
4. 若配置目录为空，使用当前工作目录
5. 若配置文件不存在，使用默认值并自动创建

### 默认值
```go
const (
    DefaultLogDir   = "logs"
    DefaultLogLevel = "info"
    DefaultLanguage = "zh-CN"
)
```

## 风险与依赖
### 依赖
- **sysconfig 模块**：需要 `sysconfig.GetConfigDir()` 和 `sysconfig.GetConfigFile()` 获取配置文件路径
- **viper 库**：需要在 `go.mod` 中添加 `github.com/spf13/viper` 依赖

### 风险
- 配置文件权限问题：若配置目录不可写，保存配置将失败
- 配置文件格式错误：需要处理 YAML 解析错误
- 并发安全：多模块同时读写配置时需要加锁（本提案暂不涉及，留待后续优化）

## 验收标准
1. `core/config` 模块实现完成，包含 `Manager` 结构体和相关方法
2. 支持从 `sysconfig` 获取配置文件路径
3. 支持加载 YAML 配置文件，不存在时使用默认值
4. 支持保存配置文件
5. 单元测试覆盖核心功能（加载、保存、读取、设置）
6. 文档说明配置文件格式与使用方式
7. 通过 `openspec-cn validate --strict`

## 相关规范
- **sysconfig 规范**：本模块依赖 `sysconfig.ConfigDir` 和 `sysconfig.ConfigFile`
- **project-structure 规范**：本模块位于 `core/config`，符合核心能力层定位

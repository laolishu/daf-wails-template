# 应用配置模块（core/config）使用指南

本模块用于管理运行时可修改的应用配置，采用 YAML 格式并使用 viper 进行读写。配置文件路径由系统配置决定（`sysconfig.ConfigDir` + `sysconfig.ConfigFile`）。

## 配置文件示例

```yaml
log:
  dir: "logs"
  level: "info"
  retention_days: 7

i18n:
  language: "zh-CN"

window:
  width: 1024
  height: 768
  title: "DAF Wails App"
```

## 配置项说明

| 配置项 | 类型 | 默认值 | 说明 |
| --- | --- | --- | --- |
| log.dir | string | `logs` | 日志目录（相对或绝对路径） |
| log.level | string | `info` | 日志级别（debug/info/warn/error） |
| log.retention_days | int | `7` | 日志保留天数（按天清理） |
| i18n.language | string | `zh-CN` | 语言代码（如 zh-CN/en-US/ja-JP） |
| window.width | int | `1024` | 窗口宽度（像素） |
| window.height | int | `768` | 窗口高度（像素） |
| window.title | string | `DAF Wails App` | 窗口标题 |

## 配置文件路径

配置文件路径由 `core/sysconfig` 提供：

- 目录：`sysconfig.GetConfigDir()`
- 文件名：`sysconfig.GetConfigFile()`
- 最终路径：`filepath.Join(ConfigDir, ConfigFile)`

当 `ConfigDir` 为空时，使用当前工作目录。

## 使用示例

```go
manager, err := config.NewManager()
if err != nil {
	return err
}

if err := manager.Load(); err != nil {
	return err
}

logDir := manager.GetLogDir()
logLevel := manager.GetLogLevel()
retentionDays := manager.GetLogRetentionDays()
language := manager.GetLanguage()
windowWidth := manager.GetWindowWidth()
windowHeight := manager.GetWindowHeight()
windowTitle := manager.GetWindowTitle()

manager.SetLogLevel("debug")
manager.SetLogRetentionDays(7)
manager.SetWindowWidth(1280)
manager.SetWindowHeight(720)
manager.SetWindowTitle("DAF Wails App")
if err := manager.Save(); err != nil {
	return err
}
```

## 常见问题

### 1. 配置文件不存在怎么办？
当配置文件不存在时，系统会使用默认值并自动创建配置文件。

### 2. 配置文件格式错误
YAML 解析失败会返回明确错误，请检查缩进与字段名称。

### 3. 配置目录不可写
保存配置时会返回权限错误，请确保配置目录具备写权限。

# 配置模块

## 功能定位
配置体系分为两层：
- **系统配置（core/sysconfig）**：编译时注入的只读变量，用于版本、构建信息、配置目录与升级端点等系统级数据。
- **应用配置（core/config）**：运行时可读写的 YAML 配置，用于日志、语言与窗口等应用行为配置。

## 主要使用方式
### 系统配置（sysconfig）
- 通过 `-ldflags` 注入，运行时只读。
- 常用入口：`sysconfig.GetInfo()`、`sysconfig.GetUpdateEndpoint()`。

### 应用配置（config）
- 通过 `config.NewManager()` 创建并加载配置文件。
- 读取或更新后调用 `Save()` 进行持久化。

```go
manager, err := config.NewManager()
if err != nil {
	return err
}

if err := manager.Load(); err != nil {
	return err
}

logLevel := manager.GetLogLevel()
manager.SetLanguage("zh-CN")
if err := manager.Save(); err != nil {
	return err
}
```

## 关键配置或接口
- sysconfig 变量：`Version`、`BuildTime`、`GitCommit`、`ConfigDir`、`ConfigFile`、`UpdateEndpoint`
- appconfig 关键项：
  - `log.dir`、`log.level`、`log.retention_days`
  - `i18n.language`
  - `window.width`、`window.height`、`window.title`

## 相关文档
- [系统配置模块（sysconfig）设置指南](../sysconfig-setup.md)
- [应用配置模块（core/config）使用指南](../appconfig-guide.md)
- [sysconfig 规范](../../openspec/specs/sysconfig/spec.md)

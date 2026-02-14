# 日志模块

## 功能定位
日志模块基于 `slog` 提供统一日志接口，支持按天文件输出、保留策略与控制台输出，并提供 Wails 日志适配。

## 主要使用方式
- 在启动阶段通过 `logger.Init` 初始化日志，并将配置来源于应用配置。
- 在业务中使用 `slog.Default()` 或带上下文的 `slog.Logger` 写入日志。
- 如需接入 Wails 日志，使用 `logger.NewWailsAdapter` 作为适配器。

```go
_, err := logger.Init(logger.Config{
	Dir:            configManager.GetLogDir(),
	Level:          configManager.GetLogLevel(),
	RetentionDays:  configManager.GetLogRetentionDays(),
	ConsoleEnabled: true,
})
if err != nil {
	return err
}

slog.Info("app.start")
```

## 关键配置或接口
- `logger.Config`：`Dir`、`Level`、`RetentionDays`、`ConsoleEnabled`
- 应用配置关键项：`log.dir`、`log.level`、`log.retention_days`
- Wails 适配：`logger.NewWailsAdapter`

## 相关文档
- [应用配置模块（core/config）使用指南](../appconfig-guide.md)
- [日志模块实现](../../core/logger/logger.go)
- [Wails 日志适配器](../../core/logger/wails_adapter.go)

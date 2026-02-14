# Logging Module

## Purpose
The logging module uses `slog` to provide a unified logging interface, daily file output, retention, console logging, and a Wails adapter.

## Usage
- Initialize logging at startup with `logger.Init`, sourcing values from app config.
- Use `slog.Default()` or a scoped `slog.Logger` for writes.
- Use `logger.NewWailsAdapter` to pipe Wails logs into core logging.

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

## Key Fields and APIs
- `logger.Config`: `Dir`, `Level`, `RetentionDays`, `ConsoleEnabled`
- App config keys: `log.dir`, `log.level`, `log.retention_days`
- Wails adapter: `logger.NewWailsAdapter`

## Related Docs
- [App config guide](../appconfig-guide.md)
- [Logger implementation](../../core/logger/logger.go)
- [Wails adapter](../../core/logger/wails_adapter.go)

# Configuration Module

## Purpose
The configuration system has two layers:
- **System config (core/sysconfig)**: compile-time, read-only variables for version/build info, config path, and update endpoint.
- **App config (core/config)**: runtime YAML config for logging, language, and window behavior.

## Usage
### System config (sysconfig)
- Injected via `-ldflags`, read-only at runtime.
- Common entry points: `sysconfig.GetInfo()`, `sysconfig.GetUpdateEndpoint()`.

### App config (config)
- Use `config.NewManager()` to create and load the YAML file.
- Call `Save()` after updates.

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

## Key Fields and APIs
- sysconfig variables: `Version`, `BuildTime`, `GitCommit`, `ConfigDir`, `ConfigFile`, `UpdateEndpoint`
- app config keys:
  - `log.dir`, `log.level`, `log.retention_days`
  - `i18n.language`
  - `window.width`, `window.height`, `window.title`

## Related Docs
- [System config setup](../sysconfig-setup.md)
- [App config guide](../appconfig-guide.md)
- [Sysconfig spec](../../openspec/specs/sysconfig/spec.md)

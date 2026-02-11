# 变更：添加日志模块

## 为什么
当前项目缺少统一的日志模块，导致日志输出、分割与保留策略难以标准化。需要新增 core/logger 以提供一致的日志能力，并由应用初始化阶段从 core/config 读取配置后进行初始化。

## 变更内容
- 新增 core/logger 模块，基于 Go 标准库 slog 实现日志输出。
- 支持按日期分割日志文件，文件名格式为 YYYY-MM-DD.log，使用本地时区。
- 支持日志保留策略（默认保留 7 天），清理由模块执行。
- 同时输出到文件与控制台，日志级别一致。
- 在 backend/app.go 初始化流程中从 core/config 读取 log.dir、log.level、log.retention_days，并调用 logger 初始化。
- 扩展 appconfig 规范，新增 log.retention_days 默认值为 7。

## 影响
- 受影响规范：logger（新增）、appconfig（修改）
- 受影响代码：core/logger、backend/app.go、core/config 默认项（后续实施阶段）

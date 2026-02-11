# Design Doc：Wails 日志适配器

## 目标
将 Wails 日志输出统一接入 core/logger，确保日志文件与控制台输出保持一致。

## 设计要点
1. **接口适配**
   - 实现 `github.com/wailsapp/wails/v2/pkg/logger.Logger` 接口。
   - 适配器内部调用 `slog`（core/logger 初始化后的默认 logger）。

2. **级别映射**
   - TRACE → slog.Debug
   - DEBUG → slog.Debug
   - INFO → slog.Info
   - WARNING → slog.Warn
   - ERROR → slog.Error
   - FATAL → slog.Error + 退出进程（以保持 Wails 的 fatal 语义）

3. **接入点**
   - 在 `options.App` 中设置 `Logger` 字段为适配器实例。
   - 使用与 core/logger 相同的输出目录与级别配置。

## 影响范围
- backend：Wails 应用初始化时注入 Logger
- core/logger：新增适配器实现

# 变更：添加 Wails 日志适配器

## 为什么
当前 Wails 的内部日志未接入 core/logger，导致运行期日志分散在控制台，难以统一归档与检索。需要提供一个 Wails 日志适配器，将 Wails 日志路由到现有日志系统，保证输出一致性与可观测性。

## 变更内容
- 新增 Wails 日志适配器，实现 `wails/pkg/logger.Logger`
- 将 Wails 日志路由到 core/logger（slog）
- 统一日志级别映射（TRACE/DEBUG/INFO/WARNING/ERROR/FATAL）

## 影响
- 受影响规范：logger
- 受影响代码（后续实施阶段）：backend/app.go、core/logger（新增适配器）

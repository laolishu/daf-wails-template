# 设计：日志模块（core/logger）

## 背景与约束
- core/logger 必须与 Wails 解耦，仅依赖标准库。
- 初始化由应用层完成：在 backend/app.go 中读取 core/config，再将参数传入 logger。
- 模块仅负责接收参数并执行日志输出、分割与保留策略。

## 方案概述
- 使用 Go 标准库 slog 作为统一日志 API。
- 提供简单的初始化入口（例如 Init 或 New），接收如下参数：
  - LogDir：日志目录
  - Level：日志级别（debug/info/warn/error）
  - RetentionDays：保留天数
  - ConsoleEnabled：是否输出到控制台（此提案固定为开启）
- 输出目标为：文件 + 控制台，级别一致。

## 日志分割策略
- 按日期分割，每日一份日志。
- 文件命名：YYYY-MM-DD.log。
- 使用本地时区判断日期边界。
- 每次写入时判断当前日期与已打开文件是否一致，不一致则切换文件。

## 保留策略
- 在初始化时清理超过 RetentionDays 的日志文件。
- 仅在指定日志目录内匹配 YYYY-MM-DD.log 形式的文件进行清理。
- RetentionDays 为 0 或负数时，视为不清理（由实现阶段明确）。

## 与配置模块的关系
- core/config 提供：log.dir、log.level、log.retention_days。
- backend/app.go 负责读取配置并调用 logger 初始化。
- logger 不直接依赖 core/config，保持可复用与解耦。

## 未涵盖范围
- 不涉及日志结构化字段的具体规范。
- 不涉及多级别输出到不同目标。
- 不涉及异步队列或采样策略。

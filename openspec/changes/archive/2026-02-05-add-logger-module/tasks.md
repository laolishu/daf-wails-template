# 任务清单：添加日志模块

## 1. 规范与设计
- [x] 建立 logger 规范增量（含分割、保留、输出目标）
- [x] 更新 appconfig 规范增量（新增 log.retention_days 默认值）
- [x] 校验提案与规范通过 openspec-cn validate

## 2. 实施（批准后）
- [x] 实现 core/logger（基于 slog，支持每日分割与保留清理）
- [x] 在 backend/app.go 中读取 core/config 并初始化 logger
- [x] 更新 core/config 默认值与文档（加入 log.retention_days）

## 3. 验证
- [x] 添加/更新单元测试（日志分割、保留策略）
- [x] 运行相关测试与 lint（如有）

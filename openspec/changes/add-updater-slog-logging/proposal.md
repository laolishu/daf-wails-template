# 变更：updater 关键路径增加 slog 日志

## 为什么
升级检查与下载涉及外部请求与文件操作，目前缺少统一日志，排障与用户反馈成本高。

## 变更内容
- 在 core/updater 关键路径增加 slog 日志，覆盖检查、下载、校验与安装流程。
- 对外请求与返回状态、关键参数与结果以结构化字段记录。

## 影响
- 受影响规范：新增 updater
- 受影响代码：core/updater（client/service/installer/verifier 等）

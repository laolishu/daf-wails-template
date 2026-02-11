# 变更：新增 Examples 页面（真实系统调用示例）

## 为什么
当前缺少一个聚合式的系统调用示例页面，不便验证前端与 Wails API 的真实链路。新增 Examples 页面可提供标准化、可重复的验证入口。

## 变更内容
- 新增 Examples 页面，放在主导航，支持国际化显示名（默认英文 Examples）
- 示例采用卡片展示，包含 Logger / System Info / Config Read 三类真实调用
- Logger Example：前端调用 Wails 日志接口写入一次 INFO 日志，提示到日志路径查看
- System Info Example：读取版本号、构建时间、运行平台并只读展示
- Config Read Example：读取用户配置，展示语言与日志级别

## 影响
- 受影响规范：frontend-ui、frontend-shell、frontend-guidelines
- 受影响代码（后续实施阶段）：frontend/src/pages、frontend/src/services、backend/app.go

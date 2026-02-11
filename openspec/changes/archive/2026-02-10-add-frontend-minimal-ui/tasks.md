# 任务清单：前端最小 UI 与 System Info

## 1. 规范与设计
- [x] 起草 frontend-ui 规范增量（页面结构、路由、UI 约束）
- [x] 起草 core-system-info 规范增量（系统信息能力与 API 约束）
- [x] 校验提案与规范通过 openspec-cn validate

## 2. 实施（批准后）
- [x] 新增 core/system-info 能力并由 backend 暴露 Wails API
- [x] 前端实现 Health Check 与 System Info 两页（Mantine + Router）
- [x] 衔接前端 services 调用 Wails API 获取系统信息

## 3. 验证
- [x] 运行前端与后端相关测试
- [x] 手动验证 UI 路由与系统信息展示

# 设计：前端最小 UI + System Info

## 关键决策
- UI 采用 Mantine AppShell，左侧固定 Sidebar，主内容单一焦点。
- 路由使用 React Router，系统路径使用 /__system/* 命名。
- System Info 由 core 提供只读能力，backend 负责对外暴露 Wails API。

## System Info 接口边界
- core：提供系统信息数据结构与读取逻辑（版本、构建时间、运行环境等）。
- backend：将 core 能力暴露给前端，不在前端拼装或模拟数据。

## 风格约束
- 风格参考 CCleaner / 1Password：简洁、可信、克制。
- 禁止后台管理风格布局与表格。

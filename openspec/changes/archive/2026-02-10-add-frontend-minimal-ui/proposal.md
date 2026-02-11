# 变更：前端最小 UI 与 System Info 能力

## 为什么
当前前端仅有基础占位内容，缺少面向用户的最小可用界面与系统信息展示。需要新增两页最小 UI，并补齐系统信息获取能力，以满足桌面基础设施产品的可信展示与健康检查入口。

## 变更内容
- 新增前端最小 UI：Health Check（默认首页）与 About / System Info
- 使用 Mantine AppShell、左侧固定 Sidebar、主内容单一焦点
- 使用 React Router，系统页面路径为 /__system/health 与 /__system/about
- System Info 所需 Wails Go API 新增，能力落在 core 目录并由 backend 暴露
- 所有系统信息通过 Wails API 获取，前端不模拟数据

## 影响
- 受影响规范：frontend-ui（新增）、core-system-info（新增）
- 受影响代码：frontend、core、backend（后续实施阶段）

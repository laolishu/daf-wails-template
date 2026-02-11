# 变更：Settings 集成 System Info 与语言切换

## 为什么
当前 System Info 作为独立页面分散了用户入口，且缺少统一的设置管理能力。将 System Info 整合到 Settings，并提供中英文即时切换，有助于提升配置集中度与可用性。

## 变更内容
- 将 System Info 内容合并到 Settings 页面，不再作为独立主菜单页面
- 在 Settings 页面新增“语言”选项，置于第一项，支持中/英文切换
- 语言选项读取与写入后端用户配置项，并在切换后立即生效

## 影响
- 受影响规范：新增 frontend-settings 规范增量；参考 appconfig、frontend-guidelines
- 受影响代码：frontend/src/App.tsx、frontend/src/pages/Settings.tsx、frontend/src/pages/SystemInfo.tsx、frontend/src/services/*、backend/app.go、core/config/*

## 待确认
- 语言选项仅限 zh-CN/en-US（默认），不包含其它语言
- 语言切换即时生效的具体范围：仅当前前端文案，还是同时影响后端日志/诊断中的语言

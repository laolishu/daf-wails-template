# 变更：添加窗口尺寸配置

## 为什么
当前窗口宽度和高度在 backend/app.go 中硬编码为 1024x768，用户无法通过配置文件自定义窗口初始尺寸。需要将窗口尺寸配置化，使用户可通过配置文件调整。

## 变更内容
- 在 core/config 添加窗口配置项：window.width 和 window.height
- 设置默认值为 1024x768
- Wails 初始化时从 core/config 读取窗口宽高并应用
- 更新配置文件示例与文档

## 影响
- 受影响规范：appconfig（修改）
- 受影响代码：core/config（新增配置项）、backend/app.go（读取配置）

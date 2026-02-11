# 任务清单：添加窗口尺寸配置

## 1. 规范与设计
- [x] 更新 appconfig 规范增量（新增 window.width 和 window.height 配置项）
- [x] 校验提案与规范通过 openspec-cn validate

## 2. 实施（批准后）
- [x] 在 core/config 添加窗口尺寸配置项与默认值
- [x] 在 backend/app.go 中读取配置并应用到 Wails 初始化
- [x] 更新配置文件示例与文档

## 3. 验证
- [x] 添加/更新单元测试（窗口配置读写）
- [x] 手动验证窗口尺寸配置生效
- [x] 运行相关测试与 lint

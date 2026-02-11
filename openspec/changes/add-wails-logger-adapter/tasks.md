# Tasks：添加 Wails 日志适配器

## 提案阶段
- [x] 创建 proposal.md
- [x] 创建 design.md
- [x] 创建规范增量（spec.md）
- [x] 运行 `openspec-cn validate add-wails-logger-adapter --strict`

## 实施阶段（apply）
- [x] 新增 Wails Logger 适配器（实现 wails Logger 接口）
- [x] 将 Wails 日志接入 core/logger（slog）
- [x] 配置 backend 注入适配器到 options.App.Logger
- [x] 校验级别映射与输出一致性

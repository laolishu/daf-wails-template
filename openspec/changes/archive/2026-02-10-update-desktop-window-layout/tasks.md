# Tasks：桌面窗口尺寸与布局规范

## 提案阶段
- [x] 创建 proposal.md
- [x] 创建 design.md
- [x] 创建规范增量（spec.md）
- [x] 运行 `openspec-cn validate update-desktop-window-layout --strict`

## 实施阶段（apply）
- [x] 更新 Wails 窗口初始化配置（默认尺寸 1200×800，最小 1024×720，不强制最大化）
- [x] 调整前端 AppShell 布局尺寸（Sidebar 240px、Logo 64px、菜单项 40–44px、图标 18–20px、顶部栏 56–64px）
- [x] 校验主内容区宽度计算与自适应
- [x] 校验 125% 缩放下 UI 不溢出
- [x] 校验中文文本不换行挤压
- [x] 复核视觉与交互一致性

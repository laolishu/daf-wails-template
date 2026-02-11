# 变更：移除侧边栏顶部 Logo 与标题行

## 为什么
当前顶部标题栏已承担品牌展示与窗口拖动职责，Sidebar 顶部的 Logo 与标题行变为重复信息，压缩可用空间并降低导航聚焦。需要移除 Sidebar 顶部品牌行，保持导航区简洁。

## 变更内容
- 侧边栏不再显示顶部 Logo 与 title 行
- 侧边栏仅保留纵向导航与底部次级入口
- 不影响顶部标题栏（仍展示 Logo + window.title）

## 影响
- 受影响规范：frontend-shell、frontend-guidelines、frontend-ui
- 受影响代码（后续实施阶段）：frontend/src/App.tsx

# Design Doc：移除侧边栏顶部品牌行

## 目标
在新增顶部标题栏后，去除侧边栏顶部品牌行，减少重复信息并提升导航可视空间。

## 设计要点
1. **信息去重**
   - 品牌展示统一由顶部标题栏承担，Sidebar 只保留导航结构。

2. **布局简化**
   - Sidebar 不再包含 Logo 与 title 行，顶部直接进入导航列表。

3. **一致性**
   - 保持现有 Sidebar 配色与导航样式不变，仅移除品牌行。

## 影响范围
- 前端布局：移除 Sidebar 顶部品牌区
- 规范：更新 frontend-shell、frontend-guidelines、frontend-ui 中的对应要求

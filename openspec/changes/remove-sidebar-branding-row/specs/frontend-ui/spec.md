# frontend-ui 规范增量：移除 Sidebar 顶部品牌行

## 修改需求

### 需求：布局与交互
前端必须使用 Mantine AppShell，左侧固定 Sidebar，主内容单一焦点；Sidebar 仅包含导航项与底部次级入口，不包含品牌行。

#### 场景：渲染主布局
- **当** 渲染主界面
- **那么** 必须使用 AppShell 并包含固定 Sidebar 与单一焦点内容区
- **并且** Sidebar 不得显示 Logo 与标题行

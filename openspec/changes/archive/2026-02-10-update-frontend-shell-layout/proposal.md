# 变更：主页面布局与 Health Check 空状态

## 为什么
当前主页面与 Health Check 页面视觉层级不足，工具型应用的克制感与专业感不明确，需要统一为“左侧导航 + 单页焦点”的稳定布局。

## 变更内容
- 将主页面结构调整为固定 Sidebar + 主内容区，移除顶部复杂工具栏
- Sidebar 采用深色蓝灰背景，包含品牌区、纵向导航（图标+文本、选中高亮）与底部次级入口
- 品牌区 Logo 使用 public/logo.png，名称来自后端用户配置 window.title
- 底部次级入口对应 pages 目录下的页面
- 主内容区采用浅色背景、内容居中与大留白的“欢迎/空状态”结构
- Health Check 页面改为插画/占位图 + 标题 + 说明 + 单一主操作 CTA，不展示表格或复杂数据
- 视觉风格强调欧美工具型产品的可信与专业，避免炫酷动效与高饱和色彩

## 影响
- 受影响规范：新增 frontend-shell 规范增量；参考 frontend-guidelines
- 受影响代码：frontend/src/App.tsx，frontend/src/pages/HealthCheck.tsx，frontend/src/pages/SystemInfo.tsx

## 待确认
- 底部次级入口的最终文案与跳转目标（设置/帮助/升级）
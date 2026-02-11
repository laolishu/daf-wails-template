# 变更：Settings 栏目划分与语言行内操作

## 为什么
当前 Settings 缺少明确的栏目划分，语言选项以两行呈现，不利于快速浏览与操作。需要将设置内容分区并改为行内操作。

## 变更内容
- Settings 页面划分为“通用设置”和“系统信息”两个栏目
- 语言选项放在“通用设置”下，且改为单行显示与操作

## 影响
- 受影响规范：更新 frontend-settings 规范增量
- 受影响代码：frontend/src/pages/Settings.tsx

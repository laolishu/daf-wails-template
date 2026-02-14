# 变更：前端集成升级检查与系统配置地址

## 为什么
当前升级能力仅在 core/updater 层实现，前端缺少可用的升级检查入口，升级检测地址也无法通过系统配置统一管理，阻碍升级能力落地。

## 变更内容
- 将升级检测地址纳入 sysconfig 系统配置项，作为编译期注入的系统变量。
- 在 backend 暴露升级检查与下载安装 API，供前端调用。
- 在升级页面集成升级检查流程，弹出结果并提供下载/取消操作。

## 影响
- 受影响规范：sysconfig、frontend-ui、updater（新增）
- 受影响代码：core/sysconfig、backend、frontend

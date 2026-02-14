# 变更：最小可扩展远程升级架构

## 为什么
当前 core/updater 为空，缺少可扩展的远程升级架构，无法支撑后续 SaaS 升级与多策略扩展。

## 变更内容
- 在 core/updater 定义接口驱动的升级模块结构与最小实现范围。
- 规范远程升级协议（HTTP POST JSON）与 UpdateInfo 数据模型。
- 定义最小实现链路：HTTPProvider + SHA256Verifier + BasicInstaller。

## 影响
- 受影响规范：新增 updater
- 受影响代码：core/updater/*（新模块结构），未来可能影响 backend 适配层

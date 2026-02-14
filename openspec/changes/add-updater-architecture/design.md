## 上下文
当前 core/updater 为空，但项目目标要求支持本地/云升级与可扩展升级策略。需要建立最小可扩展架构，避免绑定具体 UI 或运行时。

## 目标 / 非目标
- 目标：提供接口驱动升级链路、最小 HTTP 协议与 SHA256 校验、可注入的组件编排。
- 目标：保持 core 与 Wails/UI 解耦，支持 headless。
- 非目标：实现自更新替换主程序、自动重启、签名校验、增量升级、灰度分发。

## 决策
- 采用接口驱动的模块化架构：UpdateProvider、UpdatePolicy、Downloader、Verifier、Installer。
- 升级流程由 Updater 服务编排，顺序固定为 Check -> Policy -> Download -> Verify -> Install。
- 协议采用 HTTP POST JSON，协议版本由 protocolVersion 字段声明。
- BasicInstaller 仅负责下载、校验、解压到临时目录并返回安装路径，不执行替换与重启。

## 考虑的替代方案
- 直接在 Updater 内部写死 HTTP 逻辑：被否决，违背接口优先与可扩展性。
- 仅提供下载能力，不包含校验/安装：被否决，最小升级链路不完整。

## 风险 / 权衡
- 风险：协议字段调整会影响客户端兼容性 → 通过 protocolVersion 进行显式版本化。
- 风险：安装器不替换主程序可能需要外部流程 → 由上层适配层或调用方接管。

## 迁移计划
- 新增 core/updater，不影响现有模块。

## 待决问题
- 后续是否引入签名验证或差分更新（不在本次范围）。

## 1. 提案与规范
- [x] 1.1 明确升级检测地址的 sysconfig 变量与注入方式。
- [x] 1.2 定义升级检查/下载的 backend API 与前端交互流程。
- [x] 1.3 补充升级页展示字段与弹窗行为要求。

## 2. 实施（待批准后）
- [x] 2.1 扩展 core/sysconfig 支持 UpdateEndpoint 变量与访问函数。
- [x] 2.2 在 backend 暴露升级检查与下载 API，并接入 core/updater。
- [x] 2.3 前端升级页集成检查与弹窗下载流程。
- [x] 2.4 添加最小测试或手动验证清单（API 返回、UI 交互）。

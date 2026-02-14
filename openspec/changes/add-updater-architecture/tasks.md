## 1. 提案与规范
- [x] 1.1 明确 updater 协议与数据模型规范（HTTP POST JSON + UpdateInfo）。
- [x] 1.2 定义接口驱动的模块责任与依赖注入边界。
- [x] 1.3 补充最小实现链路（HTTPProvider、SHA256Verifier、BasicInstaller）行为约束。

## 2. 实施（待批准后）
- [x] 2.1 创建 core/updater 目录结构与模型定义。
- [x] 2.2 实现接口与最小实现组件（provider/policy/downloader/verifier/installer）。
- [x] 2.3 实现 Updater 服务编排（Check -> Policy -> Download -> Verify -> Install）。
- [x] 2.4 添加最小单元测试覆盖关键路径（模型解析、校验/流程）。

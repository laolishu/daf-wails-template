## 1. 规范
- [x] 1.1 创建 appconfig 规范文件（openspec/changes/add-appconfig-module/specs/appconfig/spec.md）

## 2. 实施（审批后）
- [x] 2.1 更新 go.mod 添加 viper 依赖
- [x] 2.2 创建 core/config/defaults.go（默认配置值定义）
- [x] 2.3 创建 core/config/config.go（配置管理器实现）
- [x] 2.4 创建 core/config/config_test.go（单元测试）
- [x] 2.5 创建 docs/appconfig-guide.md（使用文档）
- [x] 2.6 创建示例配置文件 configs/app/config.yml

## 3. 验证
- [x] 3.1 单元测试通过（包含加载、保存、读取、设置功能）
- [x] 3.2 运行 openspec-cn validate add-appconfig-module --strict

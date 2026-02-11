## 1. 规范
- [x] 1.1 修改 sysconfig 规范，将 ConfigFile 添加到标准系统变量集

## 2. 实施（审批后）
- [x] 2.1 更新 core/sysconfig/sysconfig.go（添加 ConfigFile 变量与 GetConfigFile 函数）
- [x] 2.2 更新 core/sysconfig/sysconfig_test.go（添加 ConfigFile 测试）
- [x] 2.3 更新 docs/sysconfig-setup.md（文档中添加 ConfigFile 说明）
- [ ] 2.4 可选：更新构建脚本支持 ConfigFile 注入

## 3. 验证
- [x] 3.1 单元测试通过
- [x] 3.2 运行 openspec-cn validate add-sysconfig-configfile --strict

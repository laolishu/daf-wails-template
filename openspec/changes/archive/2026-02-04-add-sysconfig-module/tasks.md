## 1. 设计
- [x] 1.1 确定 sysconfig 模块的标准变量集
- [x] 1.2 设计与 core/config 的边界定义
- [x] 1.3 确认编译时 ldflags 注入方案示例

## 2. 规范
- [x] 2.1 起草 sysconfig 规范增量，列举所有标准系统变量
- [x] 2.2 定义 sysconfig 接口与行为需求

## 3. 实施（审批后）
- [x] 3.1 实现 core/sysconfig 模块（Info 结构体、初始化函数）
- [x] 3.2 更新 cmd/desktop main.go.tmpl（sysconfig 初始化示例）
- [x] 3.3 编写单元测试（变量获取、零值校验）
- [x] 3.4 添加 docs/sysconfig-setup.md（ldflags 注入示例与说明）

## 4. 验证
- [x] 4.1 编译时注入验证（make 脚本演示）
- [x] 4.2 单元测试通过
- [x] 4.3 运行 openspec-cn validate add-sysconfig-module --strict

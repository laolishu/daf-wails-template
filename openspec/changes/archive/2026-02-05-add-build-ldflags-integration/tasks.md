## 1. 设计
- [x] 1.1 确认 Wails build 命令是否支持 ldflags 传递
- [x] 1.2 确定跨平台脚本的参数化策略
- [x] 1.3 定义 ConfigDir 的平台默认值

## 2. 规范
- [x] 2.1 起草 sysconfig-build-integration 规范增量
- [x] 2.2 定义构建脚本的参数标准与 ldflags 注入规则

## 3. 实施（审批后）
- [x] 3.1 更新 build.sh（通用 Linux/Unix 构建）
- [x] 3.2 更新 build-windows.bat（Windows 构建）
- [x] 3.3 更新 build-macos.sh、build-macos-arm.sh、build-macos-intel.sh（macOS 专用）
- [x] 3.4 添加脚本使用文档与参数示例
- [x] 3.5 验证跨平台构建的 ldflags 正确性

## 4. 验证
- [x] 4.1 在各平台运行脚本并检查二进制文件中的 sysconfig 值
- [x] 4.2 单元测试（core/sysconfig 读取注入的值）
- [x] 4.3 运行 openspec-cn validate add-build-ldflags-integration --strict

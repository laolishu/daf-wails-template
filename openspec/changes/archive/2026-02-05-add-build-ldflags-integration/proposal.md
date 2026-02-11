# 变更：构建脚本集成系统配置 ldflags

## 为什么
当前构建脚本仅使用 `wails build` 命令，未将系统配置模块定义的编译时参数（版本号、构建时间、Git 提交等）注入到最终二进制文件中。这导致应用无法获取可靠的版本与构建信息，影响诊断、更新、日志记录等功能。

## 变更内容
- 更新构建脚本（build.sh、build-windows.bat、build-macos.sh、build-macos-arm.sh、build-macos-intel.sh），为 Go 编译器添加 ldflags 参数，注入 sysconfig 的四个标准变量。
- 统一构建脚本的参数化接口，支持自定义版本号、构建时间、Git 提交、配置目录。
- 提供清晰的使用文档与示例。

## 影响
- 受影响规范：sysconfig-build-integration
- 受影响代码：scripts/ 目录下所有构建脚本、Wails 配置可能需要调整以传递 ldflags
- 与 core/sysconfig 的关系：充分发挥 sysconfig 模块的作用，确保编译时信息能正确注入

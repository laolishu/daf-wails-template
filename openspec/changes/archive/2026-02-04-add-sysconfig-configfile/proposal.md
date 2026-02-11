# 变更：系统配置添加 ConfigFile 参数

## 为什么
当前系统配置模块定义了 ConfigDir（配置文件目录），但未指定默认的配置文件名称。应用需要知道默认配置文件的名称（如 config.yml），以便在初始化时正确加载配置。添加 ConfigFile 参数可以让系统配置提供完整的配置文件路径信息。

## 变更内容
- 在 core/sysconfig 模块添加 ConfigFile 变量，默认值为 `config.yml`
- 更新 sysconfig 规范，将 ConfigFile 纳入标准系统变量集
- 更新构建脚本，支持通过 ldflags 注入 ConfigFile（可选）
- 更新相关文档与测试

## 影响
- 受影响规范：sysconfig（修改标准系统变量集）
- 受影响代码：core/sysconfig/sysconfig.go、core/sysconfig/sysconfig_test.go、构建脚本（可选支持）
- 向后兼容：新增变量，不影响现有功能

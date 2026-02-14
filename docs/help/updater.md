# 远程升级模块

## 功能定位
远程升级模块位于 `core/updater`，以接口驱动的方式提供升级检查、下载、校验与解压能力：
- `UpdateProvider` 负责升级检查（默认 `HTTPProvider`）。
- `Updater` 负责串联检查、策略、下载、校验、安装流程。
- `BasicInstaller` 下载 ZIP 包并解压到临时目录，当前不负责替换正在运行的进程。

## 主要使用方式
### 1. 配置升级检查地址
升级检查地址通过 sysconfig 注入：`UpdateEndpoint`。

### 2. 进行升级检查
后端 `App.CheckForUpdate` 会构造 `UpdateRequest` 并调用 `HTTPProvider.Check`：
- 请求方法：`POST`
- Content-Type：`application/json`
- 关键字段：`appId`、`version`、`platform`、`arch`、`channel`、`language`（可选）

### 3. 下载并安装
后端 `App.DownloadUpdate` 使用 `Updater` 执行下载与解压，返回解压后的目录路径。

## 关键配置或接口
- sysconfig：`UpdateEndpoint`
- 升级请求：`core/updater/model.go` 中的 `UpdateRequest`/`UpdateInfo`
- 升级流程：`core/updater/service.go` 与 `core/updater/installer.go`

## 相关文档
- [系统配置模块（sysconfig）设置指南](../sysconfig-setup.md)
- [升级服务实现](../../core/updater/service.go)
- [HTTP 升级检查实现](../../core/updater/client.go)
- [安装器实现](../../core/updater/installer.go)

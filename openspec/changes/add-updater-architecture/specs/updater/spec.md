## 新增需求
### 需求：升级模块接口与模型
系统必须提供 core/updater 作为升级系统接口层，采用接口驱动设计且不依赖 UI/Wails。

#### 场景：新增升级能力
- **当** 开发者扩展升级策略或替换实现
- **那么** 必须通过接口与依赖注入完成替换而非修改核心流程

### 需求：UpdateInfo 数据模型
系统必须定义 UpdateInfo，并包含以下字段与 JSON 标签：
- protocolVersion
- latestVersion
- force
- channel
- downloadUrl
- checksum
- releaseNotes
- minSupportedVersion

#### 场景：解析远程升级响应
- **当** 客户端解析升级服务返回的 JSON
- **那么** 必须映射到 UpdateInfo 并保留全部字段

### 需求：远程升级协议
系统必须使用 HTTP POST JSON 进行升级检查，请求体包含 appId、version、platform、arch、channel，响应体包含 UpdateInfo 字段。

#### 场景：执行升级检查
- **当** 客户端发起升级检查
- **那么** 必须按照约定的 JSON 请求/响应格式进行通信

### 需求：升级执行链路
系统必须通过 Updater 服务按顺序执行 Check -> Policy -> Download -> Verify -> Install。

#### 场景：运行升级流程
- **当** Updater 执行 Run()
- **那么** 必须严格按顺序调用各模块并在失败时返回错误

### 需求：最小实现约束
系统必须提供最小实现组件：HTTPProvider、SHA256Verifier、BasicInstaller。

#### 场景：默认升级实现
- **当** 系统使用默认实现
- **那么** 必须具备 HTTP 检查、SHA256 校验与基础安装能力

### 需求：BasicInstaller 行为
BasicInstaller 必须下载文件、校验、解压到临时目录并返回安装路径，不得直接替换正在运行的主程序。

#### 场景：基础安装流程
- **当** 使用 BasicInstaller 安装升级包
- **那么** 必须按下载/校验/解压执行并返回安装路径

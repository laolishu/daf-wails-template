# project-structure Specification

## Purpose
TBD - created by archiving change add-project-structure. Update Purpose after archive.
## 需求
### 需求：项目根目录职责分层
项目必须使用固定的一级目录分层，并明确职责：
- /cmd：程序启动入口集合，仅负责启动，不包含业务逻辑，不直接访问 core。
- /backend：Wails 平台适配层，负责生命周期、窗口与事件绑定，不包含业务规则。
- /core：产品核心能力层，与 Wails 解耦，可复用到 CLI/Server/其他桌面框架。
- /frontend：前端 UI 壳，仅负责界面与交互，通过桥接层调用 backend API。
- /assets：桌面应用资源（icon、tray、启动图、平台资源）。
- /configs：默认配置模板（system/app/user），仅用于首次启动或 reset。
- /scripts：构建/发布脚本，不参与运行时代码。
- /internal：内部工具与非公开代码，遵循 Go internal 规则。
- /docs：产品与技术文档（架构、开发约定、商业与 License）。

#### 场景：新增模块放置
- **当** 开发者新增模块或文件
- **那么** 必须根据职责分层放置在对应一级目录中

### 需求：启动入口约束
/cmd/desktop 必须仅包含 main.go，且 main.go 只负责创建 App 并调用 Run()，不得初始化任何模块或业务逻辑。

#### 场景：实现桌面启动入口
- **当** 实现桌面应用启动入口
- **那么** 入口文件仅包含 App 创建与 Run() 调用

### 需求：backend 子目录职责
backend 必须包含以下职责明确的子目录或文件：
- /backend/lifecycle：OnStartup/OnShutdown/OnDomReady 的 Wails 生命周期适配，并转发给 core/runtime。
- /backend/bridge：前后端通信桥接（JS ↔ Go API、数据转换、统一错误）。
- /backend/runtime_adapter.go：Core Runtime 到 Wails 的适配层，负责注入 context 与事件回调，屏蔽 Wails 细节。

#### 场景：新增 Wails 适配逻辑
- **当** 需要新增生命周期或通信适配
- **那么** 必须放入 backend 的对应子目录或 runtime_adapter.go 中

### 需求：core 子模块职责
core 必须包含以下核心子模块并遵守职责：
- /core/runtime：启动顺序与模块生命周期管理。
- /core/config：分层配置（system/app/user），含合并与覆盖规则。
- /core/logger：统一日志接口，多输出与轮转策略。
- /core/i18n：语言包加载与自动语言检测，前后端共用 Key。
- /core/updater：升级系统接口层（本地/云升级/回滚）。
- /core/license：授权与 License 校验与扩展点。
- /core/plugin：插件生命周期、注册/卸载与上下文隔离。
- /core/diagnostics：运行诊断与诊断包导出。

#### 场景：新增核心能力
- **当** 新增核心能力模块
- **那么** 必须归入 core 对应子模块并保持与 Wails 解耦

### 需求：frontend 源码职责
/frontend/src 必须仅包含前端 UI 源码（组件与路由），禁止包含业务规则，且禁止直接操作 core。

#### 场景：新增 UI 功能
- **当** 新增前端页面或组件
- **那么** 代码必须位于 frontend/src 且通过 backend 桥接调用能力

### 需求：关键架构约束
系统必须遵守以下架构约束：
- main.go 只能启动，不写逻辑。
- backend 不写业务。
- core 不依赖 Wails。
- frontend 不直接操作 core。
- 商业能力必须位于 core。

#### 场景：代码评审
- **当** 进行代码评审
- **那么** 必须依据上述约束判断是否存在跨层依赖或职责越界


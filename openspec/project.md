# 项目上下文

## 目的
DAF-Wails（Desktop App Foundation for Wails）旨在为 Wails 桌面应用提供一套**标准化、可复用、可商业化的基础设施框架**，解决桌面应用在配置、日志、诊断、升级、运维与长期演进中的重复开发问题。

项目目标包括：
- 降低桌面应用的工程启动与维护成本
- 将“模板工程”升级为“可演进的产品级基础设施”
- 支持 Headless（无 UI）与 UI 共存
- 为商业化（插件 / SaaS / License）预留清晰扩展路径

---

## 技术栈
- **Go**（核心 Runtime / 基础能力实现）
- **Wails**（桌面应用宿主，当前以 Wails 2 为主，预留 Wails 3 适配）
- **Web 前端技术栈**（Base Shell UI，可选）
  - 推荐：TypeScript + 任意主流前端框架（React / Vue / Svelte 均可）
- **YAML / JSON**（配置文件格式）
- **ZIP**（诊断包格式）

---

## 项目约定

### 代码风格
- Go：
  - 遵循 `gofmt`
  - 显式错误处理，禁止 panic 作为正常流程
  - 接口优先（interface-first）设计
- 命名约定：
  - Core 模块使用清晰的领域名（config / logger / diagnostics）
  - 避免使用 runtime / utils 等泛化命名
- 禁止：
  - Core 中直接引用 Wails runtime 包
  - UI 中直接访问系统能力

---

### 架构模式
- **分层架构 + 适配器模式**
  - Core Runtime：纯 Go、无 UI、无 Wails 依赖
  - Runtime Adapter：隔离 Wails 2 / 3 差异
  - UI Shell：可选，仅通过 Core API 通信
- **插件化架构**
  - 生命周期驱动（Init / Start / Stop）
  - 插件配置与命名空间隔离
- **Headless First**
  - UI 不是产品成立的前提条件

---

### 测试策略
- Core Runtime：
  - 单元测试为主（业务逻辑、配置覆盖、诊断生成）
  - 不依赖 UI 或 Wails
- Adapter 层：
  - 最小集成测试
  - 验证接口完整性与隔离性
- UI：
  - 不作为强制测试目标
  - 重点验证与 Core API 的契约一致性

---

### Git 工作流
- 分支策略：
  - `main`：稳定可发布
  - `develop`：日常开发
  - `feature/*`：功能分支
- 提交规范：
  - 语义化提交（推荐）
    - `feat:` 新功能
    - `fix:` Bug 修复
    - `refactor:` 重构
    - `docs:` 文档
    - `chore:` 构建 / 工具
- 禁止：
  - 在同一个提交中混合架构变更与小修复

---

## 项目目录结构与职责

### 一级目录

#### `/cmd`
**用途：程序启动入口集合**
- 仅负责启动，不包含业务逻辑
- 不直接访问 core 子模块

#### `/backend`
**用途：Wails 平台适配层**
- 负责对接 Wails Runtime（生命周期、窗口、事件）
- 作为 core 与 Wails 的"胶水层"
- 禁止包含业务规则

#### `/core`
**用途：产品核心能力层**
- 所有"值钱"的能力都在这里
- 必须完全与 Wails 解耦
- 可复用到 CLI / Server / 其他桌面框架

#### `/frontend`
**用途：前端 UI 壳（可替换）**
- 仅负责界面与交互
- 不直接包含业务规则
- 通过桥接层调用 backend 暴露的 API

#### `/assets`
**用途：桌面应用资源**
- 图标（icon / tray）
- 启动画面
- 平台相关资源文件

#### `/configs`
**用途：默认配置模板**
- system / app / user 的默认值
- 仅用于首次启动或 reset
- 不存放运行时生成配置

#### `/scripts`
**用途：构建与发布脚本**
- build / release / sign
- CI/CD 辅助脚本
- 不参与运行时代码

#### `/internal`
**用途：内部工具与非公开代码**
- 仅供项目内部使用
- 不作为公共 API
- Go `internal` 规则生效

#### `/docs`
**用途：产品与技术文档**
- 架构说明
- 开发约定
- 商业与 License 说明

### 关键二级目录

#### `/cmd/desktop`
启动入口，仅包含 `main.go`，仅负责创建 App 并调用 Run()，不初始化任何模块。

#### `/backend/lifecycle`
Wails 生命周期适配（OnStartup / OnShutdown / OnDomReady），转发给 core/runtime。

#### `/backend/bridge`
前后端通信桥接（JS ↔ Go API 封装、数据结构转换、统一错误格式）。

#### `/backend/runtime_adapter.go`
Core Runtime 到 Wails 的适配层，负责注入 context、注入事件回调、屏蔽 Wails 具体实现细节。

#### `/core/runtime`
应用运行时中枢（控制启动顺序、管理模块生命周期、持有全局 Context）。

#### `/core/config`
配置系统（system / app / user 分层、配置合并与覆盖规则）。

#### `/core/logger`
日志系统（统一日志接口、多输出目标、日志级别与轮转策略）。

#### `/core/i18n`
国际化系统（语言包加载、自动语言检测、前后端共用 Key 规范）。

#### `/core/updater`
升级系统接口层（本地升级、云升级、回滚机制）。

#### `/core/license`
授权与 License（License 校验、授权类型判断、企业授权扩展点）。

#### `/core/plugin`
插件系统（插件生命周期、插件注册/卸载、插件隔离上下文）。

#### `/core/diagnostics`
运行诊断（状态快照、环境信息、一键导出诊断包）。

#### `/frontend/src`
前端源代码（UI 组件、页面路由、不存放业务规则）。

### 关键架构约束
- **main.go 只能启动，不写逻辑**
- **backend 不写业务**
- **core 不依赖 Wails**
- **frontend 不直接操作 core**
- **所有商业能力必须在 core**

---

## 领域上下文
- 本项目属于 **桌面应用基础设施 / Developer Tooling** 领域
- 主要关注：
  - 应用生命周期管理
  - 运维与问题诊断
  - 配置与环境一致性
- 与传统“最终用户桌面软件”不同：
  - 更偏向开发者 / 运维 / 企业 IT 使用
  - UI 以“运行控制台”为主，而非业务界面

---

## 重要约束
- **必须支持 Headless 模式**
- **Core 不能强依赖 Wails**
- **UI 可被完全禁用**
- 商业能力（云升级 / License）不得污染 Core
- 需兼顾跨平台一致性（Windows / macOS / Linux）

---

## 外部依赖
- Wails（桌面宿主框架）
- 操作系统原生能力（通过 Adapter 间接访问）

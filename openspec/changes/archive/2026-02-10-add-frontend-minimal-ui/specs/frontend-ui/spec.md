# frontend-ui Specification (Delta)

## 新增需求
### 需求：最小 UI 页面
前端必须提供两个页面：
- Health Check（默认首页）
- About / System Info

#### 场景：访问默认首页
- **当** 用户打开应用
- **那么** 默认显示 Health Check 页面

### 需求：路由与路径规范
前端必须使用 React Router，系统页面路径必须使用 /__system/* 命名：
- /__system/health
- /__system/about

#### 场景：访问系统页面
- **当** 用户访问系统页面
- **那么** 路径必须位于 /__system/* 且路由可用

### 需求：布局与交互
前端必须使用 Mantine AppShell，左侧固定 Sidebar，主内容单一焦点；Sidebar 仅包含品牌与两个导航项。

#### 场景：渲染主布局
- **当** 渲染主界面
- **那么** 必须使用 AppShell 并包含固定 Sidebar 与单一焦点内容区

### 需求：页面内容结构
每页必须包含：标题、简短说明、一个主操作按钮（Primary CTA）。
- Health Check CTA 为 Start Scan
- System Info CTA 为主操作按钮（文本由前端提供）

#### 场景：渲染页面内容
- **当** 渲染 Health Check 或 System Info 页面
- **那么** 必须包含标题、说明与一个主操作按钮

### 需求：System Info 展示
System Info 页面必须只读展示：Version、Build Time、Environment。

#### 场景：展示系统信息
- **当** 用户打开 System Info 页面
- **那么** 必须展示 Version、Build Time、Environment

### 需求：数据来源约束
System Info 数据必须通过 Wails Go API 获取，前端不得模拟数据。

#### 场景：获取系统信息
- **当** 前端需要系统信息
- **那么** 必须调用 Wails Go API 获取并展示

### 需求：风格约束
UI 风格必须简洁、可信、克制，参考 CCleaner / 1Password，禁止后台管理风格布局与表格。

#### 场景：设计界面风格
- **当** 设计页面样式与布局
- **那么** 必须符合风格约束且禁止使用表格与后台布局

# frontend-guidelines Specification

## Purpose
规范 DAF-Wails 前端开发的技术栈、代码风格、设计约束，确保前端实现的一致性、可维护性与产品品质。
## 需求
### 需求：包管理器约束
前端项目必须使用 pnpm（>= 9.x）进行包管理，禁止使用 npm 或 yarn。

#### 场景：初始化或维护前端项目
- **当** 初始化或维护前端项目时
- **那么** 必须使用 pnpm install；生成 pnpm-lock.yaml；禁止 npm/yarn

### 需求：React 版本锁定
前端必须使用 React 18.2.0（精确版本），TypeScript >= 5.x。禁止 React 19+。

#### 场景：配置 package.json 依赖
- **当** 配置 package.json 时
- **那么** 必须 React 18.2.0、react-dom 18.2.0、@types/react@18.x、@types/react-dom@18.x、TypeScript >= 5.x

### 需求：Mantine 版本约束
UI 组件库必须使用 Mantine v7.x（@mantine/core >= 7.0.0, @mantine/hooks >= 7.0.0）。禁止混用 Tailwind CSS、Ant Design、Material UI、Chakra UI、shadcn/ui。

#### 场景：安装 UI 组件库
- **当** 选择 UI 组件库时
- **那么** 必须使用 Mantine v7；禁止混用其他 CSS 框架或组件库

### 需求：前端技术栈约束
前端必须使用以下技术栈与组件体系：
- React + TypeScript
- Vite（Wails 推荐）
- Mantine v7.x（唯一 UI 组件库）
- @tabler/icons-react（图标库）

并禁止使用：
- Ant Design / Material UI / Bootstrap / Chakra UI / shadcn/ui
- styled-components（作为独立使用）
- 混用其他 CSS-in-JS 体系

#### 场景：引入前端依赖
- **当** 引入新的前端依赖或组件库
- **那么** 必须满足上述技术栈约束，且不得引入禁止项

### 需求：构建工具与 TypeScript 配置
前端必须使用 Vite 作为构建工具，并配置 TypeScript 编译选项以支持 Wails 类型生成。

#### 场景：配置构建环境
- **当** 配置前端构建工具时
- **那么** 必须使用 Vite；vite.config.ts 配置 wailsjs 路径别名；tsconfig.app.json 包含 wailsjs includes 和 paths 映射；build 脚本包含 `tsc -b && vite build`

### 需求：函数组件与 Hooks 约束
前端必须使用函数组件 + hooks，禁止 class components、HOC、render props。

#### 场景：编写 React 组件
- **当** 编写 React 组件时
- **那么** 必须使用函数组件、useState/useEffect 等 hooks；禁止 class component、HOC、render props

### 需求：文件命名与代码组织
前端文件必须遵循 PascalCase（组件）与 camelCase（工具）命名约定，单一文件职责不超过 300 行。

#### 场景：创建新的组件或工具函数
- **当** 创建新的前端文件时
- **那么** 必须：React 组件用 PascalCase + .tsx（HealthCheck.tsx）；工具函数用 camelCase + .ts（systemInfo.ts）；每个文件一个主体导出；避免混合多个概念；文件 > 300 行必须拆分

### 需求：禁止项（CSS 框架、组件库、模式）
禁止以下技术与模式：Tailwind CSS、Ant Design、Material UI、Bootstrap、Chakra UI、shadcn/ui、Admin/Dashboard 风格、class component、Redux/Zustand（小型应用）、styled-components（独立使用）。

#### 场景：选择技术栈或架构
- **当** 引入新的库、框架或设计模式时
- **那么** 禁止上述技术与模式，确保一致性和桌面应用定位

### 需求：Mantine AppShell 布局约束
前端应用必须使用 Mantine AppShell 作为顶层容器，实现标准的左侧导航栏布局。禁止手写 div + className 布局。

#### 场景：设计应用整体布局
- **当** 设计前端应用的顶层布局时
- **那么** 必须使用 Mantine <AppShell>；配置 <AppShell.Navbar>（200-250px 宽）和 <AppShell.Main>；禁止手写 CSS flexbox/grid 布局；桌面单窗口模式（不考虑响应式）

### 需求：设计风格参考（工具型、克制风格）
前端 UI 设计必须遵循克制、工具型的设计风格，参考 CCleaner、1Password、Dropbox Desktop 等应用。

#### 场景：设计页面布局与交互
- **当** 设计前端页面的 UI 与交互时
- **那么** 必须大留白（页面宽度 <= 800px）；工具型、专注功能；清晰层级；单一主操作；状态指示用 Alert；避免表格（< 10 行除外）；避免 3+ 层嵌套
前端必须仅作为展示层，所有系统数据必须来自 Wails Backend API：
- 前端不得模拟系统状态
- 前端不得包含业务逻辑判断
- 仅展示“事实”，不推断“结论”

#### 场景：展示系统状态
- **当** 前端展示系统状态
- **那么** 必须来自 backend API 且不在前端推断业务结论

### 需求：前端目录结构约束
前端目录必须遵循以下结构：
```
frontend/
├─ public/
├─ src/
│  ├─ app/
│  ├─ pages/
│  ├─ components/
│  ├─ hooks/
│  ├─ services/
│  ├─ styles/
│  ├─ types/
│  └─ main.tsx
```

#### 场景：新增前端文件
- **当** 新增前端文件或目录
- **那么** 必须放置在规定目录结构中

### 需求：UI 设计约束（Mantine）
前端 UI 设计必须符合以下要求：
- 欧美桌面工具风格
- 低信息密度
- 大留白
- 明确层级
- 单一主操作

并必须使用以下组件：
- AppShell
- Stack / Group / Box

禁止：
- div + className 布局
- 手写颜色值
- 覆盖组件内部样式

#### 场景：构建界面布局
- **当** 构建界面布局
- **那么** 必须使用 Mantine 组件与布局约束，且不得使用禁止项

### 需求：状态与数据约束
- 所有数据必须来自 services/*
- services 必须仅调用 Wails API
- 禁止缓存关键系统状态

#### 场景：新增数据获取逻辑
- **当** 新增数据获取逻辑
- **那么** 必须通过 services 调用 Wails API 且不得缓存关键系统状态

### 需求：错误与异常处理
- 必须使用 Mantine Alert
- 错误文本必须来自 backend
- 禁止吞错或美化错误

#### 场景：前端展示错误信息
- **当** 前端展示错误信息
- **那么** 必须使用 Mantine Alert 且错误文本来自 backend

### 需求：AI 生成代码原则
必须：
- 遵循目录结构
- 最小可用实现
- 不假设不存在的 API

禁止：
- 过度封装
- 提前引入状态管理
- 装饰性 UI

#### 场景：AI 生成或修改前端代码
- **当** 生成或修改前端代码
- **那么** 必须遵循上述原则与禁止项

### 需求：最高优先级规则
本项目必须以桌面系统基础设施产品定位为最高优先级，不是营销网站或后台管理系统。

#### 场景：设计前端界面
- **当** 设计前端界面与交互
- **那么** 必须以桌面系统基础设施产品定位为最高优先级

### 需求：桌面窗口初始化尺寸
桌面应用必须使用固定的初始化窗口尺寸，并设置最小尺寸，且不得强制启动最大化。

#### 场景：应用首次启动
- **当** 应用首次启动
- **那么** 必须设置默认窗口尺寸为 1200×800
- **并且** 必须设置最小窗口尺寸为 1024×720
- **并且** 不得强制最大化启动

---

### 需求：Sidebar 尺寸与菜单元素
Sidebar 必须采用固定宽度与尺寸范围，确保左侧导航的一致性。

#### 场景：渲染侧边栏导航
- **当** 渲染侧边栏
- **那么** Sidebar 宽度必须为 240px
- **并且** Logo 区高度必须为 64px
- **并且** 菜单项高度必须在 40–44px 范围内
- **并且** 菜单图标尺寸必须在 18–20px 范围内
- **并且** 菜单文本字号必须为 14px

---

### 需求：主内容区与顶部标题栏
主内容区与顶部标题栏必须遵循固定比例与高度范围。

#### 场景：布局主内容区
- **当** 布局主内容区时
- **那么** 主内容区宽度必须为 `窗口宽度 - 240px`
- **并且** 顶部标题栏高度必须在 56–64px 范围内

---

### 需求：缩放与中文排版约束
在 125% 缩放与中文文本环境下，界面必须保持可用性与排版一致性。

#### 场景：系统缩放为 125%
- **当** 系统缩放为 125%
- **那么** UI 不得溢出或遮挡关键内容

#### 场景：中文文本展示
- **当** 页面使用中文文本
- **那么** 菜单与标题文本不得换行挤压布局


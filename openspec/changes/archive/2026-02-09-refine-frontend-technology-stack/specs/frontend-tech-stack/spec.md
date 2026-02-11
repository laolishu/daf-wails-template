# frontend-tech-stack Specification（规范增量）

本规范对现有 `frontend-guidelines` 进行**精化**，明确包管理、版本约束、代码风格等细节。

## 修改需求

### 需求：包管理器约束
**前端项目必须使用 pnpm（>= 9.x）进行包管理，禁止使用 npm 或 yarn。**

#### 场景：初始化前端项目
- **当** 初始化或维护前端项目
- **那么** 必须：
  - 使用 `pnpm install` 而非 `npm install` / `yarn install`
  - 生成 pnpm-lock.yaml（不提交 package-lock.json 或 yarn.lock）
  - 所有脚本命令使用 `pnpm` 前缀（`pnpm run build` 等）
  - 确保 pnpm 版本 >= 9.0.0

#### 反面例子
```
✗ npm install                    # 禁止
✗ yarn install                   # 禁止
✓ pnpm install                   # 正确
```

---

### 需求：React 版本锁定
**前端必须使用 React 18.x，并锁定 TypeScript 类型定义版本。禁止 React 19 及以上。**

#### 场景：配置 package.json 依赖
- **当** 配置前端项目的 package.json
- **那么** 必须：
  - React 版本：`18.2.0`（精确锁定）
  - react-dom 版本：`18.2.0`（精确锁定）
  - @types/react 版本：`18.x`（主版本锁定）
  - @types/react-dom 版本：`18.x`（主版本锁定）
  - TypeScript 版本：`>= 5.0.0`

#### 反面例子
```json
{
  "dependencies": {
    "react": "^19.2.0",          // ✗ 禁止 React 19
    "react-dom": "19.x"           // ✗ 禁止 React 19
  },
  "devDependencies": {
    "@types/react": "^19.2.5",   // ✗ 禁止 React 19 类型定义
    "typescript": "^4.9.0"        // ✗ TypeScript 过旧
  }
}
```

#### 正确例子
```json
{
  "dependencies": {
    "react": "18.2.0",
    "react-dom": "18.2.0"
  },
  "devDependencies": {
    "@types/react": "18.2.0",
    "@types/react-dom": "18.2.0",
    "typescript": "~5.9.0"
  }
}
```

---

### 需求：Mantine 版本约束
**UI 组件库必须使用 Mantine v7.x（及以上 v7 维护版本）。禁止混用其他 CSS 框架。**

#### 场景：安装 UI 组件库
- **当** 安装 UI 组件库时
- **那么** 必须：
  - @mantine/core >= 7.0.0
  - @mantine/hooks >= 7.0.0
  - @mantine/notifications >= 7.0.0（若需通知功能）
  - 不得混用 Tailwind CSS、Ant Design、Material UI、Chakra UI、shadcn/ui
  - 不得使用 styled-components、emotion（除非作为 Mantine 的依赖）

#### 反面例子
```json
{
  "dependencies": {
    "@mantine/core": "8.3.14",           // ✗ v8（不符合）
    "tailwindcss": "^3.0.0",             // ✗ 禁止混用 Tailwind
    "@emotion/react": "^11.0.0"          // ✗ 不独立使用 emotion
  }
}
```

#### 正确例子
```json
{
  "dependencies": {
    "@mantine/core": "^7.3.0",
    "@mantine/hooks": "^7.3.0",
    "@mantine/notifications": "^7.3.0"
  }
}
```

---

### 需求：构建工具与 TypeScript 配置
**前端必须使用 Vite 作为构建工具，并配置 TypeScript 编译选项以支持 Wails 类型生成。**

#### 场景：配置构建环境
- **当** 配置前端构建工具
- **那么** 必须：
  - 使用 Vite（推荐 rolldown-vite）
  - vite.config.ts 配置路径别名以支持 `wailsjs` 导入
  - tsconfig.app.json 包含 `wailsjs` 路径、baseUrl、paths 映射
  - `build` 脚本包含 `tsc -b` 预编译步骤
  - 编译输出到 `dist/` 目录

#### 反面例子
```typescript
// ✗ 未配置 wailsjs 路径
import { GetSystemInfo } from "../../wailsjs/go/backend/App"

// ✗ 未在 tsconfig 中声明
// 导致 TypeScript 无法识别 wailsjs 类型
```

#### 正确例子
```typescript
// vite.config.ts
import path from "path"
export default {
  resolve: {
    alias: {
      wailsjs: path.resolve(__dirname, 'wailsjs')
    }
  }
}

// tsconfig.app.json
{
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "wailsjs/*": ["./wailsjs/*"]
    }
  },
  "include": ["src", "wailsjs"]
}

// package.json
{
  "scripts": {
    "build": "tsc -b && vite build"
  }
}
```

---

### 需求：函数组件与 Hooks 约束
**前端必须使用函数组件 + hooks，禁止使用 class components。所有组件必须从 React hooks 中获取状态与生命周期。**

#### 场景：编写 React 组件
- **当** 编写 React 组件
- **那么** 必须：
  - 使用函数组件语法（const Component = () => ...）
  - 使用 useState、useEffect、useContext 等 hooks
  - 禁止 class 组件（extends React.Component）
  - 禁止 HOC（High Order Component）—— 使用 custom hooks 替代
  - 禁止 render props

#### 反面例子
```typescript
// ✗ class component
class HealthCheck extends React.Component {
  render() {
    return <h2>Health Check</h2>
  }
}

// ✗ 没有使用 hooks 管理状态
function SystemInfo() {
  let info = null  // 禁止
  return <div>{info}</div>
}
```

#### 正确例子
```typescript
// ✓ 函数组件 + useState
function HealthCheck() {
  const [scanning, setScanning] = useState(false)
  return <h2>Health Check</h2>
}

// ✓ 使用 useEffect 获取数据
function SystemInfo() {
  const [info, setInfo] = useState(null)
  useEffect(() => {
    fetchSystemInfo().then(setInfo)
  }, [])
  return <div>{info?.version}</div>
}
```

---

### 需求：文件命名与代码组织
**前端文件必须遵循 PascalCase（组件）与 camelCase（工具）命名约定，单一文件职责不超过 300 行。**

#### 场景：创建新的组件或工具函数
- **当** 创建新的前端文件
- **那么** 必须：
  - React 组件文件：PascalCase + .tsx 扩展（如 HealthCheck.tsx、SystemInfo.tsx）
  - 工具函数：camelCase + .ts 扩展（如 systemInfo.ts、parseError.ts）
  - 每个文件仅导出一个主体（组件或函数）
  - 避免混合多个概念（不将 A 组件和 B 工具函数放在同一文件）
  - 文件行数 > 300 行时必须拆分

#### 反面例子
```typescript
// ✗ 小写命名
export function healthcheck() { }      // 应为 HealthCheck

// ✗ 文件过长（混合多个概念）
// HealthCheckAndSystemInfo.tsx (600+ lines)
function HealthCheck() { }
function SystemInfo() { }
export { HealthCheck, SystemInfo }

// ✗ 工具函数用 PascalCase
export function GetSystemInfo() { }    // 应为 getSystemInfo
```

#### 正确例子
```typescript
// ✓ HealthCheck.tsx（仅包含 HealthCheck 组件，< 50 行）
export default function HealthCheck() { ... }

// ✓ SystemInfo.tsx（仅包含 SystemInfo 组件，< 100 行）
export default function SystemInfo() { ... }

// ✓ systemInfo.ts（工具函数，< 100 行）
export async function fetchSystemInfo(): Promise<SystemInfo> { }

// ✓ parseError.ts（工具函数，< 50 行）
export function parseErrorMessage(err: Error): string { }
```

---

### 需求：禁止项（CSS 框架、组件库、模式）
**前端禁止以下技术与模式，以保证工具型设计风格和代码一致性。**

#### 场景：选择技术栈或架构
- **当** 引入新的库、框架或设计模式时
- **那么** 禁止：
  - Tailwind CSS（CSS 工具框架）
  - Ant Design、Material UI、Bootstrap（UI 组件库）
  - Chakra UI、shadcn/ui（UI 组件库）
  - Admin 或 Dashboard 设计风格（复杂表格、深层嵌套、高信息密度）
  - class component（使用函数组件替代）
  - Redux / Zustand（小型应用无需复杂状态管理）
  - styled-components（使用 Mantine 主题 + CSS modules 替代）

#### 反面例子
```typescript
// ✗ 使用 Tailwind
<div className="flex items-center justify-between bg-gray-100">

// ✗ Dashboard 风格
<Table rows={1000} columns={20} pagination scrollable />

// ✗ class component
class MyComponent extends React.Component { }

// ✗ Redux
import { useDispatch } from 'react-redux'
```

#### 正确例子
```typescript
// ✓ Mantine 组件
<Group justify="space-between">

// ✓ 工具型、简洁的信息展示
<Stack>
  <Text>Version: 1.0.0</Text>
  <Text>Build Time: 2024-01-01</Text>
</Stack>

// ✓ 函数组件
function MyComponent() { }

// ✓ React Context + hooks（无需 Redux）
const [state, setState] = useState(...)
```

---

### 需求：Mantine AppShell 布局约束
**前端应用必须使用 Mantine AppShell 作为顶层容器，实现标准的左侧导航栏布局。禁止手写 div + className 布局。**

#### 场景：设计应用整体布局
- **当** 设计前端应用的顶层布局
- **那么** 必须：
  - 使用 Mantine `<AppShell>` 组件（而非 div + flexbox）
  - 配置 `<AppShell.Navbar>` 作为左侧导航（宽度 200-250px）
  - 配置 `<AppShell.Main>` 作为主内容区
  - 禁止手写 CSS flexbox 或 grid 实现布局
  - 布局必须不考虑响应式断点（桌面单窗口模式）

#### 反面例子
```typescript
// ✗ 手写 div 布局
<div style={{ display: 'flex', height: '100vh' }}>
  <div style={{ width: '240px', borderRight: '1px solid...' }}>
    导航栏
  </div>
  <div style={{ flex: 1 }}>
    主内容
  </div>
</div>

// ✗ 响应式考虑（不符合桌面单窗口约束）
<AppShell navbar={{ width: 240, breakpoint: 'sm' }} />
```

#### 正确例子
```typescript
// ✓ 使用 AppShell
<AppShell navbar={{ width: 240 }} withBorder>
  <AppShell.Navbar>
    <NavBar />
  </AppShell.Navbar>
  <AppShell.Main>
    <MainContent />
  </AppShell.Main>
</AppShell>
```

---

### 需求：设计风格参考（工具型、克制风格）
**前端 UI 设计必须遵循克制、工具型的设计风格，参考 CCleaner、1Password、Dropbox Desktop 等应用。禁止过度设计与装饰。**

#### 场景：设计页面布局与交互
- **当** 设计前端页面的 UI 与交互
- **那么** 必须：
  - **大留白**：避免信息密度过高，页面宽度不超过 800px
  - **工具型**：专注功能，避免装饰性 UI（无不必要的图标、渐变、动画）
  - **清晰层级**：使用 Mantine 的 Title / Text / Button 建立视觉层级
  - **单一主操作**：每个页面一个主要操作（如 "Start Scan"）
  - **状态指示**：用 Alert 或 Text 清晰展示错误、加载、成功状态
  - **禁止表格**：避免使用 Table 组件（除非数据 < 10 行）
  - **禁止嵌套**：避免 3 层以上的组件嵌套

参考应用特征：
- **CCleaner 风格**：简洁的主按钮、清晰的扫描结果、深浅配色对比
- **1Password 风格**：大留白、侧边栏导航、只读展示（无复杂表单）
- **Dropbox Desktop 风格**：状态指示为主、避免复杂弹窗、信息浓缩显示

#### 反面例子
```typescript
// ✗ 过度设计
<div style={{ background: 'linear-gradient(...)', boxShadow: '...' }}>
  <div style={{ animation: 'fadeIn 0.5s' }}>
    装饰性图标和渐变背景
  </div>
</div>

// ✗ 信息密度过高（表格）
<Table rows={100} columns={10} paginated scrollable />

// ✗ 嵌套过深
<Box><Group><Stack><Box><Group>...

// ✗ Dashboard 风格
<div>
  <TopBar><Charts /></TopBar>
  <LeftPanel><Menu /></LeftPanel>
  <MainPanel><ComplexForm /></MainPanel>
</div>
```

#### 正确例子
```typescript
// ✓ 克制风格
<Stack gap="lg" p="md" maw={720} mx="auto">
  <Title order={2}>Health Check</Title>
  <Text c="dimmed">Run a quick scan to verify services.</Text>
  <Button size="lg">Start Scan</Button>
</Stack>

// ✓ 工具型信息展示
<Stack gap="md">
  <Group justify="space-between">
    <Text>Version</Text>
    <Text fw={500}>1.0.0</Text>
  </Group>
  <Group justify="space-between">
    <Text>Status</Text>
    <Badge>Healthy</Badge>
  </Group>
</Stack>

// ✓ 状态指示
{error && <Alert color="red" title="Error">{error}</Alert>}
{loading && <Loader />}
{data && <DataDisplay data={data} />}
```

---

## 总结

本规范增量涵盖 **8 个新增需求**，精化 `frontend-guidelines`：
1. 包管理器：pnpm >= 9.x
2. React 版本：18.2.0（精确锁定）
3. Mantine 版本：v7.x
4. 构建工具：Vite + TypeScript 配置
5. 函数组件 + hooks
6. 文件命名与代码组织
7. 禁止项列表
8. Mantine AppShell 布局约束
9. 设计风格参考（工具型、克制）

**实施后预期**：
- 所有前端项目使用统一的 pnpm workspace
- React 18 + Mantine v7 的稳定技术栈
- 清晰的代码组织与设计风格
- 可通过自动化规范检查确保一致性

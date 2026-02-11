# Design Doc: 前端技术栈细化

## 架构决策记录

### 为什么选择 pnpm？
- **一致性**：pnpm workspace 确保跨项目依赖一致，避免幽灵依赖
- **性能**：比 npm/yarn 更快，特别是在 node_modules 安装时
- **预留**：Wails + Go + Node.js 多工程场景下，workspace 管理更清晰
- **防守**：明确禁止 npm/yarn 防止开发者无意中破坏依赖

### 为什么选择 React 18.x 而非 19？
- **Mantine 兼容性**：Mantine v7 对标 React 18，v8 虽支持 18 但会预期 18.2+ 的稳定特性
- **稳定性**：React 18 在 2024 年底已充分验证，18.2.0 是当前 LTS 版本
- **成熟生态**：大多数中间件、工具库都基于 React 18 测试
- **防守**：React 19 引入实验性 hooks（useId、use），易导致版本混淆

### 为什么选择 Mantine v7？
- **官方推荐**：Mantine v7 是当前稳定推荐版本
- **文档完整**：v7 文档成熟、社区案例多
- **AppShell 稳定**：Mantine v7 的 AppShell 组件经过充分验证
- **与 React 18 配套**：v7 与 React 18.2 配套最佳

### 为什么禁止 Tailwind + 明确禁止 Admin/Dashboard 风格？
- **一致性**：混用 Tailwind + Mantine 会导致 CSS 冲突
- **品牌**：CCleaner / 1Password 都是克制的工具风格，避免管理系统的复杂性
- **约束**：明确风格参考，防止过度设计

## 代码风格约束的理由

### 函数组件 + hooks（禁止 class components）
- React 18 官方已弃用 class components
- 函数组件 + hooks 的性能和可读性更优
- 防止向后维护遗产代码

### PascalCase 组件、camelCase 工具
- 遵循 TypeScript 社区约定
- 区分组件（导出为 React 元素）与工具函数
- IDE 自动识别，便于搜索

### 单一职责 + 避免超长文件
- 组件最多 300 行，超过则拆分
- 每个文件一个导出，禁止混合多个概念

## 设计参考的具体含义

### CCleaner 风格
- 清晰的主操作（Start Scan、Optimize 等）
- 简洁的信息展示（过去 24h、昨日等）
- 低复杂度的交互

### 1Password 风格
- 大留白、明确优先级
- 侧边栏导航（Mantine AppShell）
- 读写分离（只读展示 vs 编辑态）

### Dropbox Desktop 风格
- 极简的主界面
- 状态指示（同步状态、错误提示）
- 尽量避免表格

## 与当前 add-frontend-minimal-ui 的矛盾
当前实现已改用：
- npm（而非 pnpm）
- React 19.2.0（而非 18.x）
- 原生 HTML/CSS（而非 Mantine AppShell）

**需要改正为**：
- 安装 pnpm，迁移 package.json
- 降级 React 到 18.2.0，确保 Mantine v7 兼容
- 改用 Mantine AppShell + 标准组件

这个改正会在 apply 阶段实施。

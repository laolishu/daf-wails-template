# Tasks: 细化前端技术栈规范

## 提案阶段
- [x] 编写 proposal.md
- [x] 编写 design.md
- [x] 编写规范增量（spec.md）
- [x] 验证提案（openspec-cn validate）

## 规范细化阶段
- [x] **Task 1**：在 `frontend-guidelines/spec.md` 中添加需求：包管理器约束（pnpm >= 9.x）
  - 输出：新增需求文本 + 场景描述 ✓
  - 验收：需求包含清晰的约束和违反示例 ✓

- [x] **Task 2**：在 `frontend-guidelines/spec.md` 中添加需求：React & TypeScript 版本锁定
  - 版本：React 18.2.0、@types/react@18.x、TypeScript >= 5.x ✓
  - 输出：需求 + 禁止项（不得使用 React 19+） ✓
  - 验收：版本号明确，包含 package.json 示例 ✓

- [x] **Task 3**：在 `frontend-guidelines/spec.md` 中添加需求：Mantine v7 版本约束
  - 版本：@mantine/core >= 7.0.0、@mantine/hooks >= 7.0.0、@mantine/notifications >= 7.0.0 ✓
  - 输出：需求 + 禁止其他 CSS 框架 ✓
  - 验收：版本号与兼容性清晰 ✓

- [x] **Task 4**：在 `frontend-guidelines/spec.md` 中添加需求：代码风格约束（函数组件、命名约定、文件职责）
  - 约束：函数组件 + hooks、PascalCase 组件、camelCase 工具、单一职责、避免超长文件 ✓
  - 输出：需求 + 反面例子 ✓
  - 验收：包含代码示例 ✓

- [x] **Task 5**：在 `frontend-guidelines/spec.md` 中添加需求：禁止项补充（Tailwind、Admin 风格、class components）
  - 输出：明确禁止列表 ✓
  - 验收：禁止项与原因清晰 ✓

- [x] **Task 6**：在 `frontend-guidelines/spec.md` 中添加需求：设计风格参考（CCleaner / 1Password / Dropbox）
  - 输出：设计参考 + 具体含义（大留白、工具型、克制） ✓
  - 验收：可指导 UI 设计决策 ✓

- [x] **Task 7**：验证规范增量
  - 运行：`openspec-cn validate refine-frontend-technology-stack --strict` ✓
  - 验收：无错误、所有需求合法 ✓

## 实施阶段（apply）
- [x] **Task 8**：迁移前端项目到 pnpm
  - 删除 node_modules、package-lock.json ✓
  - 运行 `pnpm install` ✓
  - 验收：pnpm-lock.yaml 生成、依赖正确 ✓

- [x] **Task 9**：恢复 React 18 + Mantine v7（改回当前的原生 HTML 实现）
  - 修改 package.json：React 18.2.0、Mantine v7.17.8 ✓
  - 改回 App.tsx / pages / CSS 使用 Mantine AppShell ✓
  - 验收：pnpm run build 成功、无编译错误 ✓

- [x] **Task 10**：启动 wails dev 验证
  - 启动 `wails dev`
  - 预期：Health Check 页面显示、System Info 可切换、数据通过 Wails API 获取
  - 用户测试：在浏览器查看应用是否正常显示
  - ✓ 应用成功启动于 http://localhost:34115
  - ✓ 侧边栏显示 DAF-Wails + 菜单项
  - ✓ 页面切换正常工作

- [x] **Task 11**：存档完成的变更
  - 运行：`openspec-cn archive refine-frontend-technology-stack --yes`
  - 验收：变更移到 archive 目录 ✓
  - 归档时间戳：2026-02-09-refine-frontend-technology-stack

   - HealthCheck.tsx：Mantine Stack/Title/Text/Button
   ## 规范细化阶段
   - [x] **Task 1**：在 `frontend-guidelines/spec.md` 中添加需求：包管理器约束（pnpm >= 9.x）
   - index.css：简化为 Mantine 主题配置
   - [x] **Task 2**：在 `frontend-guidelines/spec.md` 中添加需求：React & TypeScript 版本锁定
5. ✅ 前端编译成功
   - [x] **Task 3**：在 `frontend-guidelines/spec.md` 中添加需求：Mantine v7 版本约束
   - dist/assets/index-B2CG-grM.js 266.13 kB（未压缩）
   - [x] **Task 4**：在 `frontend-guidelines/spec.md` 中添加需求：代码风格约束（函数组件、命名约定、文件职责）

   - [x] **Task 5**：在 `frontend-guidelines/spec.md` 中添加需求：禁止项补充（Tailwind、Admin 风格、class components）
- 用户应在浏览器（http://localhost:*）查看应用
   - [x] **Task 6**：在 `frontend-guidelines/spec.md` 中添加需求：设计风格参考（CCleaner / 1Password / Dropbox）
  1. Sidebar 显示（DAF-Wails 标题和两个导航链接）
   - [x] **Task 7**：验证规范增量
  3. 点击 "About / System Info" 可切换页面
  4. 系统信息（Version、Build Time、Environment）正确显示
   - [x] **Task 8**：迁移前端项目到 pnpm

   - [x] **Task 9**：恢复 React 18 + Mantine v7（改回当前的原生 HTML 实现）
- 用户验证应用正常运行
   - [ ] **Task 10**：测试前端功能

   - [ ] **Task 11**：存档完成的变更
- 运行 `openspec-cn archive refine-frontend-technology-stack --yes` 存档变更

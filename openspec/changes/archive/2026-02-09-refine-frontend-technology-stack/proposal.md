# 提案：细化前端技术栈与包管理规范

## 摘要
当前 `frontend-guidelines` 规范对技术栈的约束不够具体。本提案**精化**现有规范，明确：
1. **包管理器**：必须使用 pnpm（>= 9.x）而非 npm/yarn
2. **React 版本**：锁定 React 18.x（而非 19+）以保证 Mantine 兼容性
3. **Mantine 版本**：v7.x（官方推荐与稳定版本）
4. **构建工具细节**：Vite + TypeScript 预编译配置
5. **代码风格**：函数组件 + hooks、PascalCase 组件文件名
6. **设计参考**：CCleaner / 1Password / Dropbox Desktop 风格

## 为什么
- 当前规范提到 "Mantine" 但未指定版本，导致可能出现 React 19 + Mantine 8 不兼容的情况
- 未明确包管理工具，无法保证跨开发者环境一致性
- 缺少代码风格与文件命名约定细节，影响代码审查效率
- 缺少设计风格的具体参考，设计决策缺乏方向
- 不同版本组合的构建配置差异导致打包问题（如 useId hook 兼容性）

## 问题陈述
- 当前规范提到 "Mantine" 但未指定版本，导致可能出现 React 19 + Mantine 8 不兼容的情况
- 未明确包管理工具，无法保证跨开发者环境一致性
- 缺少代码风格与文件命名约定细节
- 缺少设计风格的具体参考

## 变更内容
本提案通过更新 `frontend-guidelines` 规范来解决上述问题，涉及以下需求变更：
1. **新增需求**：包管理器约束（pnpm >= 9.x）
2. **新增需求**：React & 类型定义版本锁定（React 18.x、TypeScript >= 5.x）
3. **新增需求**：Mantine 版本约束（v7.x）
4. **新增需求**：代码风格细则（函数组件、命名约定、单一职责）
5. **新增需求**：禁止项补充（Tailwind CSS、Admin 风格、class components）
6. **新增需求**：设计风格参考（CCleaner / 1Password / Dropbox）

## 解决方案
更新 `frontend-guidelines` 规范，新增以下需求：
1. **包管理器约束**：pnpm >= 9.x，禁止 npm/yarn
2. **React & 类型定义版本锁定**：React 18.x、TypeScript >= 5.x
3. **Mantine 版本**：v7.x（@mantine/core, @mantine/hooks, @mantine/notifications）
4. **代码风格细则**：函数组件、hooks、PascalCase/camelCase 命名、单一职责
5. **禁止列表补充**：Tailwind CSS、Admin/Dashboard 风格、class components
6. **设计风格参考**：美国/欧洲用户、克制、清晰、工具型

## 影响范围
- **修改对象**：`openspec/specs/frontend-guidelines/spec.md`（精化/扩展现有规范）
- **修改类型**：REFINED（精化现有需求以增加详细度）
- **约束范围**：所有前端代码（组件、页面、服务）
- **向后兼容性**：与当前 add-frontend-minimal-ui 实现有矛盾（需要回滚到 React 18 + Mantine v7）

## 验收标准
1. 规范文档包含所有 10 个新增需求
2. 规范可通过 `openspec-cn validate` 验证
3. 现有 add-frontend-minimal-ui 实现需要根据新规范调整（改用 pnpm + React 18 + Mantine v7）

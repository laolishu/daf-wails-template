# Frontend AI 行为约束（Wails + Mantine）

> 本约束用于规范 AI（如 GitHub Copilot / Copilot Chat）在本项目中生成、修改前端代码时的行为，确保与 **Wails 架构、Mantine 设计哲学、桌面应用场景** 保持一致。

---

## 一、技术栈硬性约束（不可违反）

### 1. 框架与工具

- React + TypeScript
- Vite
- Mantine（唯一 UI 组件库）
- @tabler/icons-react

禁止：
- Ant Design / Material UI / Bootstrap / Chakra UI
- styled-components
- 混用其他 CSS-in-JS 体系

---

### 2. 与 Wails 的关系约束

- 前端仅作为展示层（View Layer）
- 所有系统数据来自 Wails Backend API
- 前端不得模拟系统状态
- 前端不得包含业务逻辑判断

原则：前端只展示“事实”，不推断“结论”。

---

## 二、目录与文件约束

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

---

## 三、UI 设计约束（Mantine）

- 欧美桌面工具风格
- 低信息密度
- 大留白
- 明确层级
- 单一主操作

必须使用：
- AppShell
- Stack / Group / Box

禁止：
- div + className 布局
- 手写颜色值
- 覆盖组件内部样式

---

## 四、状态与数据约束

- 所有数据来自 services/*
- services 仅调用 Wails API
- 不缓存关键系统状态

---

## 五、错误与异常

- 使用 Mantine Alert
- 错误文本来自 backend
- 不吞错、不美化错误

---

## 六、AI 生成代码原则

必须：
- 遵循目录结构
- 最小可用实现
- 不假设不存在的 API

禁止：
- 过度封装
- 提前引入状态管理
- 装饰性 UI

---

## 七、最高优先级规则

这是一个桌面系统基础设施产品，不是营销网站，也不是后台管理系统。

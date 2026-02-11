# core-system-info Specification (Delta)

## 新增需求
### 需求：系统信息能力
core 必须提供系统信息读取能力，包含以下只读字段：
- Version
- BuildTime
- Environment

#### 场景：读取系统信息
- **当** 需要展示系统信息
- **那么** core 必须返回包含 Version/BuildTime/Environment 的结构化数据

### 需求：Wails API 暴露
backend 必须将 core 系统信息能力通过 Wails Go API 暴露给前端。

#### 场景：前端调用系统信息 API
- **当** 前端请求系统信息
- **那么** 必须通过 Wails API 调用 backend 暴露的方法并获得数据

# frontend-ui 规范增量：Examples 页面

## 新增需求

### 需求：Examples 页面入口
前端必须新增 Examples 页面，并作为主导航入口，名称支持 i18n，默认显示英文 “Examples”。

#### 场景：渲染主导航
- **当** 渲染主导航
- **那么** 必须包含 Examples 导航项
- **并且** 导航文案支持国际化，默认英文为 “Examples”

---

### 需求：Examples 页面结构
Examples 页面必须以卡片形式展示真实系统调用示例。

#### 场景：用户进入 Examples 页面
- **当** 用户进入 Examples 页面
- **那么** 必须展示 Logger / System Info / Config Read 三个示例卡片

---

### 需求：Logger Example
Logger Example 必须触发一次真实日志写入，并返回结果状态。

#### 场景：写入测试日志
- **当** 用户点击 “Write Test Log”
- **那么** 前端必须通过 Wails API 调用 Go 写入 INFO 日志
- **并且** 日志内容包含固定前缀 `[HealthCheck]` 与时间戳
- **并且** 成功时提示“到日志路径下查看”
- **并且** 失败时展示错误原因（只读）

---

### 需求：System Info Example
System Info Example 必须从 Go 侧读取系统信息并只读展示。

#### 场景：加载系统信息
- **当** 用户点击 “Load System Info”
- **那么** 必须展示版本号、构建时间、运行平台

---

### 需求：Config Read Example
Config Read Example 必须读取用户配置并展示关键字段。

#### 场景：读取配置
- **当** 用户点击 “Load Config”
- **那么** 必须展示语言与日志级别

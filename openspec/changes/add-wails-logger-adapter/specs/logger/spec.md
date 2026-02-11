# logger 规范增量：Wails 日志适配器

## 修改需求

### 需求：Wails 日志适配
系统必须提供 Wails 日志适配器，将 Wails 运行期日志路由到 core/logger。

#### 场景：初始化 Wails 应用
- **当** 应用创建 `options.App`
- **那么** 必须设置 `options.App.Logger` 为 Wails 日志适配器
- **并且** 适配器必须将日志写入 core/logger（slog）

---

### 需求：日志级别映射
Wails 日志级别必须映射到 slog 级别，保证语义一致。

#### 场景：Wails 日志输出
- **当** Wails 输出 TRACE/DEBUG/INFO/WARNING/ERROR/FATAL
- **那么** 必须映射到 slog.Debug/Debug/Info/Warn/Error/Error
- **并且** FATAL 必须触发进程退出

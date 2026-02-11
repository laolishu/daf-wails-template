# appconfig Specification (Delta)

## 修改需求
### 需求：标准应用配置项
系统必须定义以下标准应用配置项：
- log.dir：日志目录（相对或绝对路径），默认值为 `logs`
- log.level：日志级别（debug/info/warn/error），默认值为 `info`
- log.retention_days：日志保留天数，默认值为 `7`
- i18n.language：语言代码（zh-CN/en-US/ja-JP 等），默认值为 `zh-CN`
- window.width：窗口宽度（像素），默认值为 `1024`
- window.height：窗口高度（像素），默认值为 `768`

#### 场景：获取窗口配置
- **当** Wails 应用初始化时
- **那么** 必须从 core/config 读取 window.width 和 window.height

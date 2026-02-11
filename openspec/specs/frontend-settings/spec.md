# frontend-settings Specification

## Purpose
TBD - created by archiving change update-settings-language-systeminfo. Update Purpose after archive.
## 需求
### 需求：Settings 集成 System Info
Settings 页面必须展示 System Info 的只读信息，且 System Info 不再作为独立主菜单页面出现。

#### 场景：用户打开 Settings 页面
- **当** 用户进入 Settings
- **那么** 必须看到 System Info 信息分组，且侧边栏不再提供独立的 System Info 入口

### 需求：语言选项（首项）
Settings 页面必须提供语言选项并置于第一项，仅支持中文（zh-CN）与英文（en-US）切换。

#### 场景：用户查看 Settings 语言选项
- **当** Settings 页面渲染
- **那么** 语言选项必须位于第一项且仅包含 zh-CN/en-US

### 需求：语言配置读写
语言选项必须读取并写入后端用户配置项 `i18n.language`。

#### 场景：读取当前语言
- **当** Settings 页面加载
- **那么** 必须从后端配置读取 `i18n.language` 作为当前语言值

#### 场景：写入语言配置
- **当** 用户切换语言
- **那么** 必须将新语言写入后端配置的 `i18n.language`

### 需求：语言切换即时生效
语言切换必须即时生效，当前页面文案应在切换后立即更新。

#### 场景：用户切换语言
- **当** 用户在 Settings 中切换语言
- **那么** 当前页面文案必须立即切换为对应语言


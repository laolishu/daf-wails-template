## 新增需求
### 需求：升级检查 API
系统必须在 backend 暴露升级检查 API，供前端调用并返回 UpdateInfo。

#### 场景：前端检查更新
- **当** 前端发起升级检查
- **那么** backend 必须调用 core/updater 完成检查并返回 UpdateInfo

### 需求：升级下载与安装 API
系统必须在 backend 暴露升级下载/安装 API，供前端选择下载时调用，并返回安装路径。

#### 场景：前端触发下载
- **当** 前端选择下载更新
- **那么** backend 必须调用 core/updater 执行下载、校验、解压并返回安装路径

### 需求：升级检测地址来源
升级检查必须使用 sysconfig.UpdateEndpoint 作为检测地址，禁止前端直接请求升级服务。

#### 场景：读取升级检测地址
- **当** backend 调用升级检查
- **那么** 必须从 sysconfig 获取 UpdateEndpoint 并用于 provider 配置

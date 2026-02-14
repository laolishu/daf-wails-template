## 修改需求
### 需求：标准系统变量集
系统必须定义以下标准系统变量，并通过编译时 ldflags 注入：
- Version：应用版本号（如 v1.0.0）
- BuildTime：构建时间（ISO 8601 格式）
- GitCommit：Git 提交哈希（简称，如 abc1234）
- ConfigDir：配置文件目录（如 ~/.config/daf-app）
- ConfigFile：配置文件名称（如 config.yml），默认值为 `config.yml`
- UpdateEndpoint：升级检测服务地址（HTTP URL），默认值为空字符串

其他路径相关变量（LogDir、DataDir、TempDir）由用户通过配置文件定义，不作为系统常量。

#### 场景：获取单个系统变量
- **当** 应用需要版本号或配置目录
- **那么** 必须从 sysconfig 读取标准变量

#### 场景：获取默认配置文件路径
- **当** 应用需要加载配置文件
- **那么** 可通过 sysconfig.GetConfigDir() 和 sysconfig.GetConfigFile() 组合获取完整路径

#### 场景：获取升级检测地址
- **当** 应用需要调用升级检测服务
- **那么** 必须从 sysconfig 获取 UpdateEndpoint

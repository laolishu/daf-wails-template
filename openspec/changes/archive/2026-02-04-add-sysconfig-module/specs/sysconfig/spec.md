## 新增需求
### 需求：系统配置模块
系统必须提供一个专用模块管理编译时注入的全局系统变量，该模块应位于 core/sysconfig，与 Wails 解耦，支持通过 Go 编译时 ldflags 注入。

#### 场景：应用启动获取系统信息
- **当** 应用启动时
- **那么** 必须能获取编译时注入的系统信息（如版本号、构建时间）

### 需求：标准系统变量集
系统必须定义以下标准系统变量，并通过编译时 ldflags 注入：
- Version：应用版本号（如 v1.0.0）
- BuildTime：构建时间（ISO 8601 格式）
- GitCommit：Git 提交哈希（简称，如 abc1234）
- ConfigDir：配置文件目录（如 ~/.config/daf-app）

其他路径相关变量（LogDir、DataDir、TempDir）由用户通过配置文件定义，不作为系统常量。

#### 场景：获取单个系统变量
- **当** 应用需要版本号或配置目录
- **那么** 必须从 sysconfig 读取标准变量

### 需求：系统配置与应用配置的边界
系统配置（sysconfig）必须与应用配置（core/config）清晰分离：
- sysconfig：编译时确定、运行时只读、系统级信息（版本、构建、路径）
- config：运行时加载、用户可修改、应用行为配置（日志级别、插件开关等）

#### 场景：初始化应用
- **当** 应用启动时初始化配置系统
- **那么** 必须先初始化 sysconfig（系统级）再初始化 config（应用级）

### 需求：sysconfig 访问接口
系统必须提供统一的 sysconfig 访问接口，包括：
- Info() 函数返回包含所有标准变量的结构体
- 单个 Get 函数（如 GetVersion、GetConfigDir）
- 支持零值校验（检查变量是否在编译时被正确注入）

#### 场景：获取配置目录
- **当** 应用需要查询配置目录
- **那么** 调用 sysconfig.GetConfigDir() 直接获取

### 需求：ldflags 编译注入机制
系统必须提供文档与示例说明如何通过 Go 编译时 ldflags 注入系统变量，包括：
- 变量声明位置（core/sysconfig 包级变量）
- ldflags 语法示例（-ldflags "-X core/sysconfig.Version=v1.0.0"）
- 构建脚本参考（shell 或 Makefile 示例）

#### 场景：构建与发布
- **当** 进行构建与打包时
- **那么** 必须使用 ldflags 注入版本号、构建时间等信息

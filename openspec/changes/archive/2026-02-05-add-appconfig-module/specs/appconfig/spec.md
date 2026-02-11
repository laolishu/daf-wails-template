# appconfig Specification

## Purpose
提供运行时可修改的应用配置管理能力，与系统配置（sysconfig）形成清晰分离，使用 viper 库实现配置的加载、保存与访问。

## 新增需求

### 需求：应用配置模块
系统必须提供一个专用模块管理运行时可修改的应用配置，该模块应位于 core/config，与 Wails 解耦，使用 viper 库实现配置文件的读写。

#### 场景：应用启动时加载配置
- **当** 应用启动时
- **那么** 必须加载应用配置文件，若文件不存在则使用默认值并自动创建

#### 场景：运行时修改配置
- **当** 用户通过 API 或 UI 修改配置
- **那么** 配置必须保存到配置文件，并在下次启动时生效

### 需求：配置与系统配置的边界
应用配置（core/config）必须与系统配置（core/sysconfig）清晰分离：
- core/sysconfig：编译时确定、运行时只读、系统级信息（版本、构建、ConfigDir、ConfigFile）
- core/config：运行时加载、用户可修改、应用行为配置（日志级别、语言等）

#### 场景：区分配置类型
- **当** 开发者需要添加新配置项
- **那么** 必须根据配置特性判断放入 sysconfig 还是 config

### 需求：标准应用配置项
系统必须定义以下标准应用配置项：
- log.dir：日志目录（相对或绝对路径），默认值为 `logs`
- log.level：日志级别（debug/info/warn/error），默认值为 `info`
- i18n.language：语言代码（zh-CN/en-US/ja-JP 等），默认值为 `zh-CN`

#### 场景：获取日志配置
- **当** logger 模块初始化时
- **那么** 必须从 core/config 读取 log.dir 和 log.level

#### 场景：获取语言配置
- **当** i18n 模块初始化时
- **那么** 必须从 core/config 读取 i18n.language

### 需求：配置文件路径解析
配置管理器必须基于 sysconfig.GetConfigDir() 和 sysconfig.GetConfigFile() 确定配置文件路径：
- 完整路径 = ConfigDir + ConfigFile（使用 filepath.Join 组合）
- 若 ConfigDir 为空，使用当前工作目录
- 若配置文件不存在，使用默认值并自动创建

#### 场景：构建配置文件路径
- **当** 配置管理器初始化时
- **那么** 必须从 sysconfig 获取 ConfigDir 和 ConfigFile 并组合成完整路径

#### 场景：配置目录不存在
- **当** 配置文件路径中的目录不存在
- **那么** 必须自动创建目录（使用 os.MkdirAll）

### 需求：配置文件格式
配置文件必须使用 YAML 格式，支持嵌套结构，示例：
```yaml
log:
  dir: "logs"
  level: "info"
i18n:
  language: "zh-CN"
```

#### 场景：解析配置文件
- **当** 加载配置文件时
- **那么** 必须正确解析 YAML 格式的嵌套配置

#### 场景：配置文件格式错误
- **当** 配置文件 YAML 格式错误
- **那么** 必须返回明确的错误信息（包含行号）

### 需求：配置管理接口
系统必须提供统一的配置管理接口，包括：
- NewManager()：创建配置管理器
- Load()：加载配置文件
- Save()：保存配置文件
- GetLogDir/GetLogLevel/GetLanguage：获取配置项
- SetLogDir/SetLogLevel/SetLanguage：设置配置项

#### 场景：加载配置
- **当** 应用启动时
- **那么** 调用 NewManager() 创建管理器，调用 Load() 加载配置

#### 场景：修改并保存配置
- **当** 需要修改配置并持久化
- **那么** 调用 SetXxx() 修改配置，调用 Save() 保存到文件

### 需求：默认值定义
配置管理器必须为所有配置项提供明确的默认值，定义在 core/config/defaults.go 中：
- DefaultLogDir = "logs"
- DefaultLogLevel = "info"
- DefaultLanguage = "zh-CN"

#### 场景：配置文件不存在时使用默认值
- **当** 配置文件不存在
- **那么** 所有配置项必须使用默认值

### 需求：错误处理
配置管理器必须对以下错误场景提供明确处理：
- 配置文件不存在：使用默认值并自动创建，记录 INFO 日志
- 配置文件格式错误：返回错误，记录 ERROR 日志
- 配置目录不可写：返回错误，记录 ERROR 日志

#### 场景：权限不足无法保存配置
- **当** 配置目录没有写权限
- **那么** Save() 必须返回权限错误

### 需求：依赖 viper 库
配置管理器必须使用 github.com/spf13/viper 库实现配置读写，go.mod 必须包含此依赖。

#### 场景：使用 viper 加载配置
- **当** 加载配置文件时
- **那么** 使用 viper.ReadInConfig() 读取配置

### 需求：文档与示例
系统必须提供以下文档与示例：
- docs/appconfig-guide.md：配置文件格式、使用方式与常见问题
- configs/app/config.yml：示例配置文件（包含所有配置项及注释）

#### 场景：用户查看配置说明
- **当** 用户需要了解配置项含义
- **那么** 可通过 docs/appconfig-guide.md 查看文档

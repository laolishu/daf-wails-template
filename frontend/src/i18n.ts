export type Language = "zh-CN" | "en-US";

type Messages = {
  navHealthCheck: string;
  navExamples: string;
  navSettings: string;
  navHelp: string;
  navUpgrade: string;
  healthTitle: string;
  healthDesc: string;
  healthCta: string;
  settingsTitle: string;
  settingsDesc: string;
  settingsGeneralSection: string;
  settingsSystemSection: string;
  settingsLanguageLabel: string;
  settingsLanguageZh: string;
  settingsLanguageEn: string;
  settingsSystemInfoTitle: string;
  settingsSystemInfoDesc: string;
  settingsSystemInfoErrorTitle: string;
  systemVersionLabel: string;
  systemBuildTimeLabel: string;
  systemEnvironmentLabel: string;
  examplesTitle: string;
  examplesDesc: string;
  loggerExampleTitle: string;
  loggerExampleDesc: string;
  loggerExampleButton: string;
  loggerExampleOk: string;
  loggerExampleError: string;
  loggerExampleLogDir: string;
  systemInfoExampleTitle: string;
  systemInfoExampleDesc: string;
  systemInfoExampleButton: string;
  configExampleTitle: string;
  configExampleDesc: string;
  configExampleButton: string;
  configLanguageLabel: string;
  configLogLevelLabel: string;
  configLogDirLabel: string;
  helpTitle: string;
  helpDesc: string;
  upgradeTitle: string;
  upgradeDesc: string;
  upgradeCta: string;
};

export const messages: Record<Language, Messages> = {
  "zh-CN": {
    navHealthCheck: "健康检查",
    navExamples: "示例",
    navSettings: "设置",
    navHelp: "帮助",
    navUpgrade: "升级",
    healthTitle: "欢迎来到健康检查",
    healthDesc: "运行快速扫描以确认系统状态稳定可靠。",
    healthCta: "立即扫描",
    settingsTitle: "设置",
    settingsDesc: "管理应用偏好并更新本地配置。",
    settingsGeneralSection: "通用设置",
    settingsSystemSection: "系统信息",
    settingsLanguageLabel: "语言",
    settingsLanguageZh: "中文",
    settingsLanguageEn: "English",
    settingsSystemInfoTitle: "系统信息",
    settingsSystemInfoDesc: "用于支持与诊断的只读系统信息。",
    settingsSystemInfoErrorTitle: "系统信息错误",
    systemVersionLabel: "版本",
    systemBuildTimeLabel: "构建时间",
    systemEnvironmentLabel: "运行环境",
    examplesTitle: "Examples",
    examplesDesc: "真实系统调用示例，操作结果即时可见。",
    loggerExampleTitle: "Logger Example",
    loggerExampleDesc: "写入一次测试日志，验证前端到日志系统链路。",
    loggerExampleButton: "Write Test Log",
    loggerExampleOk: "前端成功调用 Go 日志接口",
    loggerExampleError: "写入失败",
    loggerExampleLogDir: "日志路径:",
    systemInfoExampleTitle: "System Info Example",
    systemInfoExampleDesc: "读取版本号、构建时间与运行平台信息。",
    systemInfoExampleButton: "Load System Info",
    configExampleTitle: "Config Read Example",
    configExampleDesc: "读取当前配置中的语言与日志级别。",
    configExampleButton: "Load Config",
    configLanguageLabel: "语言",
    configLogLevelLabel: "日志级别",
    configLogDirLabel: "日志路径",
    helpTitle: "帮助",
    helpDesc: "获取应用使用指南与支持资源。",
    upgradeTitle: "升级",
    upgradeDesc: "检查可用更新并保持安装为最新。",
    upgradeCta: "检查更新",
  },
  "en-US": {
    navHealthCheck: "Health Check",
    navExamples: "Examples",
    navSettings: "Settings",
    navHelp: "Help",
    navUpgrade: "Upgrade",
    healthTitle: "Welcome to Health Check",
    healthDesc: "Run a quick scan to confirm the system is ready for reliable operation.",
    healthCta: "Run Health Check",
    settingsTitle: "Settings",
    settingsDesc: "Manage application preferences and update your local configuration.",
    settingsGeneralSection: "General",
    settingsSystemSection: "System Info",
    settingsLanguageLabel: "Language",
    settingsLanguageZh: "中文",
    settingsLanguageEn: "English",
    settingsSystemInfoTitle: "System Info",
    settingsSystemInfoDesc: "Read-only system details for support and diagnostics.",
    settingsSystemInfoErrorTitle: "System Info Error",
    systemVersionLabel: "Version",
    systemBuildTimeLabel: "Build Time",
    systemEnvironmentLabel: "Environment",
    examplesTitle: "Examples",
    examplesDesc: "Real system calls with immediate results.",
    loggerExampleTitle: "Logger Example",
    loggerExampleDesc: "Write a test log to validate the logging pipeline.",
    loggerExampleButton: "Write Test Log",
    loggerExampleOk: "Frontend successfully invoked Go logger",
    loggerExampleError: "Write failed",
    loggerExampleLogDir: "Log directory:",
    systemInfoExampleTitle: "System Info Example",
    systemInfoExampleDesc: "Load version, build time, and platform info.",
    systemInfoExampleButton: "Load System Info",
    configExampleTitle: "Config Read Example",
    configExampleDesc: "Read current language and log level from config.",
    configExampleButton: "Load Config",
    configLanguageLabel: "Language",
    configLogLevelLabel: "Log Level",
    configLogDirLabel: "Log Directory",
    helpTitle: "Help",
    helpDesc: "Find guidance and support resources for operating this application.",
    upgradeTitle: "Upgrade",
    upgradeDesc: "Check for available updates and keep your installation current.",
    upgradeCta: "Check for Updates",
  },
};

export function normalizeLanguage(value: string | null | undefined): Language {
  if (value === "en-US") {
    return "en-US";
  }
  return "zh-CN";
}

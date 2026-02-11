@echo off
REM DAF-Wails Windows 构建脚本
REM 支持自动注入系统配置（版本、构建时间、Git 提交、配置目录）

setlocal enabledelayedexpansion

REM 参数配置（可通过环境变量覆盖）
if "%VERSION%"=="" set VERSION=dev
if "%BUILD_TIME%"=="" (
    for /f "tokens=*" %%a in ('powershell -Command "Get-Date -Format 'yyyy-MM-ddTHH:mm:ssZ' -AsUTC"') do set BUILD_TIME=%%a
)
if "%GIT_COMMIT%"=="" (
    for /f "tokens=*" %%a in ('git rev-parse --short HEAD 2^>nul') do set GIT_COMMIT=%%a
    if "!GIT_COMMIT!"=="" set GIT_COMMIT=unknown
)
if "%CONFIG_DIR%"=="" set CONFIG_DIR=C:\ProgramData\daf-app

echo ======================================
echo DAF-Wails Build Script (Windows)
echo ======================================
echo Version:     %VERSION%
echo Build Time:  %BUILD_TIME%
echo Git Commit:  %GIT_COMMIT%
echo Config Dir:  %CONFIG_DIR%
echo ======================================

cd /d "%~dp0\.."

REM 构建 ldflags
set LDFLAGS=-X {{.ModulePath}}/core/sysconfig.Version=%VERSION%
set LDFLAGS=%LDFLAGS% -X {{.ModulePath}}/core/sysconfig.BuildTime=%BUILD_TIME%
set LDFLAGS=%LDFLAGS% -X {{.ModulePath}}/core/sysconfig.GitCommit=%GIT_COMMIT%
set LDFLAGS=%LDFLAGS% -X {{.ModulePath}}/core/sysconfig.ConfigDir=%CONFIG_DIR%

echo Building application for Windows...
wails build --clean --platform windows/amd64 -ldflags "%LDFLAGS%"

echo.
echo Build completed successfully!
echo Binary location: build\bin\

endlocal

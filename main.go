/*
 * @Descripttion:
 * @version:
 * @Author: lfzxs@qq.com
 * @Date: 2026-02-05 15:38:06
 * @LastEditors: lfzxs@qq.com
 * @LastEditTime: 2026-02-05 23:13:00
 */
package main

import (
	"daf-wails-template/backend"
	"daf-wails-template/core/sysconfig"
	"embed"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// 初始化系统配置（从编译时 ldflags 获取版本、构建信息等）
	info := sysconfig.GetInfo()
	_ = info // 应用启动时可根据需要使用 sysconfig 信息

	app := backend.NewApp(assets)

	err := app.Run()

	if err != nil {
		println("Error:", err.Error())
	}
}

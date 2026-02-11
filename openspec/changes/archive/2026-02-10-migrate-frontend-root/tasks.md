# 任务清单：前端目录迁移

## 1. 规范与设计
- [x] 更新 project-structure 规范增量（ui-shell → frontend）
- [x] 校验提案与规范通过 openspec-cn validate

## 2. 实施（批准后）
- [x] 迁移目录：ui-shell → frontend（含全部前端内容）
- [x] 更新 wails.json 的 frontend:dir 与 assetdir
- [x] 更新 Go 侧嵌入资源路径与引用
- [x] 更新脚本/文档中与 ui-shell 相关的路径

## 3. 验证
- [x] 运行 wails dev 或相关构建命令验证前端可用
- [x] 运行相关测试与 lint

# 变更：前端目录从 ui-shell 迁移到 frontend

## 为什么
当前前端根目录命名为 ui-shell，与团队统一的 frontend 目录约定不一致，增加了维护与协作成本。需要将前端根目录迁移为 frontend，并同步更新所有引用。

## 变更内容
- 将前端根目录从 ui-shell 迁移到 frontend（包含前端工程全部内容）
- 更新 Wails 配置中的前端目录与构建输出路径
- 更新 Go 侧嵌入资源路径与相关引用
- 更新项目结构规范以反映目录命名变更

## 影响
- 受影响规范：project-structure（修改）
- 受影响代码：wails.json、backend/app.go、相关脚本/文档引用

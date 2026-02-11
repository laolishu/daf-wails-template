# frontend-guidelines 规范增量：Examples 页面

## 新增需求

### 需求：真实调用约束
Examples 页面所有示例必须调用真实 Wails API，不得使用 mock。

#### 场景：执行示例操作
- **当** 用户触发示例操作
- **那么** 必须调用真实 Wails API
- **并且** 不得修改用户配置或系统状态
- **并且** 操作结果必须立即可见

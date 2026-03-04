# gen/http/router/ - 路由注册

此目录存放 Hz 自动生成的路由注册代码。

## 文件说明

- `register.go` - 路由注册入口，调用各模块的路由注册
- `xxx/xxx.go` - 各模块的具体路由定义
- `xxx/middleware.go` - 模块级中间件配置

## 注意

- **禁止手动修改**此目录下的文件
- 路由配置应通过 IDL 文件定义
- 重新生成时会覆盖现有文件
- 全局中间件应在 `internal/transport/http/middleware/` 中配置

# gen/http/handler/ - HTTP 请求处理器

此目录存放 Hz 生成的 HTTP 请求处理器骨架代码。

## 职责

- 解析和验证请求参数
- 调用 `internal/app/` 中的服务层处理业务逻辑
- 组装并返回响应

## 示例代码

```go
func GetUser(ctx context.Context, c *app.RequestContext) {
    var req model.GetUserReq
    if err := c.BindAndValidate(&req); err != nil {
        resp.BadRequest(c, err.Error())
        return
    }

    // 调用服务层
    svc := user.NewService()
    result, err := svc.GetByID(ctx, req.ID)
    if err != nil {
        resp.InternalError(c, err.Error())
        return
    }

    resp.Success(c, result)
}
```

## 注意

- Handler 只负责 HTTP 协议处理，不包含业务逻辑
- 业务逻辑应放在 `internal/app/` 中
- 数据访问应放在 `internal/repo/` 中

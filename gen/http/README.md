# gen/http/ - Hz 生成的 HTTP 代码

此目录存放 Hz 工具自动生成的 HTTP 相关代码。

## 目录结构

```
http/
├── handler/    # 请求处理器（生成的骨架代码）
├── router/     # 路由注册
└── model/      # 请求/响应 DTO 模型
```

## 注意事项

- 此目录下的代码由 `hz` 命令生成
- 不要直接修改 router/ 和 model/ 中的文件
- handler/ 中的骨架代码可以调用 `internal/app/` 中的服务层
- 重新生成时，router/ 和 model/ 会被覆盖

## 生成命令

```bash
# 基于 IDL 生成
hz new -idl idl/api.thrift -module {{.module_name}}
hz update -idl idl/api.thrift
```

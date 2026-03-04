# gen/http/model/ - 请求/响应模型

此目录存放 Hz 根据 IDL 生成的请求和响应结构体。

## 内容说明

- 请求参数结构体（XXXReq）
- 响应数据结构体（XXXResp）
- 参数验证 tag

## 注意

- **禁止手动修改**此目录下的文件
- 模型定义应在 IDL 文件中声明
- 重新生成时会覆盖现有文件
- 数据库模型应放在 `internal/repo/db/model/`
- 业务领域模型应放在 `internal/app/*/domain/`

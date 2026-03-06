# CloudWeGo 微服务模板

基于 CloudWeGo Hertz + Kitex 的 Go 微服务脚手架模板。

## 🚀 快速使用

### 方式1：使用 scaffold CLI（推荐）

```bash
# 安装 scaffold CLI
go install github.com/your-org/scaffold/cmd/scaffold@latest

# 从模板创建新项目
scaffold create /opt/project/cloudwego-template \
  --project-name my-service \
  --module-name github.com/myorg/my-service \
  --output /path/to/output
```

### 方式2：手动替换

```bash
# 1. 复制模板到新目录
cp -r /opt/project/cloudwego-template /path/to/my-service
cd /path/to/my-service

# 2. 设置变量
PROJECT_NAME="my-service"
MODULE_NAME="github.com/myorg/my-service"
DESCRIPTION="我的微服务"
AUTHOR="张三"
SERVER_PORT="8888"
DB_NAME="my_service"

# 3. 替换模板变量
find . -type f \( -name "*.go" -o -name "*.yaml" -o -name "*.yml" -o -name "*.md" -o -name "go.mod" \) -not -path "./.git/*" -exec sed -i '' \
  -e "s|{{.project_name}}|${PROJECT_NAME}|g" \
  -e "s|{{.module_name}}|${MODULE_NAME}|g" \
  -e "s|{{.description}}|${DESCRIPTION}|g" \
  -e "s|{{.author}}|${AUTHOR}|g" \
  -e "s|{{.server_port}}|${SERVER_PORT}|g" \
  -e "s|{{.db_name}}|${DB_NAME}|g" \
  {} +

# 4. 生成代码
make tools-install
make gen-update-all

# 5. 运行项目
go run cmd/server/main.go
```

### 方式3：使用 Python 脚本

```bash
# 使用提供的 Python 脚本
python scripts/create_from_template.py \
  --template /opt/project/cloudwego-template \
  --output /path/to/output \
  --project-name my-service \
  --module-name github.com/myorg/my-service
```

## 📋 模板变量

| 变量名 | 说明 | 默认值 | 必填 |
|--------|------|--------|------|
| `project_name` | 项目名称 | `my-service` | ✅ |
| `module_name` | Go Module 名称 | `github.com/example/my-service` | ✅ |
| `description` | 项目描述 | `A Hertz microservice` | ❌ |
| `author` | 作者名称 | `张三` | ❌ |
| `server_port` | 服务端口 | `8888` | ❌ |
| `db_driver` | 数据库类型 | `mysql` | ❌ |
| `db_name` | 数据库名称 | `my_service` | ❌ |
| `enable_redis` | 启用 Redis | `true` | ❌ |

## 🎯 项目特性

- ✅ **Hz IDL 驱动**: 使用 Proto 文件定义 HTTP API
- ✅ **Kitex RPC**: 支持 Protobuf RPC 服务
- ✅ **分层架构**: gen / internal / configs 清晰分离
- ✅ **多数据库**: SQLite / MySQL / PostgreSQL
- ✅ **Redis**: 缓存支持
- ✅ **统一响应**: 标准 JSON 响应格式
- ✅ **中间件**: CORS、Recovery、Logger

## 📁 项目结构

```
{{.project_name}}/
├── cmd/server/          # 服务入口
├── configs/             # 配置文件
├── gen/                 # 自动生成代码
│   ├── http/           # HTTP 代码
│   └── rpc/            # RPC 代码
├── idl/                 # 接口定义
│   ├── http/           # HTTP IDL
│   └── rpc/            # RPC IDL
├── internal/            # 项目私有代码
│   ├── app/            # 业务逻辑
│   ├── repo/           # 数据访问
│   └── pkg/            # 工具库
├── scripts/             # 脚本
├── scaffold.yaml        # 模板配置
└── Makefile
```

## 🔧 首次运行

```bash
# 1. 安装工具
make tools-install

# 2. 安装依赖
go mod tidy

# 3. 生成代码
make gen-update-all

# 4. 运行服务
make run
```

## 📚 文档

- [README.md](README.md) - 项目完整文档
- [scaffold.yaml](scaffold.yaml) - 模板配置文件

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 License

MIT

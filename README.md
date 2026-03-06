# CloudWeGo 微服务脚手架

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GitHub Stars](https://img.shields.io/github/stars/yi-nology/cloudwego-template?style=social)](https://github.com/yi-nology/cloudwego-template)

> 🚀 基于 CloudWeGo Hertz + Kitex 的 Go 微服务脚手架模板，开箱即用！

## 🔗 项目链接

- 📖 **项目主页**: [https://github.com/yi-nology/cloudwego-template](https://github.com/yi-nology/cloudwego-template)
- 🐛 **提交 Bug**: [https://github.com/yi-nology/cloudwego-template/issues](https://github.com/yi-nology/cloudwego-template/issues)
- 💬 **讨论交流**: [GitHub Discussions](https://github.com/yi-nology/cloudwego-template/discussions)
- ⭐ **欢迎 Star**: 如果这个项目对你有帮助，请给个 Star 支持一下！

---

## ✨ 特性

- **🎯 IDL 驱动开发**: 使用 Proto 文件定义 API，自动生成代码
- **🏗️ 清晰架构**: gen（生成）/ internal（业务）/ configs（配置）分层设计
- **🗄️ 多数据库支持**: 
  - **SQLite**（默认，零配置）
  - MySQL
  - PostgreSQL
- **📦 可选 Redis**: 按需启用，灵活配置
- **🛡️ 生产就绪**: 中间件、统一响应、错误处理、日志
- **🔧 开发友好**: 热重载、代码生成、配置管理
- **📚 完整文档**: 详细的使用指南和最佳实践

---

## 📋 目录

- [快速开始](#-快速开始)
- [从模板创建项目](#-从模板创建项目)
- [项目结构](#-项目结构)
- [配置说明](#-配置说明)
- [开发指南](#-开发指南)
- [部署指南](#-部署指南)
- [常见问题](#-常见问题)
- [文档资源](#-文档资源)

---

## 🚀 快速开始

### 前置要求

- Go 1.21+
- Make（可选）

### 1️⃣ 从模板创建新项目

**方式 A：使用 Python 脚本（推荐）**

```bash
python scripts/create_from_template.py \
  --template /opt/project/cloudwego-template \
  --output /path/to/my-service \
  --project-name my-service \
  --module-name github.com/myorg/my-service
```

**方式 B：手动创建**

```bash
# 1. 复制模板
cp -r /opt/project/cloudwego-template /path/to/my-service
cd /path/to/my-service

# 2. 替换模板变量
find . -type f \( -name "*.go" -o -name "*.yaml" -o -name "*.md" \) -not -path "./.git/*" -exec sed -i '' \
  -e "s|{{.project_name}}|my-service|g" \
  -e "s|{{.module_name}}|github.com/myorg/my-service|g" \
  {} +

# 3. 初始化项目
make tools-install
go mod tidy
```

### 2️⃣ 生成代码

```bash
# 安装代码生成工具
make tools-install

# 生成 HTTP 代码
make gen-update-all

# 或单独生成
make gen-update IDL=common.proto
```

### 3️⃣ 启动服务

```bash
# 开发模式
make run

# 或直接运行
go run cmd/server/main.go
```

### 4️⃣ 验证服务

```bash
# 健康检查
curl http://localhost:8888/health

# 首页
curl http://localhost:8888/

# Ping
curl http://localhost:8888/api/v1/ping
```

---

## 🎨 从模板创建项目

### 模板变量

| 变量 | 说明 | 默认值 | 必填 |
|------|------|--------|------|
| `project_name` | 项目名称 | `my-service` | ✅ |
| `module_name` | Go Module 名称 | `github.com/example/my-service` | ✅ |
| `description` | 项目描述 | `A Hertz microservice` | ❌ |
| `author` | 作者名称 | `张三` | ❌ |
| `server_port` | 服务端口 | `8888` | ❌ |
| `db_driver` | 数据库类型 | `sqlite` | ❌ |
| `db_name` | 数据库名称 | `my_service` | ❌ |
| `enable_redis` | 启用 Redis | `false` | ❌ |

### 创建示例

**最小配置（使用默认值）：**
```bash
python scripts/create_from_template.py \
  --template /opt/project/cloudwego-template \
  --output /tmp/my-service \
  --project-name my-service \
  --module-name github.com/myorg/my-service
```

**完整配置：**
```bash
python scripts/create_from_template.py \
  --template /opt/project/cloudwego-template \
  --output /tmp/my-service \
  --project-name my-service \
  --module-name github.com/myorg/my-service \
  --description "我的微服务" \
  --author "张三" \
  --server-port 9000 \
  --db-driver mysql \
  --db-name production_db \
  --enable-redis true
```

---

## 📁 项目结构

```
my-service/
├── 📂 cmd/                      # 应用入口
│   └── server/
│       ├── main.go              # 程序入口
│       └── bootstrap/           # 初始化逻辑
│           └── bootstrap.go
├── 📂 configs/                  # 配置文件
│   ├── config.yaml              # 主配置文件
│   └── config.yaml.example      # 配置示例
├── 📂 gen/                      # 自动生成（勿手动修改）
│   ├── http/                    # HTTP 代码
│   │   ├── handler/             # 请求处理器
│   │   ├── router/              # 路由注册
│   │   └── model/               # 请求/响应模型
│   └── rpc/                     # RPC 代码
├── 📂 idl/                      # 接口定义
│   ├── api/api.proto            # HTTP 注解
│   ├── http/                    # HTTP IDL
│   │   ├── common.proto         # 通用接口
│   │   └── health.proto         # 健康检查
│   └── rpc/                     # RPC IDL
├── 📂 internal/                 # 私有代码（手写）
│   ├── app/                     # 业务逻辑层
│   ├── repo/                    # 数据访问层
│   │   ├── db/                  # 数据库
│   │   │   ├── model/           # GORM 模型
│   │   │   └── dao/             # 数据访问对象
│   │   └── redis/               # Redis（可选）
│   ├── conf/                    # 配置管理
│   └── pkg/                     # 工具库
│       ├── errors/              # 错误定义
│       ├── logger/              # 日志封装
│       └── resp/                # 响应封装
├── 📂 scripts/                  # 脚本工具
│   ├── gen.sh                   # 代码生成
│   └── create_from_template.py  # 模板创建
├── 📂 docs/                     # 文档
│   ├── DATABASE.md              # 数据库配置
│   ├── REDIS.md                 # Redis 使用
│   └── FEATURES.md              # 功能说明
├── 📄 Makefile                  # Make 命令
├── 📄 go.mod                    # Go 模块
├── 📄 scaffold.yaml             # 模板配置
└── 📄 README.md                 # 项目文档
```

---

## ⚙️ 配置说明

### 数据库配置

#### SQLite（默认，推荐开发使用）

```yaml
database:
  driver: "sqlite"
  db_name: "./data/app.db"  # 数据文件路径
```

**优点：**
- ✅ 零配置，无需安装数据库
- ✅ 纯 Go 实现，无 CGO 依赖
- ✅ 数据存储在单文件，便于管理

#### MySQL

```yaml
database:
  driver: "mysql"
  host: "localhost"
  port: 3306
  user: "root"
  password: "your_password"
  db_name: "my_service"
```

#### PostgreSQL

```yaml
database:
  driver: "postgres"
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "your_password"
  db_name: "my_service"
  ssl_mode: "disable"
```

> 📖 详细配置参考：[docs/DATABASE.md](docs/DATABASE.md)

---

### Redis 配置（可选）

**默认：关闭**

**启用方式：**

在 `configs/config.yaml` 中添加：

```yaml
redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0
```

> 📖 详细使用参考：[docs/REDIS.md](docs/REDIS.md)

---

### 服务配置

```yaml
server:
  host: "0.0.0.0"
  port: 8888
  mode: "debug"  # debug, release, test

log:
  level: "info"  # debug, info, warn, error
  filename: "logs/app.log"
  max_size: 100        # MB
  max_backups: 10      # 保留旧文件数
  max_age: 30          # 保留天数
  compress: true       # 压缩旧文件

app:
  name: "my-service"
  version: "1.0.0"
```

---

## 🛠️ 开发指南

### 代码生成

#### HTTP 接口生成

```bash
# 生成所有 HTTP 接口
make gen-update-all

# 生成单个接口
make gen-update IDL=common.proto
make gen-update IDL=http/health.proto

# 初始化新项目
make gen-new IDL=common.proto
```

#### RPC 接口生成

```bash
# 生成 RPC 代码
make gen-rpc IDL=rpc/user.proto

# 生成所有 RPC
make gen-rpc-all
```

### 定义新接口

**1. 创建 Proto 文件**

```protobuf
// idl/http/user.proto
syntax = "proto3";

package http.user;

option go_package = "user";

import "api/api.proto";

message GetUserReq {
    int64 id = 1 [(api.path) = "id"];
}

message GetUserResp {
    int64 id = 1 [(api.body) = "id"];
    string name = 2 [(api.body) = "name"];
    string email = 3 [(api.body) = "email"];
}

service UserService {
    rpc GetUser(GetUserReq) returns(GetUserResp) {
        option (api.get) = "/api/v1/users/:id";
    }
}
```

**2. 生成代码**

```bash
make gen-update IDL=http/user.proto
```

**3. 实现业务逻辑**

在 `internal/app/user/service.go` 中实现：

```go
package user

import (
    "context"
    user "github.com/myorg/my-service/gen/http/model/user"
)

type Service struct{}

func NewService() *Service {
    return &Service{}
}

func (s *Service) GetUser(ctx context.Context, req *user.GetUserReq) (*user.GetUserResp, error) {
    // 实现业务逻辑
    return &user.GetUserResp{
        Id:    req.Id,
        Name:  "John Doe",
        Email: "john@example.com",
    }, nil
}
```

**4. 在 Handler 中调用**

生成的 Handler 会自动调用你的 Service。

---

### 数据库操作

#### 定义模型

```go
// internal/repo/db/model/user.go
package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"uniqueIndex;size:50"`
    Email    string `gorm:"uniqueIndex;size:100"`
}

func (User) TableName() string {
    return "users"
}
```

#### 创建 DAO

```go
// internal/repo/db/dao/user.go
package dao

import (
    "context"
    "github.com/myorg/my-service/internal/repo/db"
    "github.com/myorg/my-service/internal/repo/db/model"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository() *UserRepository {
    return &UserRepository{db: db.GetDB()}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
    return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
    var user model.User
    err := r.db.WithContext(ctx).First(&user, id).Error
    return &user, err
}
```

---

### Redis 操作

```go
import "{{.module_name}}/internal/repo/redis"

// 基本操作
redis.Set(ctx, "key", "value", time.Hour)
val, err := redis.Get(ctx, "key")
redis.Del(ctx, "key")

// Hash
redis.HSet(ctx, "user:1", "name", "John")
name, _ := redis.HGet(ctx, "user:1", "name")

// List
redis.LPush(ctx, "queue", "item1", "item2")
items, _ := redis.LRange(ctx, "queue", 0, -1)
```

---

### 统一响应格式

所有 API 返回统一格式：

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

使用方式：

```go
import "{{.module_name}}/internal/pkg/resp"

resp.Success(c, data)
resp.Page(c, list, total, page, pageSize)
resp.BadRequest(c, "参数错误")
resp.Unauthorized(c, "未授权")
resp.NotFound(c, "未找到")
resp.InternalError(c, "内部错误")
```

---

## 🚢 部署指南

### 构建

```bash
# 编译
make build

# 输出：bin/my-service
```

### Docker

```bash
# 构建镜像
make docker-build

# 运行容器
make docker-run

# 或手动运行
docker run -p 8888:8888 my-service:latest
```

### Docker Compose

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8888:8888"
    environment:
      - CONFIG_PATH=/app/configs/config.yaml
    volumes:
      - ./configs:/app/configs
      - ./data:/app/data
```

### 生产配置建议

```yaml
# configs/config.yaml
server:
  host: "0.0.0.0"
  port: 8888
  mode: "release"  # 生产模式

database:
  driver: "mysql"
  host: "${DB_HOST}"
  port: 3306
  user: "${DB_USER}"
  password: "${DB_PASSWORD}"
  db_name: "production"

log:
  level: "info"
  filename: "logs/app.log"
  max_size: 100
  max_backups: 30
  max_age: 90
  compress: true
```

---

## ❓ 常见问题

### 1. 首次运行报错找不到 gen/ 目录？

**原因：** 模板中的 gen/ 目录需要通过 IDL 生成。

**解决：**
```bash
make gen-update-all
```

### 2. 数据库连接失败？

**SQLite：** 确保数据目录存在
```bash
mkdir -p data
```

**MySQL/PostgreSQL：** 检查配置和数据库服务状态
```bash
# 测试连接
mysql -h localhost -u root -p
```

### 3. Redis 连接错误？

**原因：** Redis 未启动或配置错误。

**解决：**
- 检查 Redis 服务：`redis-cli ping`
- 或禁用 Redis：删除 `config.yaml` 中的 redis 配置

### 4. 端口被占用？

**解决：**
```bash
# 查找占用端口的进程
lsof -i :8888

# 修改端口
# configs/config.yaml
server:
  port: 9000
```

### 5. 如何切换数据库？

**从 SQLite 到 MySQL：**

```yaml
# configs/config.yaml
database:
  driver: "mysql"  # 改为 mysql
  host: "localhost"
  port: 3306
  user: "root"
  password: "password"
  db_name: "my_service"
```

数据迁移：
```bash
# 导出 SQLite
sqlite3 data/app.db .dump > backup.sql

# 导入 MySQL（需要适配语法）
mysql -u root -p my_service < backup.sql
```

---

## 📚 文档资源

- [DATABASE.md](docs/DATABASE.md) - 数据库配置详细说明
- [REDIS.md](docs/REDIS.md) - Redis 使用指南
- [FEATURES.md](docs/FEATURES.md) - 功能配置快速指南
- [CHANGELOG.md](docs/CHANGELOG.md) - 更新日志

---

## 🤝 贡献

我们非常欢迎社区贡献！

### 反馈问题

- 🐛 **提交 Bug**: [GitHub Issues](https://github.com/yi-nology/cloudwego-template/issues)
- 💡 **功能建议**: [GitHub Discussions](https://github.com/yi-nology/cloudwego-template/discussions)
- 📧 **邮件联系**: [提交 Issue](https://github.com/yi-nology/cloudwego-template/issues/new)

### 参与开发

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

### 贡献指南

- 遵循 Go 代码规范
- 添加必要的测试
- 更新相关文档
- 保持提交历史清晰

---

## 📄 License

MIT License - 详见 [LICENSE](LICENSE) 文件

---

## 👥 作者

- **murphyyi** - *Initial work*

---

## 🙏 致谢

- [CloudWeGo](https://www.cloudwego.io/) - 云原生微服务框架
- [Hertz](https://github.com/cloudwego/hertz) - HTTP 框架
- [Kitex](https://github.com/cloudwego/kitex) - RPC 框架
- [GORM](https://gorm.io/) - ORM 库

---

**Made with ❤️ by murphyyi**

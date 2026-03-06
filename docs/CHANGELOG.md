# v1.3 更新说明

## 主要变更

### 1. 数据库默认配置改为 SQLite

**变更：**
- 默认数据库从 MySQL 改为 SQLite
- 零配置，开箱即用
- 纯 Go 实现，无需 CGO

**配置：**
```yaml
database:
  driver: "sqlite"
  db_name: "./data/app.db"
```

**优点：**
- ✅ 适合开发和中小型项目
- ✅ 简化部署，无需安装数据库服务
- ✅ 数据存储在单个文件，便于备份

---

### 2. Redis 默认关闭

**变更：**
- Redis 默认关闭（`enable_redis: false`）
- 通过配置文件动态启用

**启用方式：**

在 `configs/config.yaml` 中添加：
```yaml
redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0
```

**代码变更：**
```go
// internal/conf/config.go
type Config struct {
    Redis *RedisConfig // 指针类型，可选
}

// cmd/server/bootstrap/bootstrap.go
if cfg.Redis != nil {
    // 初始化 Redis
}
```

---

### 3. 模板变量完善

**所有硬编码路径改回模板变量：**
```go
// 旧
import "github.com/zy84338719/cloudwego-template/internal/conf"

// 新
import "{{.module_name}}/internal/conf"
```

**scaffold.yaml 变量：**
- `project_name`: 项目名称
- `module_name`: Go Module 名称
- `server_port`: 服务端口（默认 8888）
- `db_driver`: 数据库类型（默认 sqlite）
- `enable_redis`: 启用 Redis（默认 false）

---

### 4. 脚本兼容性修复

**zsh 兼容性：**
- 修复 `local -n` nameref 语法问题
- 改用 `typeset -n` 和 eval 实现跨 shell 兼容

**修复内容：**
```bash
# 旧（macOS bash 3.2 不支持）
local -n _result=$1

# 新（完全兼容）
typeset -n _result=$1
# 或
eval "${result_name}=(\"\${files[@]}\")"
```

---

### 5. 路由重复问题修复

**问题：**
- `common.proto` 和 `health.proto` 都定义了 `/health` 路由

**解决：**
- 删除 `common.proto` 中的 `/health` 定义
- 由 `health.proto` 提供完整的健康检查接口

**保留路由：**
- `/` - 首页
- `/api/v1/ping` - Ping 接口
- `/health` - 健康检查
- `/ready` - 就绪检查
- `/live` - 存活检查
- `/version` - 版本信息

---

## 新增文件

### 文档
- ✅ `docs/DATABASE.md` - 数据库配置详细说明
- ✅ `docs/REDIS.md` - Redis 配置和使用指南
- ✅ `docs/FEATURES.md` - 功能配置快速指南
- ✅ `TEMPLATE_README.md` - 模板使用说明

### 脚本
- ✅ `scripts/create_from_template.py` - Python 脚本简化模板使用

### 配置示例
- ✅ `configs/config.yaml.example` - 配置示例文件

---

## 使用方式

### 从模板创建新项目

**方式1：使用 Python 脚本**
```bash
python scripts/create_from_template.py \
  --template /opt/project/cloudwego-template \
  --output /tmp/my-service \
  --project-name my-service \
  --module-name github.com/myorg/my-service
```

**方式2：手动替换**
```bash
# 复制模板
cp -r /opt/project/cloudwego-template /path/to/my-service

# 替换变量
cd /path/to/my-service
find . -type f \( -name "*.go" -o -name "*.yaml" \) -not -path "./.git/*" -exec sed -i '' \
  -e "s|{{.project_name}}|my-service|g" \
  -e "s|{{.module_name}}|github.com/myorg/my-service|g" \
  {} +

# 生成代码
make gen-update-all

# 运行
go run cmd/server/main.go
```

---

## 升级指南

### 从 v1.2 升级

**1. 更新代码**
```bash
git pull origin main
git checkout v1.3
```

**2. 更新依赖**
```bash
go mod tidy
```

**3. 检查配置**

如果你使用了 MySQL：
- 配置保持不变，继续使用

如果你想切换到 SQLite：
```yaml
database:
  driver: "sqlite"
  db_name: "./data/app.db"
```

**4. Redis 配置**

如果之前使用了 Redis：
- 在 `config.yaml` 中确认 Redis 配置存在
- 配置保持不变

如果想关闭 Redis：
- 删除或注释 `config.yaml` 中的 `redis` 部分

---

## 下一步

- 📖 阅读 [docs/FEATURES.md](docs/FEATURES.md) 了解配置选项
- 🗄️ 阅读 [docs/DATABASE.md](docs/DATABASE.md) 了解数据库配置
- 📦 阅读 [docs/REDIS.md](docs/REDIS.md) 了解 Redis 使用
- 🚀 运行 `make run` 启动服务

---

## 贡献者

- murphyyi

## License

MIT

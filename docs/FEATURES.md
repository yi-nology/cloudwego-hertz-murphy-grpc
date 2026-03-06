# 项目配置快速指南

## 数据库配置（默认：SQLite）

项目默认使用 **SQLite** 数据库，零配置开箱即用。

### 切换到 MySQL/PostgreSQL

编辑 `configs/config.yaml`：

```yaml
database:
  driver: "mysql"  # 或 "postgres"
  host: "localhost"
  port: 3306
  user: "root"
  password: "your_password"
  db_name: "my_service"
```

详细配置参考：[DATABASE.md](docs/DATABASE.md)

---

## Redis 配置（默认：关闭）

Redis 默认**关闭**，如需启用：

### 1. 编辑配置文件

在 `configs/config.yaml` 中添加：

```yaml
redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0
```

### 2. 重启服务

```bash
go run cmd/server/main.go
```

### 3. 验证 Redis 可用

```bash
# 检查 Redis 状态
curl http://localhost:8888/health
```

详细配置参考：[REDIS.md](docs/REDIS.md)

---

## 快速切换配置

### 开发环境（默认）

```yaml
# configs/config.yaml
database:
  driver: "sqlite"
  db_name: "./data/app.db"

# Redis 不配置（关闭）
```

### 生产环境

```yaml
# configs/config.yaml
database:
  driver: "mysql"
  host: "mysql.example.com"
  port: 3306
  user: "app_user"
  password: "${DB_PASSWORD}"
  db_name: "production_db"

redis:
  host: "redis.example.com"
  port: 6379
  password: "${REDIS_PASSWORD}"
  db: 0
```

### 环境变量

支持通过环境变量覆盖配置：

```bash
export DB_PASSWORD=your_password
export REDIS_PASSWORD=your_redis_password
go run cmd/server/main.go
```

---

## 配置优先级

1. 环境变量（最高）
2. `configs/config.yaml`
3. 代码默认值（最低）

---

## 常用命令

```bash
# 查看当前配置
cat configs/config.yaml

# 测试数据库连接
go run cmd/server/main.go

# 健康检查
curl http://localhost:8888/health
```

---

## 下一步

- 📖 阅读 [README.md](README.md) 了解完整功能
- 🗄️ 查看 [DATABASE.md](docs/DATABASE.md) 了解数据库配置
- 📦 查看 [REDIS.md](docs/REDIS.md) 了解 Redis 使用

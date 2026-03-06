# 数据库配置说明

## SQLite（默认）

**优点：**
- ✅ 零配置，无需安装数据库服务
- ✅ 纯 Go 实现，无需 CGO
- ✅ 适合开发、测试和小型项目
- ✅ 数据存储在单个文件中，方便备份和迁移

**配置示例：**

```yaml
database:
  driver: "sqlite"
  db_name: "./data/app.db"  # 或 ":memory:" 表示内存数据库
```

**使用场景：**
- 开发环境
- 小型应用
- 嵌入式系统
- 原型开发

---

## MySQL

**优点：**
- ✅ 成熟稳定，生态完善
- ✅ 性能优秀，适合生产环境
- ✅ 支持高并发

**配置示例：**

```yaml
database:
  driver: "mysql"
  host: "localhost"
  port: 3306
  user: "root"
  password: "your_password"
  db_name: "my_service"
```

**使用场景：**
- 生产环境
- 中大型应用
- 需要高性能读写的场景

---

## PostgreSQL

**优点：**
- ✅ 功能强大，支持高级特性
- ✅ 数据类型丰富
- ✅ 开源免费

**配置示例：**

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

**使用场景：**
- 需要复杂查询的场景
- 地理信息系统
- 科学计算

---

## 数据库迁移

从 SQLite 迁移到 MySQL/PostgreSQL：

```bash
# 1. 导出数据
sqlite3 data/app.db .dump > backup.sql

# 2. 创建新数据库
mysql -u root -p -e "CREATE DATABASE my_service CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# 3. 导入数据（需要适配 SQL 语法）
mysql -u root -p my_service < backup.sql
```

---

## 性能优化建议

### SQLite
- 定期执行 `VACUUM` 清理碎片
- 使用 WAL 模式提高并发性能
- 合理设置 `PRAGMA cache_size`

### MySQL
- 配置合适的连接池大小
- 优化索引设计
- 使用连接池（应用层）

### PostgreSQL
- 调整 `shared_buffers`
- 配置 `work_mem`
- 使用连接池（如 PgBouncer）

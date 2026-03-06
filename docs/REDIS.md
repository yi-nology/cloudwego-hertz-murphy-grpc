# Redis 配置说明

## 启用 Redis

在 `configs/config.yaml` 中添加 Redis 配置：

```yaml
redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0
```

## 禁用 Redis（默认）

**方式1：删除 Redis 配置**

直接删除或注释掉 `configs/config.yaml` 中的 `redis` 部分：

```yaml
# redis:
#   host: "localhost"
#   port: 6379
#   password: ""
#   db: 0
```

**方式2：使用模板变量**

在创建项目时设置 `--enable-redis false`（默认）：

```bash
python scripts/create_from_template.py \
  --template /opt/project/cloudwego-template \
  --output /tmp/my-service \
  --project-name my-service \
  --module-name github.com/myorg/my-service \
  --enable-redis false  # 默认值，可省略
```

## Redis 使用示例

### 基本操作

```go
import "{{.module_name}}/internal/repo/redis"

// 设置值
err := redis.Set(ctx, "key", "value", time.Hour)

// 获取值
val, err := redis.Get(ctx, "key")

// 删除
err := redis.Del(ctx, "key")

// 检查是否存在
exists, err := redis.Exists(ctx, "key")
```

### Hash 操作

```go
// 设置 Hash 字段
err := redis.HSet(ctx, "user:1", "name", "John", "age", "30")

// 获取 Hash 字段
name, err := redis.HGet(ctx, "user:1", "name")

// 获取所有字段
all, err := redis.HGetAll(ctx, "user:1")

// 删除字段
err := redis.HDel(ctx, "user:1", "age")
```

### List 操作

```go
// 推入列表
err := redis.LPush(ctx, "queue", "item1", "item2")

// 获取列表范围
items, err := redis.LRange(ctx, "queue", 0, -1)

// 弹出元素
item, err := redis.LPop(ctx, "queue")
```

### Set 操作

```go
// 添加成员
err := redis.SAdd(ctx, "tags", "go", "redis", "docker")

// 获取所有成员
members, err := redis.SMembers(ctx, "tags")

// 检查是否是成员
isMember, err := redis.SIsMember(ctx, "tags", "go")
```

### 过期时间

```go
// 设置过期时间
err := redis.Expire(ctx, "key", time.Hour)

// 获取剩余时间
ttl, err := redis.TTL(ctx, "key")
```

## 高级配置

### 连接池

Redis 客户端默认使用连接池，配置如下：

```go
// internal/repo/redis/redis.go
&redis.Options{
    Addr:         cfg.Addr(),
    Password:     cfg.Password,
    DB:           cfg.DB,
    PoolSize:     10,     // 连接池大小
    MinIdleConns: 5,      // 最小空闲连接数
    MaxRetries:   3,      // 最大重试次数
    DialTimeout:  5 * time.Second,
    ReadTimeout:  3 * time.Second,
    WriteTimeout: 3 * time.Second,
}
```

### 生产环境建议

```yaml
redis:
  host: "redis.example.com"
  port: 6379
  password: "${REDIS_PASSWORD}"  # 使用环境变量
  db: 0
  pool_size: 20
  min_idle_conns: 10
  max_retries: 3
  dial_timeout: 5s
  read_timeout: 3s
  write_timeout: 3s
```

## 监控和调试

### 查看连接状态

```go
stats := redis.Client().PoolStats()
fmt.Printf("TotalConns: %d\n", stats.TotalConns)
fmt.Printf("IdleConns: %d\n", stats.IdleConns)
fmt.Printf("StaleConns: %d\n", stats.StaleConns)
```

### 健康检查

```go
func checkRedis(ctx context.Context) error {
    if redis.Client() == nil {
        return fmt.Errorf("redis not initialized")
    }
    return redis.Client().Ping(ctx).Err()
}
```

## 常见问题

### 1. 连接超时

**原因：**
- Redis 服务未启动
- 网络问题
- 防火墙阻止

**解决：**
```bash
# 检查 Redis 是否运行
redis-cli ping

# 检查端口
telnet localhost 6379
```

### 2. 内存占用过高

**原因：**
- 缓存数据过多
- 未设置过期时间

**解决：**
```go
// 始终设置过期时间
redis.Set(ctx, "key", "value", time.Hour)
```

### 3. 性能问题

**优化建议：**
- 使用 Pipeline 批量操作
- 合理设置连接池大小
- 使用 Lua 脚本减少网络往返

```go
// Pipeline 示例
pipe := redis.Client().Pipeline()
pipe.Set(ctx, "key1", "value1", 0)
pipe.Set(ctx, "key2", "value2", 0)
pipe.Get(ctx, "key3")
cmds, err := pipe.Exec(ctx)
```

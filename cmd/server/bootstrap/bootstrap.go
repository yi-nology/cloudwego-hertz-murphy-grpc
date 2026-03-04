package bootstrap

import (
	"fmt"
	"os"

	"github.com/cloudwego/hertz/pkg/app/server"
	hertzrouter "{{.module_name}}/gen/http/router"
	"{{.module_name}}/internal/conf"
	"{{.module_name}}/internal/pkg/logger"
	"{{.module_name}}/internal/repo/db"
	"{{.module_name}}/internal/repo/redis"
	"{{.module_name}}/internal/transport/http/middleware"
)

// Bootstrap 初始化所有组件并返回 Hertz 服务器实例
func Bootstrap() (*server.Hertz, error) {
	// 1. 初始化配置
	if err := initConfig(); err != nil {
		return nil, fmt.Errorf("init config failed: %w", err)
	}

	cfg := conf.GlobalConfig

	// 2. 初始化日志
	if err := initLogger(cfg); err != nil {
		return nil, fmt.Errorf("init logger failed: %w", err)
	}

	// 3. 初始化数据库
	if err := initDatabase(cfg); err != nil {
		return nil, fmt.Errorf("init database failed: %w", err)
	}

	// 4. 初始化 Redis
	if err := initRedis(cfg); err != nil {
		return nil, fmt.Errorf("init redis failed: %w", err)
	}

	// 5. 初始化 HTTP 服务器
	h := initServer(cfg)

	return h, nil
}

func initConfig() error {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/config.yaml"
	}

	if err := conf.Init(configPath); err != nil {
		fmt.Printf("Failed to load config from %s: %v\n", configPath, err)
		if err := conf.InitWithDefault(); err != nil {
			return err
		}
	}
	return nil
}

func initLogger(cfg *conf.Config) error {
	return logger.Init(&logger.Config{
		Level:      cfg.Log.Level,
		Filename:   cfg.Log.Filename,
		MaxSize:    cfg.Log.MaxSize,
		MaxBackups: cfg.Log.MaxBackups,
		MaxAge:     cfg.Log.MaxAge,
		Compress:   cfg.Log.Compress,
	})
}

func initDatabase(cfg *conf.Config) error {
	return db.Init(&cfg.Database)
}

func initRedis(cfg *conf.Config) error {
	return redis.Init(&cfg.Redis)
}

func initServer(cfg *conf.Config) *server.Hertz {
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	h := server.New(server.WithHostPorts(addr))

	// 注册全局中间件
	h.Use(middleware.Recovery())
	h.Use(middleware.Logger())
	h.Use(middleware.CORS())

	// 注册路由
	hertzrouter.GeneratedRegister(h)

	return h
}

// Cleanup 清理所有资源
func Cleanup() {
	logger.Sync()
	_ = db.Close()
	_ = redis.Close()
}

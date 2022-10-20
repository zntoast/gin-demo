package initialize

import (
	"context"
	"gindemo/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// 初始化redis
func InitializeRedis() *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     global.GVA_SERVER.Redis.Addr,
		Password: global.GVA_SERVER.Redis.Password,
		DB:       global.GVA_SERVER.Redis.DB,
	})
	if _, err := cli.Ping(context.Background()).Result(); err != nil {
		global.GVA_LOG.Warn("redis 加载失败", zap.Error(err))
		return nil
	} else {
		global.GVA_LOG.Info("redis 加载成功 .........")
		return cli
	}
}

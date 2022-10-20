package utils

import (
	"context"
	"gindemo/global"
	"time"

	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type RedisStore struct {
	Expiration time.Duration
	Prekey     string
	Context    context.Context
}

var (
	expiration        = time.Minute * 15 //十五分钟过期
	cont              = context.Background()
	RedisStoreDefault = NewRedisStore(expiration, cont)
)

func NewRedisStore(ex time.Duration, con context.Context) base64Captcha.Store {
	store := new(RedisStore)
	store.Expiration = ex
	store.Prekey = "redis-"
	store.Context = con
	return store
}

func (r *RedisStore) Set(id string, value string) error {
	err := global.GVA_REDIS.Set(r.Context, r.Prekey+id, value, r.Expiration).Err()
	if err != nil {
		global.GVA_LOG.Error("redisStoreSetErr", zap.Error(err))
		return err
	}
	return nil
}

func (r *RedisStore) Get(id string, clear bool) string {
	if value, err := global.GVA_REDIS.Get(r.Context, r.Prekey+id).Result(); err != nil {
		global.GVA_LOG.Error("redisStoreGetErr", zap.Error(err))
		return ""
	} else if clear {
		err := global.GVA_REDIS.Del(r.Context, r.Prekey+id).Err()
		if err != nil {
			global.GVA_LOG.Error("redisStoreClearErr", zap.Error(err))
			return ""
		}
		return value
	} else {
		return value
	}
}

func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	value := r.Get(id, clear)
	return answer == value
}

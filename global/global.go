package global

import (
	"gindemo/config"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_SERVER config.Server
	GVA_DB     *gorm.DB
	GVA_LOG    *zap.Logger
	GVA_REDIS  *redis.Client
)

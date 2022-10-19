package global

import (
	"gindemo/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_SERVER config.Server
	GVA_DB     *gorm.DB
	GVA_LOG    *zap.Logger
)

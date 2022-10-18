package global

import (
	"gindemo/config"

	"gorm.io/gorm"
)

var (
	GVA_SERVER config.Server
	GVA_DB     *gorm.DB
)

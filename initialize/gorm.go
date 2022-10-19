package initialize

import (
	"gindemo/global"
	"gindemo/model/system"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化gorm
func InitGormMysql() *gorm.DB {
	m := global.GVA_SERVER.Mysql
	if m.DbName == "" {
		return nil
	}
	//mysql配置
	mysqlconfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	//gorm配置
	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	}
	if gormDb, err := gorm.Open(mysql.New(mysqlconfig), gormConfig); err != nil {
		global.GVA_LOG.Warn("加载gorm错误err:" + err.Error())
		return nil
	} else {
		global.GVA_LOG.Info("初始化gorm成功............")
		return gormDb
	}
}

func RegisterTables(goreDb *gorm.DB) {
	goreDb.AutoMigrate(
		system.User{},
		system.LoginRecord{},
	)
}

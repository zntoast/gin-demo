package initialize

import (
	"fmt"
	"gindemo/global"
	"gindemo/model/system"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化gorm
func InitGormMysql() {
	m := global.GVA_SERVER.Mysql
	fmt.Printf("m.Dsn(): %v\n", m.Dsn())
	if m.DbName == "" {
		return
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
		fmt.Printf("\x1b[%d 加载gorm错误err:%v \x1b[0m\n", 31, err)
		return
	} else {
		global.GVA_DB = gormDb
		RegisterTables()
	}
}

// 初始化表
func RegisterTables() {
	global.GVA_DB.AutoMigrate(
		system.User{},
	)
}

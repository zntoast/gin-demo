package main

import (
	"gindemo/global"
	"gindemo/initialize"
)

func main() {
	// fmt.Printf("time.Now(): %v\n", time.Now())
	Run()
}

func Run() {
	//加载文件配置
	initialize.InitConfigFile("config1.yaml")
	//初始化日志文件
	global.GVA_LOG = initialize.InitializeZap()
	//初始化redis
	global.GVA_REDIS = initialize.InitializeRedis()
	//初始化gorm
	global.GVA_DB = initialize.InitGormMysql()
	if global.GVA_DB == nil {
		global.GVA_LOG.Warn("mysql未初始化 .............")
	} else {
		//关联数据库表(数据库中不存在则创建)
		initialize.RegisterTables(global.GVA_DB)
	}
	//初始化路由
	initialize.InitializeRoute()
}

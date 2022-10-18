package main

import (
	"gindemo/initialize"
	"gindemo/utils"
)

func main() {
	Run()
}

func Run() {
	utils.Print("打印白色", utils.Blank)
	// fmt.Printf("utils.White: %v\n", utils.White)
	//加载文件配置
	initialize.InitConfigFile("config1.yaml")
	//初始化gorm
	initialize.InitGormMysql()
	//初始化路由
	initialize.InitializeRoute()
}

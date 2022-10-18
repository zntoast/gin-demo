package main

import (
	"gindemo/initialize"
)

func main() {
	Run()
}

func Run() {
	initialize.InitConfigFile("")
	//初始化路由
	initialize.InitializeRoute()
}

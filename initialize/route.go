package initialize

import (
	"fmt"
	"gindemo/global"
	"gindemo/middleware"
	"gindemo/router"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitializeRoute() {
	r := gin.Default()
	userRouter := router.RouterGroupApp.UserRouter

	//开放接口
	publicGroup := r.Group("")
	publicGroup.Use()
	{
		userRouter.InitUserRouter(publicGroup) // 系统用户
	}

	//需要登录
	privateGroup := r.Group("")
	privateGroup.Use(middleware.HelloMiddleware())
	{

	}

	//管理员

	addr := fmt.Sprintf(":%s", global.GVA_SERVER.System.Addr)
	r.Run(addr)
}

package initialize

import (
	"gindemo/middleware"
	"gindemo/router"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitializeRoute() {
	r := gin.Default()
	userRouter := router.RouterGroupApp.UserRouter

	publicGroup := r.Group("")
	publicGroup.Use()
	{
		userRouter.InitUserRouter(publicGroup)
	}

	//需要登录
	privateGroup := r.Group("")
	privateGroup.Use(middleware.HelloMiddleware())
	{

	}

	r.Run(":8000")
}

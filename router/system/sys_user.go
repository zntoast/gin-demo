package system

import (
	"github.com/gin-gonic/gin"

	v1 "gindemo/api"
)

type UserRouter struct{}

func (user *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	UserRouter := r.Group("user").Use()
	sysApi := v1.SystemUserApi{}
	{
		UserRouter.POST("/login", sysApi.Login)       // 用户登录
		UserRouter.POST("/register", sysApi.Register) // 用户注册
	}
}

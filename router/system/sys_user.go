package system

import (
	"github.com/gin-gonic/gin"

	v1 "gindemo/api"
)

type UserRouter struct{}

func (user *UserRouter) UserRouterGroup(r *gin.RouterGroup) {
	UserRouter := r.Group("user").Use()
	sysUserApi := v1.SystemUserApi{}
	{
		UserRouter.POST("/sms_code", sysUserApi.SMSCode)        // 短信服务
		UserRouter.POST("/captcha", sysUserApi.GenerateCaptcha) // 图形验证码
		UserRouter.POST("/login", sysUserApi.UserLogin)         // 用户登录
		UserRouter.POST("/register", sysUserApi.UserRegister)   // 用户注册
	}
}

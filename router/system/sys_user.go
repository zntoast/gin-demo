package system

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (user *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	UserRouter := r.Group("user").Use()
	{
		UserRouter.POST("sms_code", sysUserApi.SMSCode)       // 短信服务
		UserRouter.GET("captcha", sysUserApi.GenerateCaptcha) // 图形验证码
		UserRouter.POST("login", sysUserApi.UserLogin)        // 用户登录
		UserRouter.POST("register", sysUserApi.UserRegister)  // 用户注册
	}
}

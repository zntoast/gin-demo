package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (user *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	UserRouter := r.Group("user").Use()
	{
		UserRouter.GET("captcha", sysUserApi.GenerateCaptcha) // 获取图形验证码
		UserRouter.POST("sms_code", sysUserApi.SMSCode)       // 短信服务
		UserRouter.POST("login", sysUserApi.UserLogin)        // 用户登录
		UserRouter.POST("register", sysUserApi.UserRegister)  // 用户注册
		UserRouter.GET("ip", func(ctx *gin.Context) {
			ip := ctx.GetHeader("X-Forwarded-For")
			addr := ctx.GetHeader("X-Real-IP")
			host := ctx.GetHeader("host")
			if ip != "" {
				fmt.Printf("ip: %v\n", ip)
			}
			ctx.JSON(200, gin.H{
				"method": "GET",
				"ip":     ip,
				"addr":   addr,
				"host":   host,
			})
		})
	}
}

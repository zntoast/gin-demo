package api

import (
	"gindemo/utils"
	"gindemo/vo/request"
	"gindemo/vo/response"

	"github.com/gin-gonic/gin"
)

type SystemUserApi struct{}

func (s *SystemUserApi) Login(c *gin.Context) {
	req := &request.LoginRequest{}
	if err := c.ShouldBind(req); err != nil {
		response.FailWithMessage(utils.GetValidMsg(err, req), c)
		return
	}
	response.OkWithMessage("登录成功", c)
}

func (s *SystemUserApi) Register(c *gin.Context) {
	req := &request.RegisterRequest{}
	if err := c.ShouldBind(req); err != nil {
		response.FailWithMessage(utils.GetValidMsg(err, req), c)
		return
	}
	if req.Password != req.ConfirmPassword {
		response.FailWithMessage("两次输入密码不一致", c)
		return
	}
	response.OkWithMessage("注册成功", c)
}

package api

import (
	"gindemo/utils"
	"gindemo/vo/request"
	"gindemo/vo/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var (
	store = base64Captcha.DefaultMemStore
)

type SystemUserApi struct{}

// 图像验证码
func (s *SystemUserApi) GenerateCaptcha(c *gin.Context) {
	req := &request.LoginRequest{}
	if err := c.ShouldBind(req); err != nil {
		response.FailWithMessage(utils.GetValidMsg(err, req), c)
		return
	}
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		response.FailWithMessage("获取验证码失败", c)
		return
	} else {
		resp := &struct {
			Id     string
			B64s   string
			Lenght int
		}{
			Id:     id,
			B64s:   b64s,
			Lenght: 6,
		}
		response.FailWithDetailed(resp, "获取验证码成功", c)
	}
	response.Ok(c)
}

// 短信服务
func (s *SystemUserApi) SMSCode(c *gin.Context) {
	req := &request.SMSCodeRequest{}
	if err := c.ShouldBind(req); err != nil {
		response.FailWithMessage(utils.GetValidMsg(err, req), c)
		return
	}
	/* 业务实现 */
	response.Ok(c)
}

// 用户登录
func (s *SystemUserApi) UserLogin(c *gin.Context) {
	req := &request.LoginRequest{}
	if err := c.ShouldBind(req); err != nil {
		response.FailWithMessage(utils.GetValidMsg(err, req), c)
		return
	}
	response.OkWithMessage("登录成功", c)
}

// 用户注册
func (s *SystemUserApi) UserRegister(c *gin.Context) {
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

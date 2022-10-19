package system

import (
	"fmt"
	"gindemo/global"
	"gindemo/model/system"
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
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.9, 80)
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
	// 实现阿里的短信服务
	response.OkWithDetailed(nil, "发送成功", c)
}

// 用户登录
func (s SystemUserApi) UserLogin(c *gin.Context) {
	req := &request.LoginRequest{}
	if err := c.ShouldBind(req); err != nil {
		response.FailWithMessage(utils.GetValidMsg(err, req), c)
		return
	}
	//判断手机号是否合法
	if err := utils.VerifyPhoneNumber(req.PhoneNumber); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//图像码验证
	if !store.Verify(req.CaptchaId, req.Captcha, true) {
		response.FailWithMessage("验证码错误", c)
		return
	}
	//验证手机号和密码
	if user, err := userService.Login(req.PhoneNumber, req.Password); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		fmt.Printf("user: %v\n", user)
		response.OkWithMessage("登录成功", c)
		//生成token
	}
}

// 用户注册
func (s SystemUserApi) UserRegister(c *gin.Context) {
	req := &request.RegisterRequest{}
	if err := c.ShouldBind(req); err != nil {
		response.FailWithMessage(utils.GetValidMsg(err, req), c)
		return
	}
	//判断手机号是否合法
	if err := utils.VerifyPhoneNumber(req.PhoneNumber); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//图像码验证
	if !store.Verify(req.CaptchaId, req.Captcha, true) {
		response.FailWithMessage("验证码错误", c)
		return
	}
	if req.Password != req.ConfirmPassword {
		response.FailWithMessage("两次密码输入不一致", c)
		return
	}
	user := &system.User{
		Phone:    req.PhoneNumber,
		Password: req.Password,
	}
	if err := userService.Register(user); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("注册成功", c)
	global.GVA_LOG.Info(fmt.Sprintf("手机号:%s 注册成功\n", req.PhoneNumber))
}

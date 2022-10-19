package request

//登录请求
type LoginRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required" msg:"请输入手机号"`
	Password    string `json:"password" binding:"required" msg:"请输入密码"`
	CaptchaId   string `json:"captchaId" binding:""`
	Captcha     string `json:"captcha" binding:"required" msg:"请输入验证码"`
}

//注册请求
type RegisterRequest struct {
	Username        string `json:"username" binding:"required" msg:"请输入用户名"`
	Password        string `json:"password" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"`
	ConfirmPassword string `json:"nextPassword" binding:"required" msg:"请输入密码"`
	Email           string `json:"email" binding:"email" msg:"邮箱地址格式不正确"`
	Code            int    `json:"code" binding:"required" msg:"请输入短信验证码"`
}

//获取图像验证码请求
type SMSCodeRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required" msg:"手机号不能为空"`
}

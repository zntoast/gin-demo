package request

type LoginRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required" msg:"电话不能为空"`
	Password    string `json:"password" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"`
	CaptchaId   string `json:"captchaId" binding:""`
	Captcha     string `json:"captcha" binding:"required" msg:"短信验证码不能为空"`
}

type RegisterRequest struct {
	Username        string `json:"username" binding:"required" msg:"用户名不能为空"`
	Password        string `json:"password" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"`
	ConfirmPassword string `json:"nextPassword" binding:"required" msg:"密码不能为空"`
	Email           string `json:"email" binding:"email" msg:"邮箱地址格式不正确"`
	Code            int    `json:"code" binding:"required" msg:"请输入短信验证码"`
}

type SMSCodeRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required" msg:"手机号不能为空"`
}

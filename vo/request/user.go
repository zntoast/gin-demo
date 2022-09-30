package request

type LoginRequest struct {
	Username string `json:"username" binding:"required" msg:"用户名不能为空"`
	Password string `json:"password" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"`
	Email    string `json:"email" binding:"email" msg:"邮箱地址格式不正确"`
}

type RegisterRequest struct {
	Username        string `json:"username" binding:"required" msg:"用户名不能为空"`
	Password        string `json:"password" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"`
	ConfirmPassword string `json:"nextPassword" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"`
	Email           string `json:"email" binding:"email" msg:"邮箱地址格式不正确"`
	Code            int    `json:"code" binding:"required" msg:"请输入短信验证码"`
}

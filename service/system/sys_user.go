package system

import (
	"errors"
	"gindemo/global"
	"gindemo/model/system"
	"gindemo/utils"
)

type UserService struct {
}

func (u *UserService) Login(phone, password string) (*system.User, error) {
	var user system.User
	if err := global.GVA_DB.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, errors.New("该手机号未注册")
	}
	//验证hash值
	if utils.BcryptCheck(password, user.Password) {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}

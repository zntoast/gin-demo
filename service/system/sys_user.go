package system

import (
	"errors"
	"gindemo/global"
	"gindemo/model/system"
	"gindemo/utils"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
}

func (u UserService) Login(phone, password string) (*system.User, error) {
	var user system.User
	if errors.Is(global.GVA_DB.Where("phone = ?", user.Phone).First(&user).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("该手机号未注册")
	}
	//验证hash值
	if !utils.BcryptCheck(password, user.Password) {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}

func (u UserService) Register(user *system.User) error {
	if !errors.Is(global.GVA_DB.Where("phone = ?", user.Phone).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("该手机号已经被注册")
	}
	//加密
	user.Password = utils.BcryptHash(user.Password)
	user.UUID = uuid.NewV4()
	maxId := uint(0)
	err := global.GVA_DB.Model(system.User{}).Table(system.User{}.TableName()).Select("max(id)").First(&maxId).Error
	if err != nil {
		global.GVA_LOG.Warn("db find max id user error", zap.Error(err))
	}
	user.Addr = "dwadjoawjw"
	err = global.GVA_DB.Create(user).Error
	return err
}

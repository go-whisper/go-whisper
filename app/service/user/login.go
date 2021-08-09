package user

import (
	"errors"
	"github.com/go-whisper/go-whisper/app/bizerr"
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Login(name, pwd string) (model.User, error) {
	u := model.User{}
	if err := instance.DB().Where("name=?", name).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			instance.Logger().Error("尝试登录时无法找到用户", zap.String("user_name", name))
			return u, bizerr.ErrUserNotFound
		}
		instance.Logger().Error("查找用户失败", zap.String("caller", caller("Login")))
		return u, bizerr.ErrDB
	}
	if !Verify(pwd, u.Password) {
		instance.Logger().Error("尝试登录失败", zap.String("user_name", name), zap.String("pwd", pwd))
		return u, bizerr.ErrUserInvalidPwd
	}
	instance.Logger().Error("登录成功", zap.String("user_name", name))
	return u, nil
}

func UpdatePassword(name, plaintextPWD string) error {
	pwd := Encrypt(plaintextPWD)
	if err := instance.DB().Model(model.User{}).Where("name=?", name).Update("password", pwd).Error; err != nil {
		instance.Logger().Error("更新用户密码失败", zap.String("caller", caller("UpdatePassword")))
		return bizerr.ErrDB
	}
	return nil
}

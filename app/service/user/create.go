package user

import (
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"go.uber.org/zap"
)

func Create(user *model.User) error {
	if err := instance.DB().Create(user).Error; err != nil {
		instance.Logger().Error("user.Save() fail", zap.Error(err))
		return err
	}

	return model.UpdateLastChange(instance.DB())
}

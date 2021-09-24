package post

import (
	"time"

	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"go.uber.org/zap"
)

func Create(post *model.Post) error {
	post.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	if err := instance.DB().Create(post).Error; err != nil {
		instance.Logger().Error("post.Save() fail", zap.Error(err))
		return err
	}
	model.UpdateLastChange(instance.DB())
	return nil
}

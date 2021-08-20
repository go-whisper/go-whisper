package post

import (
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"go.uber.org/zap"
)

func Update(id uint, post *model.Post) error {
	if err := instance.DB().Select("title", "content", "tags", "is_pinned", "url").Where("id=?", id).
		Updates(post).Error; err != nil {
		instance.Logger().Error("post.Update() fail", zap.Error(err))
		return err
	}
	return nil
}

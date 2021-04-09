package post

import (
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"go.uber.org/zap"
)

func List(limit, offset int, opt model.Option) (int64, []model.Post, error) {
	var posts []model.Post
	db := instance.DB()
	if v := opt.GetString("keyword"); v != "" {
		db = db.Where("title LIKE ? OR content LIKE ?", "%"+v+"%", "%"+v+"%")
	}
	if v := opt.GetString("tag"); v != "" {
		db = db.Where("tags LIKE ?", "%,"+v+",%")
	}
	total := int64(0)
	if err := db.Model(&posts).Count(&total).Error; err != nil {
		instance.Logger().Error("db.Count() fail", zap.String("caller", caller("List", "db.Count()")), zap.Error(err))
		return total, posts, err
	}
	if err := db.Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		instance.Logger().Error("db.Find() fail", zap.String("caller", caller("List", "db.Find()")), zap.Error(err))
		return total, posts, err
	}
	return total, posts, nil
}

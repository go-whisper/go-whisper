package post

import (
	"errors"
	"strings"

	"github.com/go-whisper/go-whisper/app/bizerr"
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func List(limit, offset int, opt model.Option) (int64, []model.Post, error) {
	var posts []model.Post
	db := instance.DB()
	if v := opt.GetString("keyword"); v != "" {
		db = db.Where("title LIKE ? OR content LIKE ?", "%"+v+"%", "%"+v+"%")
	}
	if v := opt.GetString("tag"); v != "" {
		db = db.Where("tags LIKE ?", "%"+v+"%")
	}
	if v := opt.GetString("pinned_only"); v == "yes" {
		db = db.Where("is_pinned=?", true)
	}
	total := int64(0)
	if err := db.Model(&posts).Count(&total).Error; err != nil {
		instance.Logger().Error("db.Count() fail", zap.String("caller", caller("List", "db.Count()")), zap.Error(err))
		return total, posts, bizerr.ErrDB
	}
	if err := db.Limit(limit).Offset(offset).Order("id desc").Find(&posts).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			instance.Logger().Error("db.Find() fail", zap.String("caller", caller("List", "db.Find()")), zap.Error(err))
			return total, posts, bizerr.ErrDB
		}
	}
	separator := model.GetSite().Separator
	var tmp []string
	for k := range posts {
		tmp = strings.Split(posts[k].Content, separator)
		if len(tmp) > 1 {
			posts[k].Content = posts[k].Summary()
		}
	}
	return total, posts, nil
}

func Remove(id uint) error {
	if err := instance.DB().Where("id=?", id).Delete(model.Post{}).Error; err != nil {
		instance.Logger().Error("db.Find() fail", zap.Error(err))
		return bizerr.ErrDB
	}
	return nil
}

func Detail(id uint) (model.Post, error) {
	post := model.Post{}
	if err := instance.DB().First(&post, id).Error; err != nil {
		instance.Logger().Error("post.Detail() db.Find() fail", zap.String("caller", caller("Remove", "db.Delete()")), zap.Error(err))
		return post, err
	}
	return post, nil
}

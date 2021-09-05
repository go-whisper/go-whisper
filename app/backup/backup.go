package backup

import (
	"errors"
	"fmt"
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"path/filepath"
	"time"
)

func Do() error {
	// 检查最近备份的内容ID
	backup := model.BackupLog{}
	if err := instance.DB().Table(backup.TableName()).Order("id desc").First(&backup).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			instance.Logger().Error("Backup.Do() Db.First(backup) fail:", zap.Error(err))
			return err
		}
	}
	post := model.Post{}
	if err := instance.DB().Order("id desc").First(&post).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			instance.Logger().Error("Backup.Do() Db.First(post) fail:", zap.Error(err))
			return err
		}
	}
	if id, _ := backup.Others.GetUint("post_id"); id == post.ID {
		instance.Logger().Info("backup:no new content,exit")
		return nil
	}

	fName := "db-bak-" + time.Now().Format("20060102150405") + ".db"
	path := filepath.Join(viper.GetString("database.backupPath"), fName)
	fmt.Println("path:", path)
	if err := instance.DB().Raw(".backup ?", path).Error; err != nil {
		instance.Logger().Error("Backup.Do() Db.Exec() fail:", zap.Error(err))
		return err
	}
	//storage.Put("")
	instance.Logger().Info("backup:ok")
	return nil
}

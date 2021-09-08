package backup

import (
	"errors"
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"github.com/go-whisper/go-whisper/app/storage"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

func Do(typ string) error {
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

	rand.Seed(time.Now().UnixNano())
	rnd := 1000 + rand.Intn(8000)
	fName := "backup/db-" + time.Now().Format("20060102150405") + "-" + strconv.Itoa(rnd) + ".db"
	if err := storage.PutFromFile(fName, viper.GetString("database.dsn")); err != nil {
		instance.Logger().Error("backup.fail:", zap.Error(err))
		return err
	}
	log := model.BackupLog{
		Type:      typ,
		CloudKey:  fName,
		Others:    model.InterfaceMap{"post_id": post.ID},
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := instance.DB().Create(&log).Error; err != nil {
		instance.Logger().Error("Backup.Do() Db.Create(&log) fail:", zap.Error(err))
		return err
	}
	instance.Logger().Info("backup.ok")
	return nil
}

package backup

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"github.com/go-whisper/go-whisper/app/storage"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Do(typ string) error {
	// 检查最近备份的内容ID
	lastBackup := model.BackupLog{}
	if err := instance.DB().Table(lastBackup.TableName()).Order("id desc").First(&lastBackup).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			instance.Logger().Error("Backup.Do() Db.First(lastBackup) fail:", zap.Error(err))
			return err
		}
	}
	siteParameter := model.DBSiteParameter{}
	if err := instance.DB().Where("option='last_change'").Order("id desc").First(&siteParameter).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			instance.Logger().Error("Backup.Do() Db.First(siteParameter) fail:", zap.Error(err))
			return err
		}
	}
	lastBackupAt, err := time.Parse("2006-01-02 15:04:05", lastBackup.CreatedAt)
	if err != nil {
		instance.Logger().Error("Backup.Do() time.Parse(lastBackup.CreatedAt) fail:", zap.Error(err))
	}
	lastChange, err := time.Parse("2006-01-02 15:04:05", siteParameter.Value)
	if err != nil {
		instance.Logger().Error("Backup.Do() time.Parse(siteParameter.Value) fail:", zap.Error(err))
	}
	if lastBackupAt.After(lastChange) {
		instance.Logger().Info("BACKUP:no new content,exit")
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	rnd := 1000 + rand.Intn(8000)
	fName := "lastBackup/db-" + time.Now().Format("20060102150405") + "-" + strconv.Itoa(rnd) + ".db"
	if err := storage.PutFromFile(fName, viper.GetString("database.dsn")); err != nil {
		instance.Logger().Error("lastBackup.fail:", zap.Error(err))
		return err
	}
	log := model.BackupLog{
		Type:      typ,
		CloudKey:  fName,
		Others:    model.InterfaceMap{},
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := instance.DB().Create(&log).Error; err != nil {
		instance.Logger().Error("Backup.Do() Db.Create(&log) fail:", zap.Error(err))
		return err
	}
	instance.Logger().Info("BACKUP.OK *_*")
	return nil
}

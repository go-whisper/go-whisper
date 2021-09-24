package model

import (
	"time"

	"go.uber.org/zap"

	"github.com/go-whisper/go-whisper/app/instance"

	"gorm.io/gorm"
)

func UpdateLastChange(tx *gorm.DB) error {
	res := tx.Table("site_parameters").Where("option='last_change'").
		UpdateColumn("value", time.Now().Format("2006-01-02 15:04:05"))
	if res.Error != nil {
		instance.Logger().Error("model.UpdateLastChange() tx.Create() fail:", zap.Error(res.Error))
		return res.Error
	}
	if res.RowsAffected > 0 {
		return nil
	}
	if err := tx.Create(&DBSiteParameter{Option: "last_change", Value: time.Now().Format("2006-01-02 15:04:05")}).
		Error; err != nil {
		instance.Logger().Error("model.UpdateLastChange() tx.Create() fail:", zap.Error(err))
		return err
	}
	return nil
}

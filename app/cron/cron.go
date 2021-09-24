package cron

import (
	"fmt"

	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"github.com/go-whisper/go-whisper/app/service/backup"
	cronV3 "github.com/robfig/cron/v3"
)

func Start() {
	c := cronV3.New()
	c.AddFunc("0 3 * * *", func() {
		if err := backup.Do(model.BackupTypeManual); err != nil {
			fmt.Println("**ERROR** backup fail:", err)
		}
	})
	instance.Logger().Info("cron.Start() done")
	c.Start()
}

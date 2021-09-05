package cron

import (
	"github.com/go-whisper/go-whisper/app/instance"
	cronV3 "github.com/robfig/cron/v3"
)

func Start() {
	c := cronV3.New()
	c.AddFunc("0 3 * * *", func() {
		instance.Logger().Info("start backup")
	})
	instance.Logger().Info("cron.Start() done")
	c.Start()
}

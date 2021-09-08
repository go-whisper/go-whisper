package commands

import (
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/spf13/viper"
)

func Web() {
	service := instance.WebService()
	service.LoadHTMLGlob(viper.GetString("template.path") + "/**/*")
	service.Static("public", "./public")
	service.Run(viper.GetString("service.address"))
}

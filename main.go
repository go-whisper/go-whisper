package main

import (
	"github.com/go-whisper/go-whisper/app/instance"
	_ "github.com/go-whisper/go-whisper/app/router"
	"github.com/spf13/viper"
)

func main() {
	service := instance.WebService()
	service.LoadHTMLGlob(viper.GetString("template.path") + "/**/*")
	service.Static("public", "./public")
	service.Run(viper.GetString("service.address"))
}

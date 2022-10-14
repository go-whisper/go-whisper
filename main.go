package main

import (
	"fmt"
	"os"

	"github.com/go-whisper/go-whisper/app/commands"
	"github.com/go-whisper/go-whisper/app/cron"
	"github.com/go-whisper/go-whisper/app/model"
	_ "github.com/go-whisper/go-whisper/app/router"
	"github.com/go-whisper/go-whisper/app/service/backup"
)

func main() {
	flag := "web"
	args := os.Args
	if len(args) > 1 {
		flag = args[1]
	}
	switch flag {
	case "version":
		// TODO
		fmt.Println("version:todo")
	case "backup":
		if err := backup.Do(model.BackupTypeManual); err != nil {
			fmt.Println("fail:", err)
		}
	case "web":
		cron.Start()
		commands.Web()
	case "install":
		commands.Install()
	}
}

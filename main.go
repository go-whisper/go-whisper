package main

import (
	"fmt"
	"github.com/go-whisper/go-whisper/app/backup"
	"github.com/go-whisper/go-whisper/app/commands"
	"github.com/go-whisper/go-whisper/app/cron"
	"github.com/go-whisper/go-whisper/app/model"
	_ "github.com/go-whisper/go-whisper/app/router"
	"os"
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
	}
}

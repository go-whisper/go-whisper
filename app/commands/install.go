package commands

import (
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/go-whisper/go-whisper/app/model"
	"github.com/go-whisper/go-whisper/app/service/user"
	"go.uber.org/zap"
	"os"
)

const (
	installLocker   = "storage/install.lock"
	defaultUser     = "whisper"
	defaultPassword = "whisper"
)

func Install() {
	if isInstalled() {
		instance.Logger().Info("whisper has been installed.")
		return
	}

	if err := migrate(); err != nil {
		instance.Logger().Error("migrate database error", zap.Error(err))
		return
	}

	if err := createUser(); err != nil {
		instance.Logger().Error("create admin user error", zap.Error(err))
		return
	}

	instance.Logger().Info(
		"create admin user",
		zap.String("name", defaultUser),
		zap.String("password", defaultPassword),
	)

	lockInstall()
}

func isInstalled() bool {
	_, err := os.Stat(installLocker)

	if err == nil {
		return true
	}

	return !os.IsNotExist(err)
}

func migrate() error {
	return instance.DB().AutoMigrate(
		&model.DBSiteParameter{},
		&model.User{},
		&model.Post{},
		&model.BackupLog{},
	)
}

func createUser() error {
	return user.Create(&model.User{
		Name:     defaultUser,
		Password: user.Encrypt(defaultUser),
	})
}

func lockInstall() {
	_, _ = os.Create(installLocker)
}

package instance

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/spf13/viper"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var db *gorm.DB
var onceDB sync.Once

func DB() *gorm.DB {
	dsn := viper.GetString("database.dsn")
	onceDB.Do(func() {
		var err error
		dbLogger := gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			gormLogger.Config{
				SlowThreshold:             time.Second,     // Slow SQL threshold
				LogLevel:                  gormLogger.Info, // Log level
				IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,            // Disable color
			},
		)

		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{Logger: dbLogger})
		if err != nil {
			panic("failed to connect database")
		}
	})
	return db
}

package instance

import (
	"sync"

	"github.com/spf13/viper"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var db *gorm.DB
var onceDB sync.Once

func DB() *gorm.DB {
	dbpath := viper.GetString("database.path")
	onceDB.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	})
	return db
}

package instance

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func init() {
	projectPath := os.Getenv("GO_PROJECT_PATH")
	if projectPath == "" {
		projectPath = "."
	}
	viper.AddConfigPath(projectPath)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

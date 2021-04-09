package instance

import (
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

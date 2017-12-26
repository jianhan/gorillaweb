package bootstrap

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	if err := initializeConfigs(); err != nil {
		panic(fmt.Sprintf("panic : %s", err.Error()))
	}
}

func initializeConfigs() error {
	viper.SetConfigName("app")
	viper.AddConfigPath("configs")
	viper.SetConfigType("yaml")
	viper.WatchConfig()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

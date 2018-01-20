package bootstrap

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Boot() (err error) {
	return initializeConfigs()
}

func initializeConfigs() error {
	// set default values
	viper.SetDefault("appName", "Go App")
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", 8888)
	viper.SetDefault("server.writeTimeout", 15)
	viper.SetDefault("server.readTimeout", 15)
	viper.SetDefault("server.readHeaderTimeout", 15)
	viper.SetDefault("environment", "development")
	viper.SetDefault("enableLog", false)
	// mongo db configs
	viper.SetDefault("mongo.url", "localhost")
	// start initialize loading
	viper.SetConfigName("app")
	viper.AddConfigPath("configs")
	viper.SetConfigType("yaml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

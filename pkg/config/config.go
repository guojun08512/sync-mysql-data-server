package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config is a singleton config manager
var Config = func() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yml")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	config.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	config.SetEnvKeyReplacer(replacer)
	return config
}()

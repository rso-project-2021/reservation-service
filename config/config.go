package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"db_driver"`
	DBSource      string `mapstructure:"db_source"`
	ServerAddress string `mapstructure:"server_address"`
	GinMode       string `mapstructure:"gin_mode"`
}

// Reads configuration from file or environment variables.
func New(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	// Some small change
	err = viper.Unmarshal(&config)

	log.Println(config.GinMode)
	return
}

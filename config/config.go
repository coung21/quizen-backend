package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MySqlUri      string `mapstructure:"MYSQL_URI"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig loads the config from the environment
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

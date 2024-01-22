package config

import (
	"github.com/spf13/viper"
)

// Config is the configuration for the application
type Config struct {
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	MySqlUri            string `mapstructure:"MYSQL_URI"`
	RedisAddress        string `mapstructure:"REDIS_ADDRESS"`
	EmailSenderName     string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`
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

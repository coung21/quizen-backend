package config

import (
	"github.com/spf13/viper"
)

// Config is the configuration for the application
type Config struct {
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	MysqlUser           string `mapstructure:"MYSQL_USER"`
	MysqlPassword       string `mapstructure:"MYSQL_PASSWORD"`
	MysqlDb             string `mapstructure:"MYSQL_DB"`
	MysqlHost           string `mapstructure:"MYSQL_HOST"`
	MysqlPort           string `mapstructure:"MYSQL_PORT"`
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

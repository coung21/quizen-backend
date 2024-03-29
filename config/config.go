package config

import (
	"github.com/spf13/viper"
)

// Config is the configuration for the application
type Config struct {
	ServerAddress        string `mapstructure:"SERVER_ADDRESS"`
	SecretKey            string `mapstructure:"SECRET_KEY"`
	AccessTokenDuration  int    `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration int    `mapstructure:"REFRESH_TOKEN_DURATION"`
	MysqlUser            string `mapstructure:"MYSQL_USER"`
	MysqlPassword        string `mapstructure:"MYSQL_PASSWORD"`
	MysqlDb              string `mapstructure:"MYSQL_DB"`
	MysqlHost            string `mapstructure:"MYSQL_HOST"`
	MysqlPort            string `mapstructure:"MYSQL_PORT"`
	RedisAddress         string `mapstructure:"REDIS_ADDRESS"`
	EmailSenderName      string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string `mapstructure:"EMAIL_SENDER_PASSWORD"`
	S3BucketName         string `mapstructure:"S3_BUCKET_NAME"`
	S3Region             string `mapstructure:"S3_REGION"`
	S3Domain             string `mapstructure:"S3_DOMAIN"`
	S3AccessKey          string `mapstructure:"S3_ACCESS_KEY"`
	S3SecretKey          string `mapstructure:"S3_SECRET_KEY"`
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

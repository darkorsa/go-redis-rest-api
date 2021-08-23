package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	RedisServer         string        `mapstructure:"REDIS_SERVER"`
	RedisPort           string        `mapstructure:"REDIS_PORT"`
	RedisPaswword       string        `mapstructure:"REDIS_PASSWORD"`
	RedisDB             int           `mapstructure:"REDIS_DB"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	AuthUsername        string        `mapstructure:"AUTH_USERNAME"`
	AuthPassword        string        `mapstructure:"AUTH_PASSWORD"`
	AllowedOrigins      []string      `mapstructure:"ALLOWED_ORIGINS"`
}

// LoadConfig reads configuration from file or environment variables.
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

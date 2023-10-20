package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	DevMode  = "DEVELOPMENT"
	TestMode = "TESTING"
	ProdMode = "PRODUCTION"
)

type Config struct {
	AppMode             string        `mapstructure:"APP_MODE"`
	LogLevel            string        `mapstructure:"LOG_LEVEL"`
	DbDriver            string        `mapstructure:"DB_DRIVER"`
	DbSource            string        `mapstructure:"DB_SOURCE"`
	RedisAddress        string        `mapstructure:"REDIS_ADDRESS"`
	MigrationUrl        string        `mapstructure:"MIGRATION_URL"`
	HttpServerHost      string        `mapstructure:"HTTP_SERVER_HOST"`
	HttpServerPort      string        `mapstructure:"HTTP_SERVER_PORT"`
	GrpcServerAddress   string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RedisExpiration     time.Duration `mapstructure:"REDIS_EXPIRATION"`
}

func NewConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return
}

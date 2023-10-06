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
	AppMode              string        `mapstructure:"APP_MODE"`
	LogLevel             string        `mapstructure:"LOG_LEVEL"`
	DbDriver             string        `mapstructure:"DB_DRIVER"`
	DbSource             string        `mapstructure:"DB_SOURCE"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	MigrationUrl         string        `mapstructure:"MIGRATION_URL"`
	HttpServerPort       string        `mapstructure:"HTTP_SERVER_PORT"`
	GrpcServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func NewConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

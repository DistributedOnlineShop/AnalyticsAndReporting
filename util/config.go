package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment               string        `mapstructure:"ENVIRONMENT"`
	DBDriver                  string        `mapstructure:"DB_DRIVER"`
	DBSource                  string        `mapstructure:"DB_SOURCE"`
	AllowHeaders              []string      `mapstructure:"ALLOW_HEADERS"`
	RedisAddress              string        `mapstructure:"REDIS_ADDRESS"`
	RedisPassword             string        `mapstructure:"REDIS_PASSWORD"`
	MigrationURL              string        `mapstructure:"MIGRATION_URL"`
	AnalyticsAndReportingPort string        `mapstructure:"ANALYTICS_AND_REPORTING_PORT"`
	TokenSymmetricKey         string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration       time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
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

package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	PostgresConnStr string `mapstructure:"PostgresConnStr"`
	ServerAdd       string `mapstructure:"ServerAdd"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName("app")

	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return &config, err
	}

	err = viper.Unmarshal(&config)
	return &config, err
}

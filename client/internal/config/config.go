package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	IntervalSec   int    `mapstructure:"INTERVAL_SEC"`
}

func LoadConfig() (config *Config, err error) {
	viper.SetConfigName(".env")
	viper.SetConfigFile("./.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_ADDRESS", "localhost")
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("INTERVAL_SEC", 5)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

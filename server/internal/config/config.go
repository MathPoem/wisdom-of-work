package config

import (
	"fmt"

	"github.com/spf13/viper"
)


type Config struct {
	Port       string `mapstructure:"PORT"`
	Difficulty int    `mapstructure:"DIFFICULTY"`
	POWType    string `mapstructure:"POW_TYPE"`
}

func LoadConfig() (config *Config, err error) {
	viper.SetConfigName(".env")
	viper.SetConfigFile("./.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetDefault("PORT", ":8080")
	viper.SetDefault("DIFFICULTY", 5)
	viper.SetDefault("POW_TYPE", "hashcash")
	
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	if config.POWType != "hashcash" && config.POWType != "quadratic_residue" {
		return nil, fmt.Errorf("invalid type of PoW: %s", config.POWType)
	}

	return config, nil
}

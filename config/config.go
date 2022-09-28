package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Log struct {
		Level string
	}
	HTTP struct {
		Port string
	}
	Database struct {
		URL string
	}
}

func Init() (*Config, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	constant := &Config{}
	err = viper.Unmarshal(constant)
	if err != nil {
		return nil, err
	}
	return constant, nil
}

package config

import (
	"github.com/spf13/viper"
)

func ReadConfig(cfgPath string) (*Config, error) {
	viper.AddConfigPath(cfgPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

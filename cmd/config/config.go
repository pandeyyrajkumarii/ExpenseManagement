package config

import (
	"ExpenseManagement/packages/database"
	"github.com/spf13/viper"
)

type Config struct {
	App      App
	Database database.Config
}

type App struct {
	Env         string
	ServiceName string
	Host        string
}

func NewConfig(env string) (*Config, error) {
	var config *Config
	config, err := LoadConfig(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func LoadConfig(config *Config) (*Config, error) {
	configPath := "config/"
	viper.SetConfigName("default")
	viper.AddConfigPath(configPath)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

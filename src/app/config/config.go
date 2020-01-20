package config

import "github.com/jinzhu/configor"

var Config Configuration

type Configuration struct {
	ServiceName string `env:"SERVICE_NAME"`
	Env         string `env:"ENV"`
	LogLevel    string `env:"LOG_LEVEL"`
	Port        string `env:""`
}

func Load(configPath string) error {
	return configor.Load(&Config, configPath)
}

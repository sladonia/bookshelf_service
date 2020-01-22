package config

import "github.com/jinzhu/configor"

var Config Configuration

type Configuration struct {
	ServiceName  string `env:"SERVICE_NAME"`
	Env          string `env:"ENV"`
	LogLevel     string `env:"LOG_LEVEL"`
	Port         string `env:"PORT"`
	WriteTimeout int    `env:"WRITE_TIMEOUT"`
	ReadTimeout  int    `env:"READ_TIMEOUT"`
	IdleTimeout  int    `env:"IDLE_TIMEOUT"`
}

func Load(configPath string) error {
	return configor.Load(&Config, configPath)
}

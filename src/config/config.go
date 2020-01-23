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
	BookshelfDb  BookshelfDb
}

type BookshelfDb struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	DbName   string `env:"POSTGRES_DB"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
}

func Load(configPath string) error {
	return configor.Load(&Config, configPath)
}

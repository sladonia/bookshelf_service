package config

import "github.com/jinzhu/configor"

var Config Configuration

type Configuration struct {
	ServiceName string `env:"SERVICE_NAME"`
	Env         string `env:"ENV"`
	LogLevel    string `env:"LOG_LEVEL"`
	Port        string `env:"PORT"`
	BookshelfDb BookshelfDb
}

type BookshelfDb struct {
	Host                  string `env:"POSTGRES_HOST"`
	Port                  string `env:"POSTGRES_PORT"`
	DbName                string `env:"POSTGRES_DB"`
	User                  string `env:"POSTGRES_USER"`
	Password              string `env:"POSTGRES_PASSWORD"`
	MaxOpenConnections    int    `env:"POSTGRES_MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections    int    `env:"POSTGRES_MAX_IDLE_CONNECTIONS"`
	ConnectionMaxLifetime int    `env:"POSTGRES_CONNECTION_MAX_LIFETIME"`
}

func Load(configPath string) error {
	return configor.Load(&Config, configPath)
}

package config

import "github.com/jinzhu/configor"

var Config Configuration

type Configuration struct {
	ServiceName string `env:"SERVICE_NAME" default:"bookshelf_service"`
	Env         string `env:"ENV" default:"dev"`
	LogLevel    string `env:"LOG_LEVEL" default:"debug"`
	Port        string `env:"PORT" default:":8080"`
	BookshelfDb BookshelfDb
}

type BookshelfDb struct {
	Host                  string `env:"POSTGRES_HOST" default:"localhost"`
	Port                  string `env:"POSTGRES_PORT" default:"5432"`
	DbName                string `env:"POSTGRES_DB" default:"bookshelf_db"`
	User                  string `env:"POSTGRES_USER" default:"user"`
	Password              string `env:"POSTGRES_PASSWORD" default:"password"`
	MaxOpenConnections    int    `env:"POSTGRES_MAX_OPEN_CONNECTIONS" default:"25"`
	MaxIdleConnections    int    `env:"POSTGRES_MAX_IDLE_CONNECTIONS" default:"25"`
	ConnectionMaxLifetime int    `env:"POSTGRES_CONNECTION_MAX_LIFETIME" default:"5"`
}

func Load() error {
	return configor.Load(&Config)
}

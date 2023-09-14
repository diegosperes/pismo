package util

import (
	"github.com/caarlos0/env/v9"
)

type ServerSettings struct {
	Port         string `env:"SERVER_PORT,required"`
	LogLevelName string `env:"SERVER_LOG_LEVEL_NAME"`
}

type DatabaseSettings struct {
	Host         string `env:"DATABASE_HOST,required"`
	Port         int    `env:"DATABASE_PORT" envDefault:"5432"`
	Name         string `env:"DATABASE_NAME,required"`
	User         string `env:"DATABASE_USER,required"`
	Pass         string `env:"DATABASE_PASS,required"`
	SSLMode      string `env:"DATABASE_SSL_MODE" envDefault:"enable"`
	LogLevelName string `env:"DATABASE_LOG_LEVEL_NAME"`
}

type AppSettings struct {
	Server   ServerSettings
	Database DatabaseSettings
}

func LoadSettings() (*AppSettings, error) {
	settings := &AppSettings{}
	err := env.Parse(settings)
	return settings, err
}

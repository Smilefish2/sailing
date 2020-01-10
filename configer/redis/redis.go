package redis

import (
	"github.com/caarlos0/env/v6"
)

type Redis interface {
	GetHost() string
	GetPort() uint16
	GetDatabase() uint
	GetPassword() string
}

type config struct {
	Host     string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	Port     uint16 `env:"REDIS_PORT" envDefault:"6379"`
	Database uint   `env:"REDIS_DATABASE" envDefault:"0"`
	Password string `env:"REDIS_PASSWORD" envDefault:"" `
}

func NewConfig() Redis {
	var envConfig = config{}
	if err := env.Parse(&envConfig); err != nil {
		panic(err)
	}
	return envConfig
}

func (c config) GetHost() string {
	return c.Host
}

func (c config) GetPort() uint16 {
	return c.Port
}

func (c config) GetDatabase() uint {
	return c.Database
}

func (c config) GetPassword() string {
	return c.Password
}

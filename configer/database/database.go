package database

import (
	"github.com/caarlos0/env/v6"
)

type Database interface {
	GetConnection() string
	GetHost() string
	GetPort() uint16
	GetDatabase() string
	GetUsername() string
	GetPassword() string
	GetCharset() string
	GetCollation() string
}

type config struct {
	Connection string `env:"DB_CONNECTION" envDefault:"mysql"`
	Host       string `env:"DB_HOST" envDefault:"127.0.0.1"`
	Port       uint16 `env:"DB_PORT" envDefault:"3306"`
	Database   string `env:"DB_DATABASE" envDefault:"app"`
	Username   string `env:"DB_USERNAME" envDefault:"root"`
	Password   string `env:"DB_PASSWORD" envDefault:"root"`
	Charset    string `env:"DB_CHARSET" envDefault:"utf8mb4"`
	Collation  string `env:"DB_COLLATION" envDefault:"utf8mb4_unicode_ci"`
}

func NewConfig() Database {
	var envConfig = config{}
	if err := env.Parse(&envConfig); err != nil {
		panic(err)
	}
	return envConfig
}

func (c config) GetConnection() string {
	return c.Connection
}

func (c config) GetHost() string {
	return c.Host
}

func (c config) GetPort() uint16 {
	return c.Port
}

func (c config) GetDatabase() string {
	return c.Database
}

func (c config) GetUsername() string {
	return c.Username
}

func (c config) GetPassword() string {
	return c.Password
}

func (c config) GetCharset() string {
	return c.Charset
}

func (c config) GetCollation() string {
	return c.Collation
}

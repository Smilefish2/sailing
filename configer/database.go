package configer

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

type databaseConfig struct {
	Connection string `env:"DB_CONNECTION" envDefault:"mysql"`
	Host       string `env:"DB_HOST" envDefault:"127.0.0.1"`
	Port       uint16 `env:"DB_PORT" envDefault:"3306"`
	Database   string `env:"DB_DATABASE" envDefault:"test"`
	Username   string `env:"DB_USERNAME" envDefault:"root"`
	Password   string `env:"DB_PASSWORD" envDefault:"root"`
	Charset    string `env:"DB_CHARSET" envDefault:"utf8mb4"`
	Collation  string `env:"DB_COLLATION" envDefault:"utf8mb4_unicode_ci"`
}

var databaseConf *databaseConfig

func NewDatabaseConfig() Database {
	if databaseConf == nil {
		databaseConf = &databaseConfig{}
		once.Do(func() {
			if err := env.Parse(databaseConf); err != nil {
				panic(err)
			}
		})
	}

	return databaseConf
}

func (c databaseConfig) GetConnection() string {
	return c.Connection
}

func (c databaseConfig) GetHost() string {
	return c.Host
}

func (c databaseConfig) GetPort() uint16 {
	return c.Port
}

func (c databaseConfig) GetDatabase() string {
	return c.Database
}

func (c databaseConfig) GetUsername() string {
	return c.Username
}

func (c databaseConfig) GetPassword() string {
	return c.Password
}

func (c databaseConfig) GetCharset() string {
	return c.Charset
}

func (c databaseConfig) GetCollation() string {
	return c.Collation
}

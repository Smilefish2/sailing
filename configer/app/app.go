package app

import (
	"github.com/caarlos0/env/v6"
)

type App interface {
	GetEnv() string
	GetKey() string
	GetUrl() string
	GetName() string
	GetDebug() bool
	GetVersion() string
	GetTimezone() string
}

type config struct {
	Env      string `env:"APP_ENV" envDefault:"local"`
	Key      string `env:"APP_KEY" envDefault:""`
	URL      string `env:"APP_URL" envDefault:"http://localhost"`
	Name     string `env:"APP_NAME" envDefault:"app name"`
	Debug    bool   `env:"APP_DEBUG" envDefault:"true"`
	Version  string `env:"APP_VERSION" envDefault:"latest"`
	Timezone string `env:"APP_TIMEZONE" envDefault:"UTC"`
}

func NewConfig() App {
	var envConfig = config{}
	if err := env.Parse(&envConfig); err != nil {
		panic(err)
	}
	return envConfig
}

func (c config) GetEnv() string {
	return c.Env
}

func (c config) GetKey() string {
	return c.Key
}

func (c config) GetUrl() string {
	return c.URL
}

func (c config) GetName() string {
	return c.Name
}

func (c config) GetDebug() bool {
	return c.Debug
}

func (c config) GetVersion() string {
	return c.Version
}

func (c config) GetTimezone() string {
	return c.Timezone
}

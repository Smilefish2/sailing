package configer

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

type appConfig struct {
	Env      string `env:"APP_ENV" envDefault:"local"`
	Key      string `env:"APP_KEY" envDefault:""`
	URL      string `env:"APP_URL" envDefault:"http://localhost"`
	Name     string `env:"APP_NAME" envDefault:"app name"`
	Debug    bool   `env:"APP_DEBUG" envDefault:"true"`
	Version  string `env:"APP_VERSION" envDefault:"latest"`
	Timezone string `env:"APP_TIMEZONE" envDefault:"UTC"`
}

var appConf *appConfig

func AppConfig() App {
	if appConf == nil {
		appConf = &appConfig{}
		once.Do(func() {
			if err := env.Parse(appConf); err != nil {
				panic(err)
			}
		})
	}
	return appConf
}

func (c appConfig) GetEnv() string {
	return c.Env
}

func (c appConfig) GetKey() string {
	return c.Key
}

func (c appConfig) GetUrl() string {
	return c.URL
}

func (c appConfig) GetName() string {
	return c.Name
}

func (c appConfig) GetDebug() bool {
	return c.Debug
}

func (c appConfig) GetVersion() string {
	return c.Version
}

func (c appConfig) GetTimezone() string {
	return c.Timezone
}

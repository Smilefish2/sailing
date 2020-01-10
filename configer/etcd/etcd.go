package etcd

import (
	"github.com/caarlos0/env/v6"
)

type Etcd interface {
	GetHost() string
	GetPort() uint16
	GetEnabled() bool
}

type config struct {
	Host    string `env:"ETCD_HOST" envDefault:"127.0.0.1"`
	Port    uint16 `env:"ETCD_PORT" envDefault:"2379"`
	Enabled bool   `env:"ETCD_Enabled" envDefault:"false"`
}

func NewConfig() Etcd {
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

func (c config) GetEnabled() bool {
	return c.Enabled
}

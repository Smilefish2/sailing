package configer

import (
	"github.com/caarlos0/env/v6"
)

type Etcd interface {
	GetHost() string
	GetPort() uint16
	GetEnabled() bool
}

type etcdConfig struct {
	Host    string `env:"ETCD_HOST" envDefault:"127.0.0.1"`
	Port    uint16 `env:"ETCD_PORT" envDefault:"2379"`
	Enabled bool   `env:"ETCD_Enabled" envDefault:"false"`
}

func NewEtcdConfig() Etcd {
	var envConfig = etcdConfig{}
	if err := env.Parse(&envConfig); err != nil {
		panic(err)
	}
	return envConfig
}

func (c etcdConfig) GetHost() string {
	return c.Host
}

func (c etcdConfig) GetPort() uint16 {
	return c.Port
}

func (c etcdConfig) GetEnabled() bool {
	return c.Enabled
}

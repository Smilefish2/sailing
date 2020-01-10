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

var etcdConf *etcdConfig

func NewEtcdConfig() Etcd {
	if etcdConf == nil {
		etcdConf = &etcdConfig{}
		once.Do(func() {
			if err := env.Parse(etcdConf); err != nil {
				panic(err)
			}
		})
	}
	return etcdConf
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

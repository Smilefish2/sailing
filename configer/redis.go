package configer

import "github.com/caarlos0/env/v6"

type Redis interface {
	GetHost() string
	GetPort() uint16
	GetDatabase() uint
	GetPassword() string
}

type redisConfig struct {
	Host     string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	Port     uint16 `env:"REDIS_PORT" envDefault:"6379"`
	Database uint   `env:"REDIS_DATABASE" envDefault:"0"`
	Password string `env:"REDIS_PASSWORD" envDefault:"" `
}

var redisConf *redisConfig

func RedisConfig() Redis {
	if redisConf == nil {
		redisConf = &redisConfig{}
		once.Do(func() {
			if err := env.Parse(redisConf); err != nil {
				panic(err)
			}
		})
	}
	return redisConf
}

func (c redisConfig) GetHost() string {
	return c.Host
}

func (c redisConfig) GetPort() uint16 {
	return c.Port
}

func (c redisConfig) GetDatabase() uint {
	return c.Database
}

func (c redisConfig) GetPassword() string {
	return c.Password
}

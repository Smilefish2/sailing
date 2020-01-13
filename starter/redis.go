package starter

import (
	"configer"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/micro/go-micro/util/log"
	"sync"
)

// DB Global DB connection
var redisClient *redis.Client
var redisOnce sync.Once

func RedisClient() *redis.Client {
	if redisClient == nil {
		redisOnce.Do(func() {
			redisConfig := configer.RedisConfig()
			redisClient = redis.NewClient(&redis.Options{
				Addr:     fmt.Sprintf("%v:%v", redisConfig.GetHost(), redisConfig.GetPort()),
				Password: redisConfig.GetPassword(),      // no password set
				DB:       int(redisConfig.GetDatabase()), // use default DB
			})

			log.Logf("[Init] 服务启动：Redis\n")
		})
	}
	return redisClient
}

func RedisPing() {

	pong, err := RedisClient().Ping().Result()
	if err != nil {
		panic(err)
	}
	log.Logf("[Init] Redis：%s\n", pong)
}

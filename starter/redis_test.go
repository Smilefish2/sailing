package starter_test

import (
	"starter"
	"testing"
)

func TestRedis(t *testing.T) {
	//redis := starter.RedisClient()

	starter.RedisPing()

}

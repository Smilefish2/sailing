package starter_test

import (
	"github.com/Smilefish2/sailing/starter"
	"testing"
)

func TestRedis(t *testing.T) {
	//redis := starter.RedisClient()

	starter.RedisPing()

}

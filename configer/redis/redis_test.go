package redis_test

import (
	"configer/redis"
	"fmt"
	"testing"
)

func TestApp(t *testing.T) {
	f := redis.NewConfig()
	fmt.Printf("Host: %v \n", f.GetHost())
	fmt.Printf("Port: %v \n", f.GetPort())
	fmt.Printf("Database: %v \n", f.GetDatabase())
	fmt.Printf("Password: %v \n", f.GetPassword())
}

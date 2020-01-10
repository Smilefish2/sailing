package configer_test

import (
	"configer"
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	f := configer.RedisConfig()
	fmt.Printf("Host: %v \n", f.GetHost())
	fmt.Printf("Port: %v \n", f.GetPort())
	fmt.Printf("Database: %v \n", f.GetDatabase())
	fmt.Printf("Password: %v \n", f.GetPassword())
}

package etcd_test

import (
	"configer/etcd"
	"fmt"
	"testing"
)

func TestEtcd(t *testing.T) {
	f := etcd.NewConfig()
	fmt.Printf("Host: %v \n", f.GetHost())
	fmt.Printf("Port: %v \n", f.GetPort())
	fmt.Printf("Enabled: %v \n", f.GetEnabled())
}

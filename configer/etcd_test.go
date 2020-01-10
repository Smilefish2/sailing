package configer_test

import (
	"configer"
	"fmt"
	"testing"
)

func TestEtcd(t *testing.T) {
	f := configer.EtcdConfig()
	fmt.Printf("Host: %v \n", f.GetHost())
	fmt.Printf("Port: %v \n", f.GetPort())
	fmt.Printf("Enabled: %v \n", f.GetEnabled())
}

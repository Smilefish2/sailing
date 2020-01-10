package configer_test

import (
	"configer"
	"fmt"
	"testing"
)

func TestApp(t *testing.T) {
	f := configer.AppConfig()
	fmt.Printf("Env: %v \n", f.GetEnv())
	fmt.Printf("Key: %v \n", f.GetKey())
	fmt.Printf("URL: %v \n", f.GetUrl())
	fmt.Printf("Name: %v \n", f.GetName())
	fmt.Printf("Debug: %v \n", f.GetDebug())
	fmt.Printf("Version: %v \n", f.GetVersion())
	fmt.Printf("Timezone: %v \n", f.GetTimezone())
}

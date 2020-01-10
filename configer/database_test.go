package configer_test

import (
	"configer"
	"fmt"
	"testing"
)

func TestDatabase(t *testing.T) {
	f := configer.DatabaseConfig()
	fmt.Printf("Connection: %v \n", f.GetConnection())
	fmt.Printf("Host: %v \n", f.GetHost())
	fmt.Printf("Port: %v \n", f.GetPort())
	fmt.Printf("Database: %v \n", f.GetDatabase())
	fmt.Printf("Username: %v \n", f.GetUsername())
	fmt.Printf("Password: %v \n", f.GetPassword())
	fmt.Printf("Charset: %v \n", f.GetCharset())
	fmt.Printf("Collation: %v \n", f.GetCollation())
}

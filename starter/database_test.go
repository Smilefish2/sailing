package starter_test

import (
	"starter"
	"testing"
)

func TestDatabase(t *testing.T) {

	db := starter.DB().Raw("select * from user")

	db.GetErrors()
}

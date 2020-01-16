package starter_test

import (
	"github.com/Smilefish0/sailing/starter"
	"testing"
)

func TestDatabase(t *testing.T) {

	db := starter.DB().Raw("select * from user")

	db.GetErrors()
}

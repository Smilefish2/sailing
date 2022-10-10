package starter_test

import (
	"github.com/Smilefish2/sailing/starter"
	"testing"
)

func TestDatabase(t *testing.T) {

	db := starter.DB().Raw("select * from user")

	db.GetErrors()
}

package model

import (
	"time"
)

type User struct {
	Model
	Email    string `form:"email"`
	Password string
	Username string `form:"name"`
	Gender   string
	Role     string
	Birthday *time.Time
	Balance  float32
}

func (user User) DisplayName() string {
	return user.Email
}

func (user User) AvailableLocales() []string {
	return []string{"en-US", "zh-CN"}
}

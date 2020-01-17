package user

import (
	"github.com/gf-third/mysql"
	"time"
)

// User represents a row from 'sailing_user.user'.
type User struct {
	ID        uint64         `json:"id"`         // id
	UID       string         `json:"uid"`        // uid
	Type      Type           `json:"type"`       // type
	Email     string         `json:"email"`      // email
	Mobile    string         `json:"mobile"`     // mobile
	Avatar    string         `json:"avatar"`     // avatar
	Nickname  string         `json:"nickname"`   // nickname
	Username  string         `json:"username"`   // username
	Password  string         `json:"password"`   // password
	Salt      string         `json:"salt"`       // salt
	Active    bool           `json:"active"`     // active
	Status    Status         `json:"status"`     // status
	CreatedAt time.Time      `json:"created_at"` // created_at
	UpdatedAt time.Time      `json:"updated_at"` // updated_at
	DeletedAt mysql.NullTime `json:"deleted_at"` // deleted_at
}

package model

import (
	"time"
)

type User struct {
	ID string `json:"user_id" db:"user_id"`

	Username string `json:"username" db:"username"`
	Balance  uint64 `json:"balance" db:"-"`
	Token    string `json:"access_token,omitempty" db:"-"`

	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

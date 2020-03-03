package model

import (
	"time"
)

type User struct {
	ID       string `json:"id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"-" db:"password"`

	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`

	Balance float64 `json:"balance" db:"-"`
}

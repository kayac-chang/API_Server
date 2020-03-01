package model

import (
	"time"
)

type User struct {
	ID       string `json:"id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`

	CreateAt time.Time `json:"created_at" db:"created_at"`
	UpdateAt time.Time `json:"updated_at" db:"updated_at"`
}

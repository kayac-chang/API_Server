package model

import (
	"time"
)

type User struct {
	ID       string `json:"id" db:"user_id"`
	Username string `json:"username" db:"username"`

	Balance Balance `json:"balance" db:"-"`
	Token   string  `json:"-" db:"access_token"`

	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

type Balance float64

type Token struct {
	AccessToken string `json:"access_token"`

	CreatedAt time.Time `json:"issued_at"`
}

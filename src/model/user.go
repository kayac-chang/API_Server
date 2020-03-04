package model

import (
	"time"
)

type User struct {
	ID       string `json:"id" db:"user_id"`
	Username string `json:"username" db:"username"`

	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`

	Balance Balance `json:"balance" db:"-"`
	Token   Token   `json:"token" db:"-"`
}

type Balance float64

type Token struct {
	ServiceID   string `json:"service_id"`
	AccessToken string `json:"access_token"`
	Type        string `json:"token_type"`
	Expire      int64  `json:"expires_in"`
}

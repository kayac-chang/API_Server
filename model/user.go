package model

import "time"

// User ...
type User struct {
	ID       string  `json:"user_id" db:"user_id"`
	Username string  `json:"username" db:"username"`
	Balance  float64 `json:"balance" db:"balance"`
	Session  string  `json:"session" db:"session"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

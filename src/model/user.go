package model

import "time"

type User struct {
	ID       string  `json:"user_id"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
	Session  string  `json:"session"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

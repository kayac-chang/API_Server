package model

import (
	"time"
)

type State string

const (
	Pending   State = "P"
	Fulfilled       = "F"
	Rejected        = "R"
)

type Order struct {
	ID string `json:"id" db:"order_id"`

	State State   `json:"state" db:"state"`
	Bet   float64 `json:"bet" db:"bet"`

	GameID int    `json:"game_id" db:"game_id"`
	UserID string `json:"user_id" db:"user_id"`

	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	CompletedAt time.Time `json:"completed_at" db:"completed_at"`
}

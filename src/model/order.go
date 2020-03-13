package model

import "time"

type State string

const (
	Pending   State = "P"
	Completed       = "C"
	Rejected        = "R"
)

type Order struct {
	ID string `json:"order_id" db:"order_id"`

	GameID string `json:"game_id" db:"game_id"`
	UserID string `json:"user_id" db:"user_id"`

	State State  `json:"state" db:"state"`
	Bet   uint64 `json:"bet" db:"bet"`

	CreatedAt   *time.Time `json:"created_at" db:"created_at"`
	CompletedAt *time.Time `json:"completed_at" db:"completed_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}

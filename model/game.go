package model

import (
	"time"
)

// Game ...
type Game struct {
	ID string `json:"game_id" db:"game_id"`

	Name     string `json:"name" db:"name"`
	Href     string `json:"href" db:"href"`
	Category string `json:"category" db:"category"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

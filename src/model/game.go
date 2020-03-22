package model

import (
	"time"
)

type Game struct {
	ID string `json:"game_id,omitempty" db:"game_id"`

	Name     string `json:"name" db:"name"`
	Href     string `json:"href" db:"href"`
	Category string `json:"category" db:"category"`

	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

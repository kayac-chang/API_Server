package model

import (
	"time"
)

type Game struct {
	ID string `json:"game_id"`

	Name     string `json:"name"`
	Href     string `json:"href"`
	Category string `json:"category"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

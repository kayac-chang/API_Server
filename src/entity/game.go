package entity

import (
	"time"
)

// Game Represent Game Data
type Game struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Href     string    `json:"href"`
	UpdateAt time.Time `json:"updated_at"`
	CreateAt time.Time `json:"created_at"`
}

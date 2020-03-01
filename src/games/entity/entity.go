package entity

import (
	"time"
)

type Game struct {
	ID       string    `json:"id" db:"game_id"`
	Name     string    `json:"name" db:"name"`
	Href     string    `json:"href" db:"href"`
	Category string    `json:"category" db:"category"`
	CreateAt time.Time `json:"created_at" db:"created_at"`
	UpdateAt time.Time `json:"updated_at" db:"updated_at"`
}

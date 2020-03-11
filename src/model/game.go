package model

import (
	"net/url"
	"time"
)

type Game struct {
	ID string `json:"game_id" db:"game_id"`

	Name     string  `json:"name" db:"name"`
	Href     url.URL `json:"href" db:"href"`
	Category string  `json:"category" db:"category"`

	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

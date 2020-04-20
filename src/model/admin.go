package model

import (
	"time"
)

type Admin struct {
	ID string `json:"admin_id"`

	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

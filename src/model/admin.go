package model

import (
	"time"
)

type Admin struct {
	ID string `json:"admin_id" db:"admin_id"`

	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

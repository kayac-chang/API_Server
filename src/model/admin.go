package model

import (
	"time"
)

type Admin struct {
	ID string `json:"admin_id" db:"admin_id"`

	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"username"`
	Password string `json:"-" db:"password"`
	Org      string `json:"organization" db:"organization"`

	Token string `json:"-" db:"-"`

	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

package model

import "time"

type Token struct {
	AccessToken string    `json:"access_token"`
	Type        string    `json:"token_type"`
	ServiceID   string    `json:"service_id"`
	CreatedAt   time.Time `json:"issued_at"`
}

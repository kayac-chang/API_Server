package model

type User struct {
	Data

	ID       string  `json:"user_id"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
	Session  string  `json:"session"`
}

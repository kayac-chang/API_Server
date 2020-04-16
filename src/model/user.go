package model

type User struct {
	Data

	Username string `json:"username"`
	Balance  uint64 `json:"balance"`
}

package model

import "errors"

type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

var (
	ErrUserNotFound = errors.New("User Not Found")

	ErrDBConstraint = errors.New("Integrity Constraint Violation")
)

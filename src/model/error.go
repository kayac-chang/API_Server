package model

import "errors"

type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

var (
	ErrUserNotFound = errors.New("User Not Found")
	ErrUnauthorized = errors.New("The request has not been applied because it lacks valid authentication credentials for the target resource")
	ErrDBConstraint = errors.New("Integrity Constraint Violation")
)

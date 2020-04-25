package model

import "errors"

type Error struct {
	Code    int
	Name    string
	Message string
}

func (it *Error) Error() string {
	return it.Message
}

var (
	ErrNotFound        = errors.New("Resource Not Found")
	ErrUserNotFound    = errors.New("User Not Found")
	ErrGameNotFound    = errors.New("Game Not Found")
	ErrOrderNotFound   = errors.New("Order Not Found")
	ErrExisted         = errors.New("Resource already existed")
	ErrUnauthorized    = errors.New("The request has not been applied because it lacks valid authentication credentials for the target resource")
	ErrDBConstraint    = errors.New("Integrity Constraint Violation")
	ErrUnexpectPayload = errors.New("Unexpected Request Payload")
)

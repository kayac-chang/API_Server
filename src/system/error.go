package system

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrGenTokenError       = errors.New("Token generation failed")

	ErrNotFound = errors.New("Requested Item not found")

	ErrConflict = errors.New("Item already exist")

	ErrBadParamInput = errors.New("Given Param is not valid")

	/*
		401 Unauthorized
		the request has not been applied
		because it lacks valid authentication credentials for the target resource.
	*/
	ErrUnauthorized = errors.New("Request requires user authentication")
	ErrAuthFailure  = errors.New("Authentication Failure")
)

func GetStatusCode(err error) int {

	logrus.Error(err)

	if err, ok := err.(pgx.PgError); ok {

		if err.Code == "23505" {

			return http.StatusConflict
		}
	}

	switch err {

	case ErrGenTokenError, ErrInternalServerError:
		return http.StatusInternalServerError

	case ErrNotFound:
		return http.StatusNotFound

	case ErrConflict:
		return http.StatusConflict

	case ErrUnauthorized, ErrAuthFailure:
		return http.StatusUnauthorized

	default:
		return http.StatusInternalServerError
	}
}

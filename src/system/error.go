package system

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")

	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Requested Item not found")

	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Item already exist")

	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
)

func GetStatusCode(err error) int {

	logrus.Error(err)

	if err, ok := err.(pgx.PgError); ok {

		if err.Code == "23505" {

			return http.StatusConflict
		}
	}

	switch err {

	case ErrInternalServerError:
		return http.StatusInternalServerError

	case ErrNotFound:
		return http.StatusNotFound

	case ErrConflict:
		return http.StatusConflict

	default:
		return http.StatusInternalServerError
	}
}

package token

import (
	"api/model"
	"net/http"
)

func (it Usecase) CheckHeader(session, contentType string) error {

	if err := checkSession(session); err != nil {
		return err
	}

	if err := checkContentType(contentType); err != nil {
		return err
	}

	return nil
}

func checkSession(session string) *model.Error {

	if session == "" {
		return &model.Error{

			Code:    http.StatusBadRequest,
			Message: "Request header must contain session",
		}
	}

	return nil
}

func checkContentType(contentType string) *model.Error {

	compare := "application/json"

	if contentType != compare {

		return &model.Error{
			Code:    http.StatusBadRequest,
			Message: "Content-Type must be " + compare,
		}
	}

	return nil
}

func (it Usecase) CheckPayload(game, username string) error {

	if game == "" || username == "" {

		return &model.Error{
			Code:    http.StatusBadRequest,
			Message: "Request payload must contain <game> and <username>",
		}
	}

	return nil
}

package token

import (
	"api/model"
	"net/http"
)

func (it *Usecase) CheckHeader(session, contentType string) error {

	var err *model.Error

	err = checkSession(session)
	err = checkContentType(contentType)

	if err != nil {
		err.Code = http.StatusBadRequest
		err.Name = "Check Request #1"
	}

	return err
}

func (it *Usecase) CheckPayload(game, username string) error {

	if game == "" || username == "" {
		return &model.Error{
			Code:    http.StatusBadRequest,
			Name:    "Check Request Payload #3",
			Message: "Request payload must contain <game> and <username>",
		}
	}

	return nil
}

func checkSession(session string) *model.Error {

	if session == "" {
		return &model.Error{
			Message: "Request header must contain session",
		}
	}

	return nil
}

func checkContentType(contentType string) *model.Error {
	compare := "application/json"

	if contentType != compare {
		return &model.Error{
			Message: "Content-Type must be " + compare,
		}
	}

	return nil
}

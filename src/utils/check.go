package utils

import (
	"api/model"
	"fmt"
	"net/http"

	"github.com/badoux/checkmail"
)

func CheckMail(email string) error {

	err := checkmail.ValidateFormat(email)

	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {

		return fmt.Errorf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
	}

	return nil
}

func CheckSession(session string) *model.Error {

	if session == "" {
		return &model.Error{

			Code:    http.StatusBadRequest,
			Message: "Request header must contain session",
		}
	}

	return nil
}

func CheckContentType(contentType string, compare string) *model.Error {

	if contentType != compare {

		return &model.Error{
			Code:    http.StatusBadRequest,
			Message: "Content-Type must be " + compare,
		}
	}

	return nil
}

func CheckPayload(json map[string]interface{}, keys ...string) error {

	for _, key := range keys {

		if _, exist := json[key]; !exist {

			return &model.Error{
				Code:    http.StatusBadRequest,
				Message: "Request payload must contain " + key,
			}
		}
	}

	return nil
}

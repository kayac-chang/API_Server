package admin

import (
	"api/model"
	"api/model/response"
	"api/utils"

	"net/http"
)

// POST create admin account
func (it Handler) POST(w http.ResponseWriter, r *http.Request) {

	main := func() interface{} {

		// == Check Content-Type #1 ==
		contentType := r.Header.Get("Content-Type")
		if err := utils.CheckContentType(contentType, "application/json"); err != nil {

			err.Name = "Check Request #1"

			return err
		}

		// == Parse JSON #2 ==
		req, err := utils.ParseJSON(r.Body)
		if err != nil {
			err := err.(*model.Error)

			err.Name = "Parse JSON #2"

			return err
		}

		// == Check Request Payload #3 ==
		if err := it.checkPayload(req); err != nil {
			err := err.(*model.Error)

			err.Name = "Check Request Payload #3"

			return err
		}

		// == Create Admin Account #4 ==
		email := req["email"].(string)
		username := req["username"].(string)
		password := req["password"].(string)

		admin, err := it.usecase.Store(email, username, password)
		if err != nil {
			err := err.(*model.Error)

			err.Name = "Create Admin Account #4"

			return err
		}

		// == Send Response ==
		return response.JSON{
			Code: http.StatusCreated,

			Data: admin,
		}
	}

	it.Send(w, main())
}

func (it Handler) checkPayload(req map[string]interface{}) error {

	err := utils.CheckPayload(req, "secret", "email", "username", "password")
	if err != nil {
		return err
	}

	if req["secret"] != string(it.env.Secret) {

		return &model.Error{
			Code:    http.StatusBadRequest,
			Message: "Secret Code is wrong",
		}
	}

	email := req["email"].(string)
	if err := utils.CheckMail(email); err != nil {
		return err
	}

	password := req["password"].(string)
	if err := utils.CheckPassword(password); err != nil {
		return err
	}

	return nil
}

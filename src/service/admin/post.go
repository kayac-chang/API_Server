package admin

import (
	"api/model"
	"api/model/response"
	"api/utils"

	"net/http"
)

func (it Handler) POST(w http.ResponseWriter, r *http.Request) {

	main := func() interface{} {

		// == Check Content-Type #1 ==
		contentType := r.Header.Get("Content-Type")
		if err := utils.CheckContentType(contentType); err != nil {

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
		admin, err := it.usecase.Store(req["email"], req["username"], req["password"])
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

func (it Handler) checkPayload(req map[string]string) error {

	keys := []string{"secret", "email", "username", "password"}

	for _, key := range keys {

		if val, exist := req[key]; !exist || val == "" {

			return &model.Error{
				Code:    http.StatusBadRequest,
				Message: "Request payload must contain " + key,
			}
		}
	}

	if req["secret"] != string(it.env.Secret) {

		return &model.Error{
			Code:    http.StatusBadRequest,
			Message: "Secret Code is wrong",
		}
	}

	if err := utils.CheckMail(req["email"]); err != nil {
		return err
	}

	if err := utils.CheckPassword(req["password"]); err != nil {
		return err
	}

	return nil
}

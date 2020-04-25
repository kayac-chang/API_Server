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
		if r.Header.Get("Content-Type") != "application/json" {

			return &model.Error{
				Code:    http.StatusBadRequest,
				Name:    "Check Content-Type #1",
				Message: "Content-Type must be application/json",
			}
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

		email := req["email"].(string)
		username := req["username"].(string)
		password := req["password"].(string)

		// == Check Admin Email Already Exist #4 ==
		found, err := it.usecase.Find(email)
		if found != nil {

			return &model.Error{
				Name:    "Check Admin Email #4",
				Code:    http.StatusConflict,
				Message: "Email has been used",
			}
		}
		if err != model.ErrNotFound {

			return &model.Error{
				Name:    "Check Admin Email #4",
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}

		// == Create Admin Account #5 ==
		admin, err := it.usecase.Store(email, username, password)
		if err != nil {
			err := err.(*model.Error)

			err.Name = "Create Admin Account #5"

			return err
		}

		// == Send Response ==
		return response.JSON{
			Code: http.StatusCreated,

			Data: map[string]string{
				"admin_id":   admin.ID,
				"email":      admin.Email,
				"username":   admin.Username,
				"created_at": admin.CreatedAt.String(),
				"updated_at": admin.UpdatedAt.String(),
			},
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

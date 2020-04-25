package admin

import (
	"api/framework/jwt"
	"api/model"
	"api/model/response"
	"api/utils"
	"net/http"
)

// Auth POST /admins/tokens
func (it Handler) Auth(w http.ResponseWriter, r *http.Request) {

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

		email := req["email"].(string)
		password := req["password"].(string)

		// == Check Admin Exist #3 ==
		admin, err := it.usecase.Find(email)
		if err != nil {

			if err != model.ErrNotFound {

				return &model.Error{
					Code:    http.StatusInternalServerError,
					Name:    "Check Admin Exist #3",
					Message: err.Error(),
				}
			}

			return &model.Error{
				Code:    http.StatusUnauthorized,
				Name:    "Check Admin Exist #3",
				Message: "Incorrect email address and / or password",
			}
		}

		// == Check Password #4 ==
		if err := utils.CompareHash(admin.Password, password); err != nil {

			return &model.Error{
				Code:    http.StatusUnauthorized,
				Name:    "Check Password #4",
				Message: "Incorrect email address and / or password",
			}
		}

		// == Generate JWT Token #5 ==
		token, err := jwt.Sign(it.env)
		if err != nil {

			return &model.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Generate JWT Token #5",
				Message: err.Error(),
			}
		}

		// == Associate Token With Admin #6 ==
		if err := it.usecase.Associate(token, admin); err != nil {

			return &model.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Associate Token With Admin #6",
				Message: err.Error(),
			}
		}

		// == Send Response ==
		return response.JSON{
			Code: http.StatusCreated,

			Data: token,
		}
	}

	it.Send(w, main())
}

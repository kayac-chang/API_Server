package admin

import (
	"api/model"
	"api/model/response"
	"encoding/json"
	"net/http"
)

func (it *Handler) Auth(w http.ResponseWriter, r *http.Request) {

	// == Parse Payload ==
	req := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusBadRequest,

			Error: model.Error{
				Name:    "Unexpect Payload",
				Message: model.ErrUnexpectPayload.Error(),
			},
		})

		return
	}

	// Authentication
	token, err := it.usecase.CheckUser(req)
	if err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusUnauthorized,

			Error: model.Error{
				Name:    "Unauthorized",
				Message: model.ErrUnexpectPayload.Error(),
			},
		})

		return
	}

	// == Send Response ==
	it.Send(w, response.JSON{
		Code: http.StatusCreated,

		Data: token,
	})
}

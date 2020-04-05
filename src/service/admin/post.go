package admin

import (
	"api/model"
	"api/model/response"

	"encoding/json"
	"net/http"
)

func (it *Handler) POST(w http.ResponseWriter, r *http.Request) {

	// == Parse Payload ==
	req := map[string]string{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusBadRequest,

			Error: response.Error{
				Name:    "Unexpect Payload",
				Message: model.ErrUnexpectPayload.Error(),
			},
		})

		return
	}

	// Check secret code
	if req["secret"] != string(it.env.Secret) {

		it.Send(w, response.JSON{
			Code: http.StatusUnauthorized,

			Error: response.Error{
				Name:    "Unexpect Secret Code",
				Message: model.ErrUnauthorized.Error(),
			},
		})

		return
	}

	// == Store ==
	admin, err := it.usecase.Store(req)
	if err != nil {

		it.Send(w, response.JSON{
			Code: http.StatusNotAcceptable,

			Error: response.Error{
				Name:    "Create Admin Failed",
				Message: err.Error(),
			},
		})

		return
	}

	// == Send Response ==
	it.Send(w, response.JSON{
		Code: http.StatusCreated,

		Data: admin,
	})
}

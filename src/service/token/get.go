package token

import (
	"api/model/pb"
	"api/model/response"

	"net/http"
)

func (it *Handler) Get(w http.ResponseWriter, r *http.Request) {

	// Get Authorization in http header
	token := it.URLParam(r, "token")

	// Pass in Auth logic
	user, err := it.token.Auth(token)
	if err != nil {
		it.Send(w, response.ProtoBuf{
			Code: http.StatusUnauthorized,

			Data: &pb.Error{
				Code:    http.StatusUnauthorized,
				Name:    "Given token not found in service",
				Message: err.Error(),
			},
		})

		return

	}

	// == Send Response ==
	it.Send(w, response.ProtoBuf{

		Code: http.StatusOK,

		Data: &pb.User{
			UserId:   user.ID,
			Username: user.Username,
			Balance:  user.Balance,
		},
	})
}

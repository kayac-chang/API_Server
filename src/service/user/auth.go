package user

import (
	"api/model"
	"api/model/pb"
	"api/model/response"

	"net/http"
	"strings"
)

func (it *Handler) Auth(w http.ResponseWriter, r *http.Request) {

	// Get Authorization in http header
	authStr := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(authStr) != 2 {

		it.Send(w, response.ProtoBuf{
			Code: http.StatusUnauthorized,

			Data: &pb.Error{
				Code:    http.StatusUnauthorized,
				Name:    "Please provide correct Authorization header",
				Message: model.ErrUnauthorized.Error(),
			},
		})

		return
	}

	// Pass in Auth logic
	user, err := it.userCase.Auth(authStr[1])
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

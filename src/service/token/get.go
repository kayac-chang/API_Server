package token

import (
	"api/model"
	"api/model/pb"
	"api/model/response"
	"net/http"
)

// GET ..
func (it *Handler) GET(w http.ResponseWriter, r *http.Request) {

	main := func() interface{} {

		// == Check Token Exist #1 ==
		token := it.URLParam(r, "token")
		id, err := it.usecase.Auth(token)
		if err != nil {

			code := http.StatusInternalServerError
			if err == model.ErrNotFound {
				code = http.StatusUnauthorized
			}

			return response.ProtoBuf{
				Code: code,

				Data: &pb.Error{
					Code:    uint32(code),
					Name:    "Check Token Exist #1",
					Message: err.Error(),
				},
			}
		}

		// == Find User #2 ==
		user, err := it.usecase.FindUserByID(id)
		if err != nil {

			code := http.StatusInternalServerError
			if err == model.ErrNotFound {
				code = http.StatusUnauthorized
			}

			return response.ProtoBuf{
				Code: code,

				Data: &pb.Error{
					Code:    uint32(code),
					Name:    "Find User #2",
					Message: err.Error(),
				},
			}
		}

		// == Send Response ==
		return response.ProtoBuf{
			Code: http.StatusOK,

			Data: &pb.User{
				UserId:   user.ID,
				Username: user.Username,
				Balance:  uint64(user.Balance),
			},
		}
	}

	it.Send(w, main())
}

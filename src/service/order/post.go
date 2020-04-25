package order

import (
	"api/model"
	"api/model/pb"
	"api/model/response"
	"api/utils"
	"net/http"
)

// POST ...
func (it Handler) POST(w http.ResponseWriter, r *http.Request) {

	main := func() interface{} {

		// == Check Authorization #1 ==
		token := r.Header.Get("Authorization")
		if err := it.usecase.Auth(token); err != nil {

			return &model.Error{
				Code:    http.StatusUnauthorized,
				Name:    "Check Authorization #1",
				Message: err.Error(),
			}
		}

		// == Check Content-Type #2 ==
		contentType := r.Header.Get("Content-Type")
		if err := utils.CheckContentType(contentType, "application/protobuf"); err != nil {

			return response.ProtoBuf{
				Code: http.StatusBadRequest,

				Data: &pb.Error{
					Code:    http.StatusBadRequest,
					Name:    "Check Content-Type #2",
					Message: err.Error(),
				},
			}
		}

		// == Parse ProtoBuf #3 ==
		order, err := it.Parse(r.Body)
		if err != nil {

			return response.ProtoBuf{
				Code: http.StatusBadRequest,

				Data: &pb.Error{
					Code:    http.StatusBadRequest,
					Name:    "Parse ProtoBuf #3",
					Message: err.Error(),
				},
			}
		}

		// == Check Game Exist #4 ==
		game, err := it.usecase.FindGameByID(order.GameID)
		if err != nil {

			code := http.StatusNotFound
			if err != model.ErrNotFound {
				code = http.StatusInternalServerError
			}

			return response.ProtoBuf{
				Code: code,

				Data: &pb.Error{
					Code:    uint32(code),
					Name:    "Check Game Exist #4",
					Message: err.Error(),
				},
			}
		}

		// == Check User Exist #5 ==
		user, err := it.usecase.FindUserByID(order.UserID)
		if err != nil {

			code := http.StatusNotFound
			msg := "User does not existed"

			if err != model.ErrNotFound {
				code = http.StatusInternalServerError
				msg = err.Error()
			}

			return response.ProtoBuf{
				Code: code,

				Data: &pb.Error{
					Code:    uint32(code),
					Name:    "Check User Exist #5",
					Message: msg,
				},
			}
		}

		// == Send Bet #6 ==
		balance, err := it.usecase.SendBet(user, game, order)
		if err != nil {
			_err := err.(*model.Error)

			return response.ProtoBuf{
				Code: _err.Code,

				Data: &pb.Error{
					Code:    uint32(_err.Code),
					Name:    "Send Bet #6",
					Message: _err.Error(),
				},
			}
		}

		// == Update Balance #7 ==
		user.Balance = balance
		if err := it.usecase.UpdateUser(user); err != nil {

			return response.ProtoBuf{
				Code: http.StatusInternalServerError,

				Data: &pb.Error{
					Code:    uint32(http.StatusInternalServerError),
					Name:    "Update Balance #7",
					Message: err.Error(),
				},
			}
		}

		// == Store Order #8 ==
		if err := it.usecase.StoreOrder(order); err != nil {

			return response.ProtoBuf{
				Code: http.StatusInternalServerError,

				Data: &pb.Error{
					Code:    http.StatusInternalServerError,
					Name:    "Store Order #8",
					Message: err.Error(),
				},
			}
		}

		// == Create Protobuf #9 ==
		data, err := order.ToProto()
		if err != nil {

			return response.ProtoBuf{
				Code: http.StatusInternalServerError,

				Data: &pb.Error{
					Code:    http.StatusInternalServerError,
					Name:    "Create Protobuf #9",
					Message: err.Error(),
				},
			}
		}

		return response.ProtoBuf{
			Code: http.StatusCreated,

			Data: data,
		}
	}

	it.Send(w, main())
}

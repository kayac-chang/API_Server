package order

import (
	"api/model"
	"api/model/pb"
	"api/model/response"
	"api/utils"
	"net/http"
)

// PUT ...
func (it *Handler) PUT(w http.ResponseWriter, r *http.Request) {

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
		if r.Header.Get("Content-Type") != "application/protobuf" {

			return response.ProtoBuf{
				Code: http.StatusBadRequest,

				Data: &pb.Error{
					Code:    http.StatusBadRequest,
					Name:    "Check Content-Type #2",
					Message: "Content-Type must be application/protobuf",
				},
			}
		}

		// == Check Order Exist #3 ==
		orderID := it.URLParam(r, "id")
		order, err := it.usecase.FindOrderByID(orderID)
		if err != nil {
			code := http.StatusNotFound
			if err != model.ErrNotFound {
				code = http.StatusInternalServerError
			}

			return response.ProtoBuf{
				Code: code,

				Data: &pb.Error{
					Code:    uint32(code),
					Name:    "Check Order Exist #3",
					Message: err.Error(),
				},
			}
		}

		// == Check Exist #4 ==
		task1 := utils.Promisefy(func() (interface{}, error) {
			return it.usecase.FindGameByID(order.GameID)
		})
		task2 := utils.Promisefy(func() (interface{}, error) {
			return it.usecase.FindUserByID(order.UserID)
		})
		res, err := utils.WaitAll(task1, task2)
		if err != nil {

			code := http.StatusNotFound
			if err != model.ErrNotFound {
				code = http.StatusInternalServerError
			}

			return response.ProtoBuf{
				Code: code,

				Data: &pb.Error{
					Code:    uint32(code),
					Name:    "Check Exist #4",
					Message: err.Error(),
				},
			}
		}

		game := res[0].(*model.Game)
		user := res[1].(*model.User)

		// == Parse ProtoBuf #5 ==
		req, err := it.Parse(r.Body)
		if err != nil {

			return response.ProtoBuf{
				Code: http.StatusBadRequest,

				Data: &pb.Error{
					Code:    http.StatusBadRequest,
					Name:    "Parse ProtoBuf #5",
					Message: err.Error(),
				},
			}
		}

		// == Update Order ==
		order.State = req.State

		switch order.State {

		case model.Completed:

			// == Checkout Order #7 ==
			order.Win = req.Win
			if err := it.usecase.Checkout(user, game, order); err != nil {

				return response.ProtoBuf{
					Code: http.StatusInternalServerError,

					Data: &pb.Error{
						Code:    http.StatusInternalServerError,
						Name:    "Checkout Order #7",
						Message: err.Error(),
					},
				}
			}

		case model.Rejected:

			return response.ProtoBuf{
				Code: http.StatusBadRequest,

				Data: &pb.Error{
					Code:    http.StatusBadRequest,
					Name:    "Rejected #7",
					Message: "Not Implement",
				},
			}

		case model.Issue:

			return response.ProtoBuf{
				Code: http.StatusBadRequest,

				Data: &pb.Error{
					Code:    http.StatusBadRequest,
					Name:    "Issue #7",
					Message: "Not Implement",
				},
			}

		default:
			return response.ProtoBuf{
				Code: http.StatusInternalServerError,

				Data: &pb.Error{
					Code:    http.StatusInternalServerError,
					Name:    "Update Order #7",
					Message: "Not Support Order State",
				},
			}
		}

		// == Create Protobuf #6 ==
		data, err := order.ToProto()
		if err != nil {

			return response.ProtoBuf{
				Code: http.StatusInternalServerError,

				Data: &pb.Error{
					Code:    http.StatusInternalServerError,
					Name:    "Create Protobuf #6",
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

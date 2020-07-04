package order

import (
	"api/model"
	"api/model/pb"
	"api/model/response"
	"api/utils"
	"fmt"
	"net/http"
)

// PUT ...
func (it *Handler) PUT(w http.ResponseWriter, r *http.Request) {

	main := func() interface{} {

		// == Check Authorization #1 ==
		token := r.Header.Get("Authorization")
		if err := it.usecase.Auth(token); err != nil {

			return response.ProtoBuf{
				Code: http.StatusUnauthorized,

				Data: &pb.Error{
					Code:    http.StatusUnauthorized,
					Name:    "Check Authorization #1",
					Message: err.Error(),
				},
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

			return response.ProtoBuf{
				Code: http.StatusNotFound,

				Data: &pb.Error{
					Code:    http.StatusNotFound,
					Name:    "Check Order Exist #3",
					Message: err.Error(),
				},
			}
		}

		// == Check Exist #5 ==
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
					Name:    "Check Exist #5",
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

		fmt.Printf("%+v\n", req)

		switch req.State {

		case model.Completed:

			// == Checkout Order #6 ==
			order.State = model.Completed
			order.Win = req.Win
			order.CompletedAt = req.CompletedAt

			if err := it.usecase.Checkout(user, game, order); err != nil {

				return response.ProtoBuf{
					Code: http.StatusInternalServerError,

					Data: &pb.Error{
						Code:    http.StatusInternalServerError,
						Name:    "Checkout Order #6",
						Message: err.Error(),
					},
				}
			}

		case model.Rejected:

			return response.ProtoBuf{
				Code: http.StatusBadRequest,

				Data: &pb.Error{
					Code:    http.StatusBadRequest,
					Name:    "Rejected #6",
					Message: "Not Implement",
				},
			}

		case model.Issue:

			return response.ProtoBuf{
				Code: http.StatusBadRequest,

				Data: &pb.Error{
					Code:    http.StatusBadRequest,
					Name:    "Issue #6",
					Message: "Not Implement",
				},
			}

		default:
			return response.ProtoBuf{
				Code: http.StatusInternalServerError,

				Data: &pb.Error{
					Code:    http.StatusInternalServerError,
					Name:    "Update Order #6",
					Message: "Not Support Order State",
				},
			}
		}

		// == Create Protobuf #7 ==
		data, err := order.ToProto()
		if err != nil {

			return response.ProtoBuf{
				Code: http.StatusInternalServerError,

				Data: &pb.Error{
					Code:    http.StatusInternalServerError,
					Name:    "Create Protobuf #7",
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

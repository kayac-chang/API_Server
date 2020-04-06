package order

import (
	"api/model/pb"
	"api/model/response"
	"net/http"

	"github.com/golang/protobuf/ptypes"
)

func (it *Handler) POST(w http.ResponseWriter, r *http.Request) {

	order, err := it.usecase.Parse(r.Body)
	if err != nil {

		it.Send(w, response.ProtoBuf{
			Code: http.StatusBadRequest,

			Data: &pb.Error{
				Code:    http.StatusBadRequest,
				Name:    "Unexpect Payload",
				Message: err.Error(),
			},
		})

		return
	}

	if err := it.usecase.Create(order); err != nil {

		it.Send(w, response.ProtoBuf{
			Code: http.StatusNotAcceptable,

			Data: &pb.Error{
				Code:    http.StatusNotAcceptable,
				Name:    "Create Order Failed",
				Message: err.Error(),
			},
		})

		return
	}

	// === Send ProtoBuf ===
	created, err := ptypes.TimestampProto(order.CreatedAt.Time)
	if err != nil {

		it.Send(w, response.ProtoBuf{
			Code: http.StatusInternalServerError,

			Data: &pb.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Server Error",
				Message: err.Error(),
			},
		})

		return
	}

	it.Send(w, response.ProtoBuf{
		Code: http.StatusCreated,

		Data: &pb.Order{
			OrderId:   order.ID,
			GameId:    order.GameID,
			UserId:    order.UserID,
			State:     order.State.PbState(),
			Bet:       order.Bet,
			CreatedAt: created,
		},
	})
}

package order

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/framework/server"
	"api/model"
	"api/model/pb"
	"api/model/response"
	order "api/usecase/order"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/golang/protobuf/ptypes"
)

type Handler struct {
	*server.Server
	env     *env.Env
	usecase *order.Usecase
}

func New(s *server.Server, e *env.Env, db *postgres.DB, c *cache.Cache) *Handler {

	return &Handler{
		s,
		e,
		order.New(e, db, c),
	}
}

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

func (it *Handler) PUT(w http.ResponseWriter, r *http.Request) {

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

	order.ID = chi.URLParam(r, "order_id")

	switch order.State {

	case model.Completed:
		order, err = it.usecase.Checkout(order.ID)

	case model.Rejected:
		// TODO

	default:
		// TODO
	}

	if err != nil {
		it.Send(w, response.ProtoBuf{
			Code: http.StatusNotAcceptable,

			Data: &pb.Error{
				Code:    http.StatusNotAcceptable,
				Name:    "Error",
				Message: err.Error(),
			},
		})

		return
	}

	// === Send ProtoBuf ===
	completed, err := ptypes.TimestampProto(order.CompletedAt.Time)
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
			OrderId:     order.ID,
			GameId:      order.GameID,
			UserId:      order.UserID,
			State:       order.State.PbState(),
			Bet:         order.Bet,
			CreatedAt:   created,
			CompletedAt: completed,
		},
	})
}

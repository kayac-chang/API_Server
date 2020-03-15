package order

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/framework/server"
	"api/model"
	"api/model/pb"
	"api/model/request"
	"api/model/response"
	order "api/usecase/order"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/golang/protobuf/ptypes"
)

type handler struct {
	*server.Server
	env     *env.Env
	usecase *order.Usecase
}

func New(e *env.Env, db *postgres.DB, c *cache.Cache) {

	s := server.New(e)

	it := handler{
		s,
		e,
		order.New(e, db, c),
	}

	s.Route("/"+e.API.Version, func(s chi.Router) {
		s.With(it.Order).Post("/orders", it.POST)
		// s.Put("/orders/{order_id}", it.PUT)
	})

	s.Listen(e.API.OrderPort)
}

func (it *handler) POST(w http.ResponseWriter, r *http.Request) {

	order := r.Context().Value(request.ORDER).(*model.Order)

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

func (it *handler) PUT(w http.ResponseWriter, r *http.Request) {

	order := r.Context().Value(request.ORDER).(*model.Order)

}

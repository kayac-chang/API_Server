package order

import (
	"api/env"
	"api/framework/server"
	"api/model"
	"api/model/pb"
	"api/model/response"
	"api/order/repo/cache"
	"api/order/repo/postgres"
	"api/order/usecase"
	userCache "api/user/repo/cache"
	"io/ioutil"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/labstack/gommon/log"
)

type handler struct {
	*server.Server

	usecase usecase.Usecase
}

func New(e *env.Env) {

	s := server.New()

	c := cache.New()
	db := postgres.New(e.Postgres.ToURL(), 30)
	uc := userCache.New()

	it := handler{
		s,
		usecase.New(e, db, c, uc),
	}

	s.Route("/"+e.API.Version, func(r chi.Router) {
		s.Post("/orders", it.POST)
		s.Put("/orders/{order_id}", it.PUT)
	})

	http.ListenAndServe(e.API.OrderPort, s)
}

func (it *handler) POST(w http.ResponseWriter, r *http.Request) {

	reqByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("%s\n", err.Error())

		res := response.ProtoBuf{
			Code: http.StatusBadRequest,
			Data: &pb.Error{
				Code:    http.StatusBadRequest,
				Name:    "Unexpect Payload",
				Message: model.ErrUnexpectPayload.Error(),
			},
		}

		it.SendProtoBuf(w, res)
		return
	}

	req := pb.Order{}

	if err := proto.Unmarshal(reqByte, &req); err != nil {
		log.Errorf("%s\n", err.Error())

		res := response.ProtoBuf{
			Code: http.StatusBadRequest,
			Data: &pb.Error{
				Code:    http.StatusBadRequest,
				Name:    "Unexpect Payload",
				Message: model.ErrUnexpectPayload.Error(),
			},
		}

		it.SendProtoBuf(w, res)
		return
	}

	order := model.Order{
		UserID: req.GetUserId(),
		GameID: req.GetGameId(),
		Bet:    req.GetBet(),
	}

	if err := it.usecase.Store(&order); err != nil {
		log.Errorf("%s\n", err.Error())

		res := response.ProtoBuf{
			Code: http.StatusInternalServerError,
			Data: &pb.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Server Error",
				Message: err.Error(),
			},
		}

		it.SendProtoBuf(w, res)
		return
	}

	states := map[model.State]pb.Order_State{
		model.Pending:   pb.Order_Pending,
		model.Completed: pb.Order_Completed,
		model.Rejected:  pb.Order_Rejected,
	}

	created, err := ptypes.TimestampProto(*order.CreatedAt)
	if err != nil {
		log.Errorf("%s\n", err.Error())

		res := response.ProtoBuf{
			Code: http.StatusInternalServerError,
			Data: &pb.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Server Error",
				Message: err.Error(),
			},
		}

		it.SendProtoBuf(w, res)
		return
	}

	res := response.ProtoBuf{
		Code: http.StatusCreated,
		Data: &pb.Order{
			OrderId:   order.ID,
			GameId:    order.GameID,
			UserId:    order.UserID,
			State:     states[order.State],
			Bet:       order.Bet,
			CreatedAt: created,
		},
	}

	it.SendProtoBuf(w, res)
}

func (it *handler) PUT(w http.ResponseWriter, r *http.Request) {

	reqByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("%s\n", err.Error())

		res := response.ProtoBuf{
			Code: http.StatusBadRequest,
			Data: &pb.Error{
				Code:    http.StatusBadRequest,
				Name:    "Unexpect Payload",
				Message: model.ErrUnexpectPayload.Error(),
			},
		}

		it.SendProtoBuf(w, res)
		return
	}

	req := pb.Order{}
	if err := proto.Unmarshal(reqByte, &req); err != nil {
		log.Errorf("%s\n", err.Error())

		res := response.ProtoBuf{
			Code: http.StatusBadRequest,
			Data: &pb.Error{
				Code:    http.StatusBadRequest,
				Name:    "Unexpect Payload",
				Message: model.ErrUnexpectPayload.Error(),
			},
		}

		it.SendProtoBuf(w, res)
		return
	}

	order := model.Order{
		ID: it.URLParam(r, "order_id"),
	}

	if err := it.usecase.Find(&order); err != nil {
		log.Errorf("%s\n", err.Error())

		res := response.ProtoBuf{
			Code: http.StatusNotFound,
			Data: &pb.Error{
				Code:    http.StatusNotFound,
				Name:    "Not Found",
				Message: err.Error(),
			},
		}

		it.SendProtoBuf(w, res)
		return
	}

	completed, err := ptypes.Timestamp(req.GetCompletedAt())
	if err != nil {
		log.Errorf("%s\n", err.Error())

		res := response.ProtoBuf{
			Code: http.StatusInternalServerError,
			Data: &pb.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Server Error",
				Message: err.Error(),
			},
		}

		it.SendProtoBuf(w, res)
		return
	}

	order.CompletedAt = &completed

	if err := it.usecase.Checkout(&order); err != nil {
		log.Errorf("%s\n", err.Error())

		res := response.ProtoBuf{
			Code: http.StatusInternalServerError,
			Data: &pb.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Server Error",
				Message: err.Error(),
			},
		}

		it.SendProtoBuf(w, res)
		return
	}

	states := map[model.State]pb.Order_State{
		model.Pending:   pb.Order_Pending,
		model.Completed: pb.Order_Completed,
		model.Rejected:  pb.Order_Rejected,
	}

	created, err := ptypes.TimestampProto(*order.CreatedAt)
	if err != nil {
		log.Errorf("%s\n", err.Error())

		res := response.ProtoBuf{
			Code: http.StatusInternalServerError,
			Data: &pb.Error{
				Code:    http.StatusInternalServerError,
				Name:    "Server Error",
				Message: err.Error(),
			},
		}

		it.SendProtoBuf(w, res)
		return
	}

	res := response.ProtoBuf{
		Code: http.StatusAccepted,
		Data: &pb.Order{
			OrderId:     order.ID,
			GameId:      order.GameID,
			UserId:      order.UserID,
			State:       states[order.State],
			Bet:         order.Bet,
			CreatedAt:   created,
			CompletedAt: req.GetCompletedAt(),
		},
	}

	it.SendProtoBuf(w, res)
}

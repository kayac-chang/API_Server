package orders

import (
	"context"
	"net/http"

	"github.com/KayacChang/API_Server/model"
	"github.com/KayacChang/API_Server/orders/repo"
	"github.com/KayacChang/API_Server/orders/usecase"
	"github.com/KayacChang/API_Server/system/env"
	"github.com/KayacChang/API_Server/system/log"
	"github.com/KayacChang/API_Server/system/web"
	"github.com/labstack/echo/v4"
)

// New create orders service
func New(cfg env.Config) {

	server := web.NewServer()

	logic := usecase.New(
		repo.New(cfg.Postgres),
	)

	server.Use(web.Bind("order", newOrder))

	server.POST("/orders", create(logic))

	log.Fatal(server.StartTLS(":8080", ".private/cert.pem", ".private/key.pem"))
}

func newOrder() interface{} {

	return &model.Order{}
}

func create(logic *usecase.Usecase) echo.HandlerFunc {

	handler := func(ctx echo.Context) (interface{}, error) {

		c := ctx.Request().Context()

		if c == nil {
			c = context.Background()
		}

		user := ctx.Get("order").(*model.Order)

		err := logic.Create(c, user)

		return user, err
	}

	return web.Send(http.StatusCreated, handler)
}

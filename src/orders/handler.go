package orders

import (
	"context"
	"net/http"

	"server/model"
	"server/orders/repo"
	"server/orders/usecase"
	"server/system/log"
	"server/system/web"

	"github.com/labstack/echo/v4"
)

// New create orders service
func New() {

	server := web.NewServer()

	logic := usecase.New(repo.New())

	server.Use(web.Bind("order", newOrder))

	server.POST("/orders", create(logic))

	// TODO: PUT /orders/:order_id

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

		order := ctx.Get("order").(*model.Order)

		err := logic.Create(c, order)

		return order, err
	}

	return web.Send(http.StatusCreated, handler)
}

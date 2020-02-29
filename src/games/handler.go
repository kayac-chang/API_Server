package games

import (
	"context"
	"log"
	"net/http"

	"github.com/KayacChang/API_Server/games/entity"
	"github.com/KayacChang/API_Server/games/repo"
	"github.com/KayacChang/API_Server/games/usecase"
	"github.com/KayacChang/API_Server/system"
	"github.com/KayacChang/API_Server/system/env"
	"github.com/KayacChang/API_Server/system/web"
	"github.com/labstack/echo/v4"
)

// New create game service
func New(cfg env.Config) {

	server := web.NewServer()

	logic := usecase.New(
		repo.New(cfg.Postgres),
	)

	server.POST("/games", post(logic), bindGame, validate)

	server.GET("/games", get(logic))

	log.Fatal(server.StartTLS(":8080", ".private/cert.pem", ".private/key.pem"))
}

func bindGame(next echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		var game entity.Game

		err := ctx.Bind(&game)

		if err != nil {
			/*
				422 Unprocessable Entity
				the server understands the content type of the request entity,
				and the syntax of the request entity is correct,
				but it was unable to process the contained instructions.
			*/
			return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		}

		ctx.Set("game", &game)

		return next(ctx)
	}
}

func validate(next echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		// TODO: validation

		isValid := true

		if !isValid {
			/*
				400 Bad Request
				the server cannot or will not process the request due to something that is perceived to be a client error
				(e.g., malformed request syntax, invalid request message framing, or deceptive request routing).
			*/
			return ctx.JSON(http.StatusBadRequest, "")
		}

		return next(ctx)
	}
}

func post(usecase *usecase.Usecase) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		c := ctx.Request().Context()

		if c == nil {
			c = context.Background()
		}

		game := ctx.Get("game").(*entity.Game)

		return exec(
			ctx,
			http.StatusCreated,
			func() (interface{}, error) { return game, usecase.Store(c, game) },
		)
	}
}

func get(usecase *usecase.Usecase) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		c := ctx.Request().Context()

		if c == nil {
			c = context.Background()
		}

		games := &[]entity.Game{}

		return exec(
			ctx,
			http.StatusOK,
			func() (interface{}, error) { return games, usecase.Find(c, games) },
		)
	}
}

type execFunc func() (interface{}, error)

func exec(ctx echo.Context, code int, fn execFunc) error {

	res, err := fn()

	if err != nil {

		code := system.GetStatusCode(err)

		res := web.ResponseError{
			Message: err.Error(),
		}

		return ctx.JSON(code, res)
	}

	return ctx.JSON(code, res)
}

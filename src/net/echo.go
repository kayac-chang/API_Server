package net

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ======================================

type Server struct {
	*echo.Echo
}

type Context interface {
	echo.Context
}

type HandlerFunc func(ctx Context) error

// ======================================

func New() *Server {

	server := &Server{echo.New()}

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	return server
}

func (it *Server) Listen(port string) {

	it.Logger.Fatal(it.Start(port))
}

func (it *Server) Get(path string, handle HandlerFunc) {

	it.GET(path, func(ctx echo.Context) error {
		return handle(ctx)
	})
}

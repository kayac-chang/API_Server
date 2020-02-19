package game

import (
	"net/http"

	"github.com/KayacChang/API_Server/net"
)

func Mount(app *net.Server) {

	app.Get("/games", fetch)
	// app.POST("/games", fetch)
}

func fetch(ctx net.Context) error {

	return ctx.String(http.StatusOK, "Hello, World!")
}

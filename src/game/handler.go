package game

import (
	"net/http"

	"github.com/KayacChang/API_Server/system"
)

func Mount(app *system.Server) {

	app.Get("/games", fetch)
	// app.POST("/games", fetch)
}

func fetch(ctx system.Context) error {

	return ctx.String(http.StatusOK, "Hello, World!")
}

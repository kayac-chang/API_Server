package game

import (
	"log"
	"net/http"

	"github.com/KayacChang/API_Server/system/db"
	"github.com/KayacChang/API_Server/system/env"
	"github.com/KayacChang/API_Server/system/web"
)

// Serve Start game service
func Serve() {

	server := web.NewServer()

	server.Get("/games", hello())

	log.Fatal(server.Start(":8080"))
}

func hello() web.HandlerFunc {

	db.New(env.Postgres().ToURL())

	return func(c web.Context) error {

		return c.String(http.StatusOK, "hello")
	}
}

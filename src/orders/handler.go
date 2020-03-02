package orders

import (
	"log"

	"github.com/KayacChang/API_Server/games/usecase"
	"github.com/KayacChang/API_Server/orders/repo"
	"github.com/KayacChang/API_Server/system/env"
	"github.com/KayacChang/API_Server/system/web"
)

// New create game service
func New(cfg env.Config) {

	server := web.NewServer()

	logic := usecase.New(
		repo.New(cfg.Postgres),
	)

	log.Fatal(server.StartTLS(":8080", ".private/cert.pem", ".private/key.pem"))
}

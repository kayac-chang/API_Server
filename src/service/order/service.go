package order

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/framework/server"
	order "api/usecase/order"
)

type Handler struct {
	*server.Server
	env     env.Env
	usecase *order.Usecase
}

func New(server *server.Server, env env.Env, db *postgres.DB, c *cache.Cache) *Handler {

	return &Handler{
		server,
		env,
		order.New(env, db, c),
	}
}

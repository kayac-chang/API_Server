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
	env     *env.Env
	usecase *order.Usecase
}

func New(s *server.Server, e *env.Env, db *postgres.DB, c *cache.Cache) *Handler {

	return &Handler{
		s,
		e,
		order.New(e, db, c),
	}
}

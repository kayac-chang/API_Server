package admin

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/framework/server"
	admin "api/usecase/admin"
)

type Handler struct {
	*server.Server
	env     *env.Env
	usecase *admin.Usecase
}

func New(s *server.Server, e *env.Env, db *postgres.DB, c *cache.Cache) *Handler {

	return &Handler{
		s,
		e,
		admin.New(e, db, c),
	}
}

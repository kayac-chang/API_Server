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
	env     env.Env
	usecase *admin.Usecase
}

func New(server *server.Server, env env.Env, db *postgres.DB, c *cache.Cache) *Handler {

	return &Handler{
		server,
		env,
		admin.New(env, db, c),
	}
}

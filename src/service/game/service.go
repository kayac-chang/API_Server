package game

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/framework/server"
	admin "api/usecase/admin"
	game "api/usecase/game"
)

type Handler struct {
	*server.Server
	env   env.Env
	game  *game.Usecase
	admin *admin.Usecase
}

func New(s *server.Server, env env.Env, db *postgres.DB, c *cache.Cache) *Handler {

	return &Handler{
		s,
		env,
		game.New(env, db, c),
		admin.New(env, db, c),
	}
}

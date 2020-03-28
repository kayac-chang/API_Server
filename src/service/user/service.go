package user

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/framework/server"
	game "api/usecase/game"
	user "api/usecase/user"
)

type Handler struct {
	*server.Server
	env      *env.Env
	userCase *user.Usecase
	gameCase *game.Usecase
}

func New(s *server.Server, e *env.Env, db *postgres.DB, c *cache.Cache) *Handler {

	return &Handler{
		s,
		e,
		user.New(e, db, c),
		game.New(e, db, c),
	}
}

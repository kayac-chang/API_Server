package game

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/framework/server"
	game "api/usecase/game"
)

type Handler struct {
	*server.Server
	env     *env.Env
	usecase *game.Usecase
}

func New(s *server.Server, e *env.Env, db *postgres.DB, c *cache.Cache) *Handler {

	return &Handler{
		s,
		e,
		game.New(e, db, c),
	}

}

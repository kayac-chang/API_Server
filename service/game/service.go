package game

import (
	"api/env"
	"api/framework/server"
	game "api/usecase/game"
)

// Handler game service instance
type Handler struct {
	server.Server
	env     env.Env
	usecase game.Usecase
}

// New create game service instance
func New(s server.Server, env env.Env, usecase game.Usecase) Handler {

	return Handler{s, env, usecase}
}

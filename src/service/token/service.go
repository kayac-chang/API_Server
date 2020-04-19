package token

import (
	"api/env"
	"api/framework/server"

	game "api/usecase/game"
	token "api/usecase/token"
)

type Handler struct {
	server.Server
	env   env.Env
	token token.Usecase
	game  game.Usecase
}

func New(server server.Server, env env.Env, token token.Usecase, game game.Usecase) Handler {

	return Handler{server, env, token, game}
}

func (it Handler) getHref(url string) string {

	return it.env.Service.Domain + "/" + it.env.API.Version + url
}

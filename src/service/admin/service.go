package admin

import (
	"api/env"
	"api/framework/server"
	admin "api/usecase/admin"
)

type Handler struct {
	server.Server
	env     env.Env
	usecase *admin.Usecase
}

func New(server server.Server, env env.Env) *Handler {

	return &Handler{
		server,
		env,
		admin.New(env, db, c),
	}
}

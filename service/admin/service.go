package admin

import (
	"api/env"
	"api/framework/server"
	"api/usecase/admin"
)

// Handler service instance
type Handler struct {
	server.Server
	env     env.Env
	usecase admin.Usecase
}

// New admin service
func New(server server.Server, env env.Env, usecase admin.Usecase) Handler {

	return Handler{
		server,
		env,
		usecase,
	}
}

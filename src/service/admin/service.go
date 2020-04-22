package admin

import (
	"api/env"
	"api/framework/redis"
	"api/framework/server"
	"api/usecase/admin"
)

type Handler struct {
	server.Server
	env     env.Env
	usecase admin.Usecase
}

func New(server server.Server, env env.Env, db redis.Redis) *Handler {

	return &Handler{
		server,
		env,
		admin.New(env, db),
	}
}

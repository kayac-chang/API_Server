package order

import (
	"api/env"
	"api/framework/server"
	order "api/usecase/order"
)

// Handler ...
type Handler struct {
	server.Server
	env     env.Env
	usecase order.Usecase
}

// New ...
func New(server server.Server, env env.Env, order order.Usecase) Handler {

	return Handler{server, env, order}
}

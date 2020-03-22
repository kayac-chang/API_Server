package main

import (
	"api/env"
	"api/framework/cache"
	"api/framework/postgres"
	"api/framework/server"
	"api/service/game"
	"api/service/order"
	"api/service/user"

	"github.com/go-chi/chi"
)

func main() {

	// === Framework ===
	env := env.New()
	cache := cache.Get()
	db := postgres.New(env.Postgres.ToURL())
	it := server.New(env)

	// === Handler ===
	game := game.New(it, env, db, cache)
	user := user.New(it, env, db, cache)
	order := order.New(it, env, db, cache)

	it.Route("/"+env.API.Version, func(server chi.Router) {
		// === Game ===
		server.Get("/games", game.GET)
		server.With(it.ParseJSON).Post("/games", game.POST)

		// === User ===
		server.With(it.ParseJSON).Post("/token", user.POST)
		server.With(it.User).Get("/auth", user.Auth)

		// === Order ===
		server.Route("/orders", func(server chi.Router) {
			server.With(it.Order).Post("/", order.POST)
			server.With(it.Order).Put("/{order_id}", order.PUT)
		})
	})

	it.Listen(env.Service.Port)
}

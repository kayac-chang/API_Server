package main

import (
	"api/env"
	"api/framework/redis"
	"api/framework/server"
	"api/service/admin"
	"api/service/token"

	gameusecase "api/usecase/game"
	tokenusecase "api/usecase/token"

	"github.com/go-chi/chi"
)

func main() {

	// === Framework ===
	env := env.New()
	redis := redis.New(env.Redis.HOST, env.Redis.PORT)
	it := server.New(env)

	// === Usecase ===
	tokenUsecase := tokenusecase.New(env, redis)
	gameUsecase := gameusecase.New(env, redis)

	// === Handler ===
	token := token.New(it, env, tokenUsecase, gameUsecase)
	admin := admin.New(it, env, redis)

	it.Route("/"+env.API.Version, func(router chi.Router) {
		// === Game ===
		// router.Route("/games", func(router chi.Router) {
		// 	router.Get("/", game.GET_ALL)
		// 	router.Get("/{name}", game.GET)
		// 	router.Post("/", game.POST)
		// 	router.Put("/{name}", game.PUT)
		// })

		// === Admin ===
		router.Route("/admins", func(router chi.Router) {
			router.Post("/", admin.POST)

			// router.Route("/tokens", func(router chi.Router) {
			// 	router.Post("/", admin.Auth)
			// })
		})

		// === User ===
		// router.Route("/users", func(router chi.Router) {
		// 	router.Get("/{token}", token.Get)
		// })

		// === Token ===
		router.Route("/tokens", func(router chi.Router) {
			router.Post("/", token.POST)
		})

		// === Order ===
		// router.Route("/orders", func(router chi.Router) {
		// 	router.Post("/", order.POST)
		// 	router.Put("/{order_id}", order.PUT)
		// })
	})

	it.Listen(env.Service.Port)
}

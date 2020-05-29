package main

import (
	"api/env"
	"api/framework/postgres"
	"api/framework/redis"
	"api/framework/server"
	"api/service/admin"
	"api/service/game"
	"api/service/order"
	"api/service/suborder"
	"api/service/token"

	adminusecase "api/usecase/admin"
	gameusecase "api/usecase/game"
	orderusecase "api/usecase/order"
	tokenusecase "api/usecase/token"

	"github.com/go-chi/chi"
)

func main() {

	// === Framework ===
	env := env.New()
	db := postgres.New(env)
	redis := redis.New(env)
	it := server.New(env)

	// === Usecase ===
	tokenUsecase := tokenusecase.New(env, redis, db)
	gameUsecase := gameusecase.New(env, redis, db)
	adminUsecase := adminusecase.New(env, redis, db)
	orderUsecase := orderusecase.New(env, redis, db)

	// === Handler ===
	token := token.New(it, env, tokenUsecase)
	admin := admin.New(it, env, adminUsecase)
	game := game.New(it, env, gameUsecase)
	order := order.New(it, env, orderUsecase)
	subOrder := suborder.New(it, env, orderUsecase)

	it.Route("/"+env.API.Version, func(router chi.Router) {
		// === Game ===
		router.Route("/games", func(router chi.Router) {
			router.Get("/", game.GETALL)
			router.Get("/{id}", game.GET)
			router.Post("/", game.POST)
			router.Put("/{id}", game.PUT)
		})

		// === Admin ===
		router.Route("/admins", func(router chi.Router) {
			router.Post("/", admin.POST)

			router.Route("/tokens", func(router chi.Router) {
				router.Post("/", admin.Auth)
			})
		})

		// === User ===
		router.Route("/users", func(router chi.Router) {
			router.Get("/{token}", token.GET)
		})

		// === Token ===
		router.Route("/tokens", func(router chi.Router) {
			router.Post("/", token.POST)
		})

		// === Order ===
		router.Route("/orders", func(router chi.Router) {
			router.Post("/", order.POST)
			router.Put("/{id}", order.PUT)
		})

		// === SubOrder ===
		router.Route("/sub-orders", func(router chi.Router) {
			router.Post("/", subOrder.POST)
		})
	})

	it.Listen(env.Service.Port)
}

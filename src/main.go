package main

import (
	"net/http"
	"user/api"
	"user/repo/cache"
	"user/repo/postgres"
	"user/usecase"
	"user/utils/env"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	e := env.New()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cache := cache.New()
	db := postgres.New(e.Postgres.ToURL(), 30)

	logic := usecase.New(db, cache)

	handler := api.New(logic)

	r.Post("/token", handler.POST)

	http.ListenAndServe(":8000", r)
}

package server

import (
	"api/env"
	"log"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	*chi.Mux
	env env.Env
}

func New(e env.Env) Server {

	server := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	server.Use(cors.Handler)
	server.Use(middleware.RequestID)
	server.Use(middleware.RealIP)
	server.Use(middleware.Logger)
	server.Use(middleware.Recoverer)

	return Server{server, e}
}

func (it Server) Listen(port string) {

	log.Fatal(
		http.ListenAndServe(port, it),
	)
}

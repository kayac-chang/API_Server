package main

import (
	"github.com/KayacChang/API_Server/games"
	"github.com/KayacChang/API_Server/system/env"
)

func main() {

	cfg := env.Config{
		Postgres: env.PostgresConfig{
			"host":   "localhost",
			"port":   "5432",
			"user":   "kayac",
			"dbname": "postgres",
		},
	}

	games.New(cfg)
}

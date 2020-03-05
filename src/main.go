package main

import "server/users"

func main() {

	// cfg := env.Config{
	// 	Postgres: env.PostgresConfig{
	// 		"host":     "localhost",
	// 		"port":     "5432",
	// 		"user":     "postgres",
	// 		"password": "123456",
	// 		"dbname":   "postgres",
	// 	},
	// }

	// games.New(cfg)

	users.New()

	// orders.New()
}

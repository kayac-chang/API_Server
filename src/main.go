package main

import (
	"fmt"
	"log"
	"server/model"
	"server/system/cache"
	"server/utils"

	"github.com/go-redis/redis"
)

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

	// users.New(cfg)

	// orders.New(cfg)

	user := &model.User{
		ID:       utils.MD5("kayac"),
		Username: "kayac",
		Balance:  123456.0,
	}

	// cache.Set(user.ID, user, 1*time.Minute)

	var res model.User

	err := cache.Get(user.ID, &res)

	if err == redis.Nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", res)
}

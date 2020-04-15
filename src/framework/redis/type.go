package redis

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	*redis.Pool
}

func New(host string, port string) *Redis {

	pool := &redis.Pool{

		MaxIdle: 3,

		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {

			addr := host + ":" + port

			conn, err := redis.Dial("tcp", addr)

			if err != nil {
				log.Fatal(err)
			}

			return conn, err
		},
	}

	return &Redis{pool}
}

package redis

import (
	"log"

	"github.com/mediocregopher/radix/v3"
)

// Redis radix client wrapper
type Redis struct {
	pool *radix.Pool
}

// Action radix wrapper
type Action interface {
	radix.Action
}

// New return Redis client
func New(host string, port string) Redis {

	pool, err := radix.NewPool("tcp", host+":"+port, 10)
	if err != nil {
		log.Fatal("Init: Failed when connect to Redis...")
	}

	return Redis{pool}
}

// Set return an action for set command
func (it Redis) Set(key string, val interface{}) error {

	switch val := val.(type) {

	case string:
		return it.pool.Do(radix.Cmd(nil, "SET", key, val))
	}

	panic("Exception on Redis.Set: Not support val type")
}

package redis

import (
	"encoding/json"
	"log"

	"github.com/mediocregopher/radix/v3"
	errs "github.com/pkg/errors"
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

// Set is used to perform redis SET command,
// it will translate data in to json string and store into redis
func (it Redis) Set(key string, val interface{}) error {

	json, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return it.pool.Do(radix.Cmd(nil, "SET", key, string(json)))
}

// Get is used to perform redis GET command
// it will get back json string by key, and parse to val
func (it Redis) Get(key string, val interface{}) error {

	var res string

	if err := it.pool.Do(radix.Cmd(&res, "GET", key)); err != nil {
		return err
	}

	if res == "" {
		return errs.Errorf("GET key [%s] is empty in redis", key)
	}

	return json.Unmarshal([]byte(res), val)
}

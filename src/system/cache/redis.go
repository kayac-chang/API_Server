package cache

import (
	"encoding/json"
	"time"

	"server/system/env"
	"server/system/log"
	"server/utils"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func init() {
	cfg := env.Redis()

	rdb = redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		PoolSize:     cfg.PoolSize,
		PoolTimeout:  cfg.PoolTimeout,
	})

	_, err := rdb.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}
}

func Set(name string, value interface{}, expires time.Duration) error {

	json := utils.Stringify(value)

	return rdb.Set(name, json, expires).Err()
}

func Get(name string, value interface{}) error {

	str, err := rdb.Get(name).Result()

	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(str), value)
}

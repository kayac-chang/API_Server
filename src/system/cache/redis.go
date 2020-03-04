package cache

import (
	"encoding/json"
	"time"

	"server/system/env"
	"server/system/log"
	utils "server/utils/json"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func init() {

	rdb = redis.NewClient(env.Redis())

	_, err := rdb.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}
}

func Set(name string, value interface{}, expires time.Duration) error {

	data := utils.Stringify(value)

	return rdb.Set(name, data, expires).Err()
}

func Get(name string, value interface{}) error {

	str, err := rdb.Get(name).Result()

	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(str), value)
}

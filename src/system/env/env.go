package env

import (
	"log"
	"os"
	"server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

// === Export ===

// Postgres URL
var Postgres func() string

// IsDebug flag for bebug mode
var IsDebug func() bool

// Domain return service domain name
var Domain func() string

// DomainKey return domain key uuid
var DomainKey func() uuid.UUID

// Redis Options
var Redis func() *redis.Options

// === Impl ===
func init() {

	err := godotenv.Load()

	if err != nil {
		log.Panicf("No [ .env ] file found...\n")
	}

	env := struct {
		Debug     bool
		Postgres  map[string]string
		Redis     map[string]string
		Domain    string
		DomainKey uuid.UUID
	}{

		Debug: getEnvAsBool("DEBUG"),

		Postgres: map[string]string{
			"host":     getEnv("PG_HOST"),
			"port":     getEnv("PG_PORT"),
			"user":     getEnv("PG_USER"),
			"password": getEnv("PG_PASSWORD"),
			"dbname":   getEnv("PG_NAME"),
		},

		Redis: map[string]string{
			"addr":          getEnv("REDIS_ADDR"),
			"dial_timeout":  getEnv("REDIS_DIAL_TIMEOUT"),
			"read_timeout":  getEnv("REDIS_READ_TIMEOUT"),
			"write_timeout": getEnv("REDIS_WRITE_TIMEOUT"),
			"pool_size":     getEnv("REDIS_POOL_SIZE"),
			"pool_timeout":  getEnv("REDIS_POOL_TIMEOUT"),
		},

		Domain: getEnv("DOMAIN"),

		DomainKey: uuid.Must(
			uuid.FromString(getEnv("DOMAIN_KEY")),
		),

		// UserRoles: getEnvAsSlice("USER_ROLES", []string{"admin"}, ","),

		// MaxUsers:  getEnvAsInt("MAX_USERS", 1),
	}

	Postgres = func() string {

		cfg := env.Postgres

		data := make([]string, len(cfg))

		for key, val := range cfg {

			data = append(data, key+"="+val)
		}

		return strings.Join(data, " ")
	}

	IsDebug = func() bool {
		return env.Debug
	}

	Domain = func() string {
		return env.Domain
	}

	DomainKey = func() uuid.UUID {
		return env.DomainKey
	}

	redisConfig := &redis.Options{
		Addr:         env.Redis["addr"],
		DialTimeout:  utils.StrToTime(env.Redis["dial_timeout"], time.Second),
		ReadTimeout:  utils.StrToTime(env.Redis["read_timeout"], time.Second),
		WriteTimeout: utils.StrToTime(env.Redis["write_timeout"], time.Second),
		PoolSize:     utils.Number(env.Redis["pool_size"]),
		PoolTimeout:  utils.StrToTime(env.Redis["pool_timeout"], time.Second),
	}

	Redis = func() *redis.Options {

		return redisConfig
	}

	log.Printf("Parse .env: \n%s\n", utils.Jsonify(env))
}

// === Func ===
func getEnv(key string) string {

	value, exists := os.LookupEnv(key)

	if !exists {
		log.Panicf("%s in .env not existed", key)
	}

	return value
}

func getEnvAsBool(key string) bool {

	valStr := getEnv(key)

	val, err := strconv.ParseBool(valStr)

	if err != nil {
		log.Panicf("%s=%s in .env is not boolean value", key, valStr)
	}

	return val
}

func getEnvAsInt(key string) int {

	valStr := getEnv(key)

	val, err := strconv.ParseInt(valStr, 10, 32)

	if err != nil {
		log.Panicf("%s=%s in .env is not int value", key, valStr)
	}

	return int(val)
}

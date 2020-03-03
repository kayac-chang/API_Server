package env

import (
	"log"
	"os"
	"server/utils"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

// === Data Structure ===

// Config Environment variable struct
type Config struct {
	Debug bool

	Postgres PostgresConfig

	Domain string

	DomainKey uuid.UUID

	Redis RedisConfig
}

type PostgresConfig map[string]string

// ToURL helper func to generate datasource string
func (cfg PostgresConfig) ToURL() string {

	data := make([]string, len(cfg))

	for key, val := range cfg {

		data = append(data, key+"="+val)
	}

	return strings.Join(data, " ")
}

type RedisConfig struct {
	Addr         string
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	PoolTimeout  time.Duration
}

// === Export ===

// Postgres getter for getting PostgresConfig
var Postgres func() PostgresConfig

// IsDebug flag for bebug mode
var IsDebug func() bool

// Domain return service domain name
var Domain func() string

// DomainKey return domain key uuid
var DomainKey func() uuid.UUID

var Redis func() RedisConfig

// === Impl ===
func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("No [ .env ] file found...\n")
	}

	env := Config{

		Postgres: map[string]string{
			"host":     getEnv("PG_HOST"),
			"port":     getEnv("PG_PORT"),
			"user":     getEnv("PG_USER"),
			"password": getEnv("PG_PASSWORD"),
			"dbname":   getEnv("PG_NAME"),
		},

		Redis: RedisConfig{
			Addr:         getEnv("REDIS_ADDR"),
			DialTimeout:  getEnvAsTime("REDIS_DIAL_TIMEOUT") * time.Second,
			ReadTimeout:  getEnvAsTime("REDIS_READ_TIMEOUT") * time.Second,
			WriteTimeout: getEnvAsTime("REDIS_WRITE_TIMEOUT") * time.Second,
			PoolSize:     getEnvAsInt("REDIS_POOL_SIZE"),
			PoolTimeout:  getEnvAsTime("REDIS_POOL_TIMEOUT") * time.Second,
		},

		Debug: getEnvAsBool("DEBUG"),

		Domain: getEnv("DOMAIN"),

		DomainKey: uuid.Must(
			uuid.FromString(getEnv("DOMAIN_KEY")),
		),

		// UserRoles: getEnvAsSlice("USER_ROLES", []string{"admin"}, ","),

		// MaxUsers:  getEnvAsInt("MAX_USERS", 1),
	}

	Postgres = func() PostgresConfig {
		return env.Postgres
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

	Redis = func() RedisConfig {
		return env.Redis
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

func getEnvAsTime(key string) time.Duration {

	num := getEnvAsInt(key)

	return time.Duration(num)
}

package env

import (
	"api/utils/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// === Postgres ===

type PostgresConfig map[string]string

func (cfg PostgresConfig) ToURL() string {

	data := make([]string, len(cfg))

	for key, val := range cfg {

		data = append(data, key+"="+val)
	}

	return strings.Join(data, " ")
}

// === Env ===

type Env struct {
	Postgres  PostgresConfig
	ServiceID string
}

func New() *Env {

	err := godotenv.Load()

	if err != nil {
		log.Panicf("No [ .env ] file found...\n")
	}

	env := &Env{

		Postgres: map[string]string{
			"host":     getEnv("PG_HOST"),
			"port":     getEnv("PG_PORT"),
			"user":     getEnv("PG_USER"),
			"password": getEnv("PG_PASSWORD"),
			"dbname":   getEnv("PG_NAME"),
		},

		ServiceID: getEnv("SERVICE_ID"),
	}

	log.Printf("Parse .env: \n%s\n", json.Jsonify(env))

	return env
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

package env

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// === Data Structure ===

// Config Environment variable struct
type Config struct {
	Debug bool

	Postgres PostgresConfig
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

// === Export ===

// ENV getter for getting environment variable
var ENV func() Config

// Postgres getter for getting PostgresConfig
func Postgres() PostgresConfig {
	return ENV().Postgres
}

// IsDebug flag for bebug mode
func IsDebug() bool {
	return ENV().Debug
}

// === Impl ===
func Init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("No [ .env ] file found...\n")
	}

	env := Config{

		Postgres: map[string]string{
			"host":     getEnv("DB_HOST"),
			"port":     getEnv("DB_PORT"),
			"user":     getEnv("DB_USER"),
			"password": getEnv("DB_PASSWORD"),
			"dbname":   getEnv("DB_NAME"),
		},

		Debug: getEnvAsBool("DEBUG"),

		// UserRoles: getEnvAsSlice("USER_ROLES", []string{"admin"}, ","),

		// MaxUsers:  getEnvAsInt("MAX_USERS", 1),
	}

	ENV = func() Config {
		res := env

		return res
	}

	log.Printf("Parse .env: %+v\n", env)
}

// === Func ===
func getEnv(key string) string {

	value, exists := os.LookupEnv(key)

	if !exists {
		log.Fatalf("%s in .env not existed", key)
	}

	return value
}

func getEnvAsBool(key string) bool {

	valStr := getEnv(key)

	val, err := strconv.ParseBool(valStr)

	if err != nil {
		log.Fatalf("%s=%s in .env is not boolean value", key, valStr)
	}

	return val
}

func getEnvAsInt(key string) int {

	valStr := getEnv(key)

	val, err := strconv.ParseInt(valStr, 10, 32)

	if err != nil {
		log.Fatalf("%s=%s in .env is not int value", key, valStr)
	}

	return int(val)
}

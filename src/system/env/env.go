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

// PostgresConfig config for postgresql database
type PostgresConfig struct {
	User string
	DB   string
}

// ToURL helper func to generate datasource string
func (cfg PostgresConfig) ToURL() string {

	data := []string{
		"user=" + cfg.User,
		"dbname=" + cfg.DB,
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
func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("No [ .env ] file found...\n")
	}

	env := Config{

		Postgres: PostgresConfig{
			User: getEnv("USER"),
			DB:   getEnv("DB"),
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

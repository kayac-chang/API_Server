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
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

// ToURL helper func to generate datasource string
func (cfg PostgresConfig) ToURL() string {

	data := []string{
		"host=" + cfg.Host,
		"port=" + cfg.Port,
		"user=" + cfg.User,
		"password=" + cfg.Password,
		"dbname=" + cfg.DB,
		"sslmode=disable",
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
			Host:     getEnv("HOST"),
			Port:     getEnv("PORT"),
			User:     getEnv("USER"),
			Password: getEnv("PASSWORD"),
			DB:       getEnv("DB"),
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

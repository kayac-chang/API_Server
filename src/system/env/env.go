package env

import (
	"log"
	"os"
	"strconv"
	"strings"

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

// Postgres getter for getting PostgresConfig
var Postgres func() PostgresConfig

// IsDebug flag for bebug mode
var IsDebug func() bool

// Domain return service domain name
var Domain func() string

// DomainKey return domain key uuid
var DomainKey func() uuid.UUID

// === Impl ===
func init() {

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

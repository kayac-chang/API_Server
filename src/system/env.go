package system

import (
	"log"
	"os"

	"github.com/KayacChang/API_Server/entity"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("ERROR: No [ .env ] file found...\n")
	}
}

func Env() *entity.Env {

	it := entity.Env{

		Postgres: entity.PostgresConfig{
			User: getEnv("USER", ""),
			DB:   getEnv("DB", ""),
		},

		// DebugMode: getEnvAsBool("DEBUG_MODE", true),

		// UserRoles: getEnvAsSlice("USER_ROLES", []string{"admin"}, ","),

		// MaxUsers:  getEnvAsInt("MAX_USERS", 1),
	}

	log.Printf("Parse [ .env ]: %+v\n", it)

	return &it
}

func getEnv(key string, defaultVal string) string {

	value, exists := os.LookupEnv(key)

	if exists {
		return value
	}

	return defaultVal
}

package env

import (
	"api/utils/json"

	"log"
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
	Postgres PostgresConfig
	Redis    Redis
	Secret   []byte
	Agent    Agent
	Service  Service
	API      API
	SSL      SSL
}

type Service struct {
	ID     string
	Domain string
	Port   string
}

type API struct {
	Version string
}

type Agent struct {
	Domain string
	API    string
	Token  string
}

type SSL struct {
	Cert string
	Key  string
}

type Redis struct {
	HOST string
	PORT string
}

func New() Env {

	err := godotenv.Load()

	if err != nil {
		log.Panicf("No [ .env ] file found...\n")
	}

	env := Env{

		Postgres: map[string]string{
			"host":     getEnv("PG_HOST"),
			"port":     getEnv("PG_PORT"),
			"user":     getEnv("PG_USER"),
			"password": getEnv("PG_PASSWORD"),
			"dbname":   getEnv("PG_NAME"),
		},

		Redis: Redis{
			HOST: getEnv("REDIS_HOST"),
			PORT: getEnv("REDIS_PORT"),
		},

		Service: Service{
			ID:     getEnv("SERVICE_ID"),
			Domain: getEnv("SERVICE_DOMAIN"),
			Port:   ":" + getEnv("SERVICE_PORT"),
		},

		API: API{
			Version: getEnv("API_VERSION"),
		},

		SSL: SSL{
			Cert: getEnv("CERT"),
			Key:  getEnv("KEY"),
		},

		Secret: []byte(getEnv("SECRET")),

		Agent: Agent{
			Domain: getEnv("AGENT_DOMAIN"),
			API:    getEnv("AGENT_API"),
			Token:  getEnv("AGENT_TOKEN"),
		},
	}

	log.Printf("Parse .env: \n%s\n", json.Jsonify(env))

	return env
}

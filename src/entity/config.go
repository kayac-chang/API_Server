package entity

import (
	"strings"
)

type Env struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	User string
	DB   string
}

func (cfg PostgresConfig) ToURL() string {

	data := []string{
		"user=" + cfg.User,
		"dbname=" + cfg.DB,
	}

	return strings.Join(data, " ")
}

package main

import (
	"api/env"
	"api/user"
)

func main() {

	e := env.New()

	user.New(e)
}

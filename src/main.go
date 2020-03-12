package main

import (
	"api/env"
	"api/game"
)

func main() {

	e := env.New()

	game.New(e)
}

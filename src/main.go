package main

import (
	"api/env"
	"api/game"
	"api/order"
	"api/user"
	"sync"
)

func main() {
	e := env.New()

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		game.New(e)

		wg.Done()
	}()

	go func() {
		user.New(e)

		wg.Done()
	}()

	go func() {
		order.New(e)

		wg.Done()
	}()

	wg.Wait()
}

package main

import (
	"api/env"
	"api/order"
	"api/user"
	"sync"
)

func main() {
	e := env.New()

	var wg sync.WaitGroup

	// game.New(e)

	wg.Add(2)

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

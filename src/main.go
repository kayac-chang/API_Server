package main

import (
	"fmt"

	"github.com/KayacChang/API_Server/db"
	"github.com/KayacChang/API_Server/game"
	"github.com/KayacChang/API_Server/net"
)

func main() {

	directory, err := db.Run("psql://kayac@localhost:5432/test")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", directory)

	app := net.New()

	game.Mount(app)

	// Start server
	app.Listen(":8080")
}

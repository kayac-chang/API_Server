package main

import (
	"github.com/KayacChang/API_Server/db"
	"github.com/KayacChang/API_Server/system"
)

func main() {

	env := system.Env()

	db.Run(env)

	// app := net.New()

	// game.Mount(app)

	// app.Listen(":8080")
}

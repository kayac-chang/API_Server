package main

import (
	"net/http"

	"github.com/KayacChang/API_Server/mongo"
	"github.com/KayacChang/API_Server/user"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	db := mongo.New("test")

	user.Serve(r, db)

	http.ListenAndServe(":8080", r)
}

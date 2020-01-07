package main

import (
	"fmt"
	"net/http"

	"github.com/KayacChang/API_Server/games"
	"github.com/KayacChang/API_Server/postgres"

	"github.com/julienschmidt/httprouter"
)

const (
	port = ":8080"
)

func cors(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Access-Control-Request-Method") != "" {
		header := w.Header()

		//	Allow Http Methods
		const allow = "POST, GET, OPTIONS, PUT, DELETE"
		header.Set("Access-Control-Allow-Methods", allow)

		//	Allow all Headers
		var allowHeaders = r.Header.Get("Access-Control-Allow-Headers")
		header.Set("Access-Control-Allow-Headers", allowHeaders)

		//	@TODO: Specifiy allow origin
		header.Set("Access-Control-Allow-Origin", "*")
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := httprouter.New()

	r.GlobalOPTIONS = http.HandlerFunc(cors)

	db := postgres.New("test")

	games.Serve(r, db)

	fmt.Printf("Server running at port: %s\n", port)

	http.ListenAndServe(port, r)
}

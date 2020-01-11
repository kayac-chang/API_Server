package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type send func(statusCode int, res interface{})

func Prepare(w http.ResponseWriter) send {

	return func(statusCode int, res interface{}) {
		w.WriteHeader(statusCode)

		sendJSON(w, res)
	}
}

func sendJSON(w http.ResponseWriter, data interface{}) {

	res, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s\n", res)
}

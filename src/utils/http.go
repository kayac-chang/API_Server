package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SendJSON(w http.ResponseWriter, data interface{}) {

	res, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s\n", res)
}

package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func sendJSON(w http.ResponseWriter, data interface{}) {
	res, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s\n", res)
}

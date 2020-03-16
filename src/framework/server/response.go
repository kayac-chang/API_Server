package server

import (
	"api/model/response"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func (it *Server) Send(w http.ResponseWriter, val interface{}) {

	switch val := val.(type) {

	case response.JSON:
		sendJSON(w, val)

	case response.ProtoBuf:
		sendProtoBuf(w, val)

	case string:
		w.Write([]byte(val))

	default:
		log.Fatalf("Unsupport Type: %t\n", val)
	}
}

func sendJSON(w http.ResponseWriter, res response.JSON) {

	output, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(output)
}

func sendProtoBuf(w http.ResponseWriter, res response.ProtoBuf) {

	fmt.Printf("%v\n", res)

	out, err := proto.Marshal(res.Data)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	w.Header().Set("Content-Type", "application/protobuf")
	w.WriteHeader(res.Code)
	w.Write(out)
}

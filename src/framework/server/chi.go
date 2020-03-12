package server

import (
	"api/model/response"

	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/golang/protobuf/proto"
)

type Server struct {
	*chi.Mux
}

func New() *Server {

	server := chi.NewRouter()

	server.Use(middleware.RequestID)
	server.Use(middleware.RealIP)
	server.Use(middleware.Logger)
	server.Use(middleware.Recoverer)

	return &Server{server}
}

func (it *Server) SendJSON(w http.ResponseWriter, data response.JSON) {

	output, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte("Serialization Error"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.Code)
	w.Write(output)
}

func (it *Server) SendProtoBuf(w http.ResponseWriter, res response.ProtoBuf) {

	out, err := proto.Marshal(res.Data)
	if err != nil {
		w.Write([]byte("Serialization Error"))
	}

	w.Header().Set("Content-Type", "application/protobuf")
	w.WriteHeader(res.Code)
	w.Write(out)
}

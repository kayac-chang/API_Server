package api

import "net/http"

type Handler interface {
	POST(w http.ResponseWriter, r *http.Request)
}

package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (it Server) URLParam(r *http.Request, key string) string {

	return chi.URLParam(r, key)
}

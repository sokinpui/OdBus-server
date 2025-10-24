package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"go-https-server/internal/handler"
)

func New() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/items", handler.PlaceholderGet).Methods(http.MethodGet)
	r.HandleFunc("/api/items", handler.PlaceholderPost).Methods(http.MethodPost)

	return r
}

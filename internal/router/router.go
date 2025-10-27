package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go-https-server/internal/handler"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("API called: %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func New(apiHandler *handler.ApiHandler) http.Handler {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/blocked_sign/qry", apiHandler.GetBlockedSigns).Methods(http.MethodGet)
	api.HandleFunc("/station_point/create", apiHandler.CreateStationPoint).Methods(http.MethodPost)
	api.HandleFunc("/station_point/qry", apiHandler.GetStationPoints).Methods(http.MethodGet)

	return r
}

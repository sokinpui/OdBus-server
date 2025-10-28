package router

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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

	// For development, allow all origins. In production, you should restrict this.
	corsOrigins := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "token", "dt"})

	r.Use(loggingMiddleware)

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/blockedSign/qry", apiHandler.GetBlockedSigns).Methods(http.MethodGet)
	api.HandleFunc("/stationPoint/create", apiHandler.CreateStationPoint).Methods(http.MethodPost)
	api.HandleFunc("/stationPoint/qry", apiHandler.GetStationPoints).Methods(http.MethodGet)

	// Wrap the router with the CORS middleware
	return handlers.CORS(corsOrigins, corsMethods, corsHeaders)(r)
}

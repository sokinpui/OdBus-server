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
	corsMethods := handlers.AllowedMethods([]string{"POST", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "token", "dt"})

	r.Use(loggingMiddleware)

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/blockedSign/qry", apiHandler.GetBlockedSigns).Methods(http.MethodPost)
	api.HandleFunc("/station/create", apiHandler.CreateStation).Methods(http.MethodPost)
	api.HandleFunc("/station/qry", apiHandler.GetStations).Methods(http.MethodPost)
	api.HandleFunc("/station/qryById", apiHandler.GetStationByID).Methods(http.MethodPost)
	api.HandleFunc("/station/update", apiHandler.UpdateStation).Methods(http.MethodPost)
	api.HandleFunc("/station/delete", apiHandler.DeleteStation).Methods(http.MethodPost)

	// Wrap the router with the CORS middleware
	return handlers.CORS(corsOrigins, corsMethods, corsHeaders)(r)
}

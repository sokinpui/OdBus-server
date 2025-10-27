package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"go-https-server/internal/handler"
)

func New(apiHandler *handler.ApiHandler) http.Handler {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/blocked_sign/qry", apiHandler.GetBlockedSigns).Methods(http.MethodGet)
	api.HandleFunc("/station_point/create", apiHandler.CreateStationPoint).Methods(http.MethodPost)
	api.HandleFunc("/station_point/qry", apiHandler.GetStationPoints).Methods(http.MethodGet)

	return r
}

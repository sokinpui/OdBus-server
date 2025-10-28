package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"go-https-server/internal/models"
	"go-https-server/internal/store"
)

// ApiHandler handles API requests.
type ApiHandler struct {
	store *store.Store
}

// NewApiHandler creates a new ApiHandler.
func NewApiHandler(s *store.Store) *ApiHandler {
	return &ApiHandler{store: s}
}

// StationCreateReq is the request DTO for creating a station.
type StationCreateReq struct {
	Name      string   `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Tags      []string `json:"tags"`
}

// GetBlockedSigns handles GET /api/blockedSign/qry
func (h *ApiHandler) GetBlockedSigns(w http.ResponseWriter, r *http.Request) {
	signs, err := h.store.GetBlockedSigns()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	respondWithJSON(w, http.StatusOK, signs)
}

// CreateStation handles POST /api/station/create
func (h *ApiHandler) CreateStation(w http.ResponseWriter, r *http.Request) {
	var req StationCreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	st := &models.Station{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Tags:      req.Tags,
	}

	if err := h.store.CreateStation(st); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	respondWithJSON(w, http.StatusCreated, st)
}

// GetStations handles GET /api/station/qry
func (h *ApiHandler) GetStations(w http.ResponseWriter, r *http.Request) {
	points, err := h.store.GetStations()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	respondWithJSON(w, http.StatusOK, points)
}

// GetStationByID handles GET /api/station/{id}
func (h *ApiHandler) GetStationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid station ID")
		return
	}

	station, err := h.store.GetStationByID(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if station == nil {
		respondWithError(w, http.StatusNotFound, "Station not found")
		return
	}

	respondWithJSON(w, http.StatusOK, station)
}

// StationUpdateReq is the request DTO for updating a station.
type StationUpdateReq struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

// UpdateStation handles PUT /api/station/{id}
func (h *ApiHandler) UpdateStation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid station ID")
		return
	}

	var req StationUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	st := &models.Station{
		ID:   id,
		Name: req.Name,
		Tags: req.Tags,
	}

	if err := h.store.UpdateStation(st); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	respondWithJSON(w, http.StatusOK, st)
}

// DeleteStation handles DELETE /api/station/{id}
func (h *ApiHandler) DeleteStation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid station ID")
		return
	}

	if err := h.store.DeleteStation(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Station deleted successfully"})
}

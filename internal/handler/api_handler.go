package handler

import (
	"encoding/json"
	"net/http"

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

// GetBlockedSigns handles POST /api/blockedSign/qry
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

// GetStations handles POST /api/station/qry
func (h *ApiHandler) GetStations(w http.ResponseWriter, r *http.Request) {
	points, err := h.store.GetStations()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	respondWithJSON(w, http.StatusOK, points)
}

// StationRequestByID is the request DTO for getting a station by ID.
type StationRequestByID struct {
	ID int `json:"id"`
}

// GetStationByID handles POST /api/station/qryById
func (h *ApiHandler) GetStationByID(w http.ResponseWriter, r *http.Request) {
	var req StationRequestByID
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	station, err := h.store.GetStationByID(req.ID)
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
	ID   int      `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

// UpdateStation handles POST /api/station/update
func (h *ApiHandler) UpdateStation(w http.ResponseWriter, r *http.Request) {
	var req StationUpdateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	st := &models.Station{
		ID:   req.ID,
		Name: req.Name,
		Tags: req.Tags,
	}

	if err := h.store.UpdateStation(st); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	respondWithJSON(w, http.StatusOK, st)
}

// StationDeleteReq is the request DTO for deleting a station.
type StationDeleteReq struct {
	ID int `json:"id"`
}

// DeleteStation handles POST /api/station/delete
func (h *ApiHandler) DeleteStation(w http.ResponseWriter, r *http.Request) {
	var req StationDeleteReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	if err := h.store.DeleteStation(req.ID); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Station deleted successfully"})
}

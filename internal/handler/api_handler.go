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

// StationPointCreateReq is the request DTO for creating a station point.
type StationPointCreateReq struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
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

// CreateStationPoint handles POST /api/stationPoint/create
func (h *ApiHandler) CreateStationPoint(w http.ResponseWriter, r *http.Request) {
	var req StationPointCreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	sp := &models.StationPoint{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	if err := h.store.CreateStationPoint(sp); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	respondWithJSON(w, http.StatusCreated, sp)
}

// GetStationPoints handles GET /api/stationPoint/qry
func (h *ApiHandler) GetStationPoints(w http.ResponseWriter, r *http.Request) {
	points, err := h.store.GetStationPoints()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	respondWithJSON(w, http.StatusOK, points)
}

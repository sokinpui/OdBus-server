package store

import (
	"database/sql"

	"go-https-server/internal/models"
)

// Store handles all database operations.
type Store struct {
	db *sql.DB
}

// New creates a new Store.
func New(db *sql.DB) *Store {
	return &Store{db: db}
}

// GetBlockedSigns retrieves all blocked signs from the database.
func (s *Store) GetBlockedSigns() ([]*models.BlockedSign, error) {
	rows, err := s.db.Query("SELECT id, ST_Y(location::geometry) AS latitude, ST_X(location::geometry) AS longitude FROM blockedSigns")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var signs []*models.BlockedSign
	for rows.Next() {
		var sign models.BlockedSign
		if err := rows.Scan(&sign.ID, &sign.Latitude, &sign.Longitude); err != nil {
			return nil, err
		}
		signs = append(signs, &sign)
	}
	return signs, nil
}

// CreateStationPoint inserts a new station point into the database.
func (s *Store) CreateStationPoint(sp *models.StationPoint) error {
	query := "INSERT INTO stationPoints (location) VALUES (ST_SetSRID(ST_MakePoint($1, $2), 4326)) RETURNING id"
	return s.db.QueryRow(query, sp.Longitude, sp.Latitude).Scan(&sp.ID)
}

// GetStationPoints retrieves all station points from the database.
func (s *Store) GetStationPoints() ([]*models.StationPoint, error) {
	rows, err := s.db.Query("SELECT id, ST_Y(location::geometry) AS latitude, ST_X(location::geometry) AS longitude FROM stationPoints")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var points []*models.StationPoint
	for rows.Next() {
		var point models.StationPoint
		if err := rows.Scan(&point.ID, &point.Latitude, &point.Longitude); err != nil {
			return nil, err
		}
		points = append(points, &point)
	}
	return points, nil
}

package store

import (
	"database/sql"
	"time"

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

	signs := make([]*models.BlockedSign, 0)
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
func (s *Store) CreateStation(st *models.Station) error {
	query := `
		INSERT INTO stations (name, location, created_by, is_active, tags)
		VALUES ($1, ST_SetSRID(ST_MakePoint($2, $3), 4326), $4, $5, $6)
		RETURNING id, created_at`
	// For simplicity, created_by is hardcoded. In a real app, this would come from auth.
	st.CreatedBy = "Current User"
	st.IsActive = true
	return s.db.QueryRow(query, st.Name, st.Longitude, st.Latitude, st.CreatedBy, st.IsActive, st.Tags).Scan(&st.ID, &st.CreatedAt)
}

// GetStationPoints retrieves all station points from the database.
func (s *Store) GetStations() ([]*models.Station, error) {
	rows, err := s.db.Query("SELECT id, name, ST_Y(location::geometry) AS latitude, ST_X(location::geometry) AS longitude, created_by, created_at, updated_at, is_active, tags FROM stations ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stations := make([]*models.Station, 0)
	for rows.Next() {
		var station models.Station
		if err := rows.Scan(&station.ID, &station.Name, &station.Latitude, &station.Longitude, &station.CreatedBy, &station.CreatedAt, &station.UpdatedAt, &station.IsActive, &station.Tags); err != nil {
			return nil, err
		}
		stations = append(stations, &station)
	}
	return stations, nil
}

// GetStationByID retrieves a single station by its ID.
func (s *Store) GetStationByID(id int) (*models.Station, error) {
	var station models.Station
	query := "SELECT id, name, ST_Y(location::geometry) AS latitude, ST_X(location::geometry) AS longitude, created_by, created_at, updated_at, is_active, tags FROM stations WHERE id = $1"
	err := s.db.QueryRow(query, id).Scan(&station.ID, &station.Name, &station.Latitude, &station.Longitude, &station.CreatedBy, &station.CreatedAt, &station.UpdatedAt, &station.IsActive, &station.Tags)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found
		}
		return nil, err
	}
	return &station, nil
}

// DeleteStation deletes a station from the database by its ID.
func (s *Store) DeleteStation(id int) error {
	query := "DELETE FROM stations WHERE id = $1"
	_, err := s.db.Exec(query, id)
	return err
}

// UpdateStation updates an existing station in the database.
func (s *Store) UpdateStation(st *models.Station) error {
	query := `
		UPDATE stations
		SET name = $1, tags = $2, updated_at = $3
		WHERE id = $4
		RETURNING id, name, ST_Y(location::geometry) AS latitude, ST_X(location::geometry) AS longitude, created_by, created_at, updated_at, is_active, tags`
	return s.db.QueryRow(query, st.Name, st.Tags, time.Now(), st.ID).Scan(&st.ID, &st.Name, &st.Latitude, &st.Longitude, &st.CreatedBy, &st.CreatedAt, &st.UpdatedAt, &st.IsActive, &st.Tags)
}

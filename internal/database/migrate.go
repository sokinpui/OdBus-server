package database

import "database/sql"

// Migrate creates the necessary tables in the database if they don't exist.
func Migrate(db *sql.DB) error {
	createBlockedSignsTable := `
	CREATE TABLE IF NOT EXISTS blocked_signs (
		id SERIAL PRIMARY KEY,
		latitude REAL NOT NULL,
		longitude REAL NOT NULL
	);`

	if _, err := db.Exec(createBlockedSignsTable); err != nil {
		return err
	}

	createStationPointsTable := `
	CREATE TABLE IF NOT EXISTS station_points (
		id SERIAL PRIMARY KEY,
		latitude REAL NOT NULL,
		longitude REAL NOT NULL
	);`

	if _, err := db.Exec(createStationPointsTable); err != nil {
		return err
	}

	return nil
}

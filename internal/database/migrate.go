package database

import "database/sql"

// Migrate creates the necessary tables in the database if they don't exist.
func Migrate(db *sql.DB) error {
	if _, err := db.Exec("CREATE EXTENSION IF NOT EXISTS postgis"); err != nil {
		return err
	}

	createBlockedSignsTable := `
	CREATE TABLE IF NOT EXISTS blocked_signs (
		id SERIAL PRIMARY KEY,
		location GEOGRAPHY(Point, 4326) NOT NULL
	);`

	if _, err := db.Exec(createBlockedSignsTable); err != nil {
		return err
	}

	createStationPointsTable := `
	CREATE TABLE IF NOT EXISTS station_points (
		id SERIAL PRIMARY KEY,
		location GEOGRAPHY(Point, 4326) NOT NULL
	);`

	if _, err := db.Exec(createStationPointsTable); err != nil {
		return err
	}

	return nil
}

package database

import "database/sql"

// Migrate creates the necessary tables in the database if they don't exist.
func Migrate(db *sql.DB) error {
	if _, err := db.Exec("CREATE EXTENSION IF NOT EXISTS postgis"); err != nil {
		return err
	}

	createBlockedSignsTable := `
	CREATE TABLE IF NOT EXISTS blockedSigns (
		id SERIAL PRIMARY KEY,
		location GEOGRAPHY(Point, 4326) NOT NULL
	);`

	if _, err := db.Exec(createBlockedSignsTable); err != nil {
		return err
	}

	createStationsTable := `
	CREATE TABLE IF NOT EXISTS stations (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		location GEOGRAPHY(Point, 4326) NOT NULL,
		"createdBy" VARCHAR(255) NOT NULL,
		"createdAt" TIMESTAMPTZ DEFAULT NOW(),
		"updatedAt" TIMESTAMPTZ,
		"isActive" BOOLEAN DEFAULT TRUE,
		tags TEXT[]
	);`
	if _, err := db.Exec(createStationsTable); err != nil {
		return err
	}

	return nil
}

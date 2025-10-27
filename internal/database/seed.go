package database

import (
	"database/sql"
	"fmt"
	"log"

	"go-https-server/internal/kml"
)

// SeedBlockedSigns populates the blockedSigns table from a KMZ file if the table is empty.
func SeedBlockedSigns(db *sql.DB, kmzPath string) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM blockedSigns").Scan(&count)
	if err != nil {
		return fmt.Errorf("could not query blockedSigns count: %w", err)
	}

	if count > 0 {
		log.Println("blockedSigns table already seeded")
		return nil
	}

	log.Printf("seeding data from %s", kmzPath)

	latLongs, err := kml.ParseKMZ(kmzPath)
	if err != nil {
		return fmt.Errorf("could not parse KMZ file: %w", err)
	}

	txn, err := db.Begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare("INSERT INTO blockedSigns (location) VALUES (ST_SetSRID(ST_MakePoint($1, $2), 4326))")
	if err != nil {
		return fmt.Errorf("could not prepare statement: %w", err)
	}
	defer stmt.Close()

	log.Println("inserting records into blockedSigns table...")

	for _, ll := range latLongs {
		if _, err := stmt.Exec(ll.Longitude, ll.Latitude); err != nil {
			return fmt.Errorf("could not execute statement: %w", err)
		}
	}

	log.Printf("seeded %d records into blockedSigns table", len(latLongs))

	return txn.Commit()
}

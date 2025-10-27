package database

import (
	"database/sql"
	"fmt"
	"log"

	"go-https-server/internal/kml"
)

// SeedBlockedSigns populates the blocked_signs table from a KMZ file if the table is empty.
func SeedBlockedSigns(db *sql.DB, kmzPath string) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM blocked_signs").Scan(&count)
	if err != nil {
		return fmt.Errorf("could not query blocked_signs count: %w", err)
	}

	if count > 0 {
		log.Println("blocked_signs table already seeded")
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

	stmt, err := txn.Prepare("INSERT INTO blocked_signs (latitude, longitude) VALUES ($1, $2)")
	if err != nil {
		return fmt.Errorf("could not prepare statement: %w", err)
	}
	defer stmt.Close()

	log.Println("inserting records into blocked_signs table...")

	for _, ll := range latLongs {
		if _, err := stmt.Exec(ll.Latitude, ll.Longitude); err != nil {
			return fmt.Errorf("could not execute statement: %w", err)
		}
	}

	log.Printf("seeded %d records into blocked_signs table", len(latLongs))

	return txn.Commit()
}

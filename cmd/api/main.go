package main

import (
	"log"
	"net/http"

	"go-https-server/internal/config"
	"go-https-server/internal/database"
	"go-https-server/internal/handler"
	"go-https-server/internal/logger"
	"go-https-server/internal/router"
	"go-https-server/internal/server"
	"go-https-server/internal/store"
)

func main() {
	logger.Init()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping database: %v", err)
	}

	log.Println("database connection successful")

	if err := database.Migrate(db); err != nil {
		log.Fatalf("could not migrate database: %v", err)
	}
	log.Println("database migration successful")

	if err := database.SeedBlockedSigns(db, "No_public_light_buses_labels.kmz"); err != nil {
		log.Fatalf("could not seed blocked signs data: %v", err)
	}

	s := store.New(db)
	apiHandler := handler.NewApiHandler(s)

	r := router.New(apiHandler)
	srv := server.New(cfg.ServerAddr, r)

	log.Printf("starting server on %s", cfg.ServerAddr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not start server: %v", err)
	}
}

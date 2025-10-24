package main

import (
	"log"
	"net/http"

	"go-https-server/internal/config"
	"go-https-server/internal/database"
	"go-https-server/internal/router"
	"go-https-server/internal/server"
)

func main() {
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

	r := router.New()
	srv := server.New(cfg.ServerAddr, r)

	log.Printf("starting server on %s", cfg.ServerAddr)
	if err := srv.ListenAndServeTLS("cert.pem", "key.pem"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not start server: %v", err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"go-https-server/internal/config"
	"go-https-server/internal/database"
	"go-https-server/internal/logger"
	"go-https-server/internal/models"
	"go-https-server/internal/store"
)

const (
	kmbRouteStopAPI = "https://data.etabus.gov.hk/v1/transport/kmb/route-stop/%s/%s/1"
	kmbStopAPI      = "https://data.etabus.gov.hk/v1/transport/kmb/stop/%s"
)

type RouteStopResponse struct {
	Data []RouteStop `json:"data"`
}

type RouteStop struct {
	Stop string `json:"stop"`
}

type StopResponse struct {
	Data Stop `json:"data"`
}

type Stop struct {
	NameEn string `json:"name_en"`
	Lat    string `json:"lat"`
	Long   string `json:"long"`
}

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

	s := store.New(db)

	route := "63X"
	directions := []string{"outbound", "inbound"}

	for _, direction := range directions {
		if err := fetchAndSeedStations(s, route, direction); err != nil {
			log.Fatalf("could not seed stations for route %s %s: %v", route, direction, err)
		}
	}

	log.Println("successfully seeded all stations")
}

func fetchAndSeedStations(s *store.Store, route string, direction string) error {
	log.Printf("fetching stations for route %s, direction %s", route, direction)

	routeStopURL := fmt.Sprintf(kmbRouteStopAPI, route, direction)
	resp, err := http.Get(routeStopURL)
	if err != nil {
		return fmt.Errorf("could not fetch route stops from %s: %w", routeStopURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code from %s: %d", routeStopURL, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %w", err)
	}

	var routeStopResponse RouteStopResponse
	if err := json.Unmarshal(body, &routeStopResponse); err != nil {
		return fmt.Errorf("could not unmarshal route stop response: %w", err)
	}

	log.Printf("found %d stops for route %s, direction %s", len(routeStopResponse.Data), route, direction)

	for _, routeStop := range routeStopResponse.Data {
		if err := processStop(s, routeStop.Stop, route, direction); err != nil {
			log.Printf("could not process stop %s: %v. skipping.", routeStop.Stop, err)
		}
		// Add a small delay to avoid hitting API rate limits.
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func processStop(s *store.Store, stopID, route, direction string) error {
	stopURL := fmt.Sprintf(kmbStopAPI, stopID)
	resp, err := http.Get(stopURL)
	if err != nil {
		return fmt.Errorf("could not fetch stop details from %s: %w", stopURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code from %s: %d", stopURL, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read stop response body: %w", err)
	}

	var stopResponse StopResponse
	if err := json.Unmarshal(body, &stopResponse); err != nil {
		return fmt.Errorf("could not unmarshal stop response: %w", err)
	}

	lat, err := strconv.ParseFloat(stopResponse.Data.Lat, 64)
	if err != nil {
		return fmt.Errorf("could not parse latitude '%s': %w", stopResponse.Data.Lat, err)
	}

	long, err := strconv.ParseFloat(stopResponse.Data.Long, 64)
	if err != nil {
		return fmt.Errorf("could not parse longitude '%s': %w", stopResponse.Data.Long, err)
	}

	station := &models.Station{
		Name:      stopResponse.Data.NameEn,
		Latitude:  lat,
		Longitude: long,
		Tags:      []string{"kmb", route, direction},
	}

	if err := s.CreateStation(station); err != nil {
		return fmt.Errorf("could not create station '%s': %w", station.Name, err)
	}

	log.Printf("successfully inserted station: %s", station.Name)
	return nil
}

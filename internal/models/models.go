package models

// BlockedSign represents a blocked sign location.
type BlockedSign struct {
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// StationPoint represents a station point location.
type StationPoint struct {
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

package models

import (
	"github.com/lib/pq"
	"time"
)

// BlockedSign represents a blocked sign location.
type BlockedSign struct {
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Station represents a station point location.
type Station struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Latitude  float64        `json:"latitude"`
	Longitude float64        `json:"longitude"`
	CreatedBy string         `json:"createdBy"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt *time.Time     `json:"updatedAt,omitempty"`
	IsActive  bool           `json:"isActive"`
	Tags      pq.StringArray `json:"tags"`
}

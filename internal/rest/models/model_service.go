package models

import (
	"time"
)

type Service struct {
	Id string `json:"id,omitempty"`
	// Time the sevice was created
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Time the sevice was updated
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Time the sevice was deleted
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Unique Title of the service
	Title string `json:"title,omitempty"`
	// description of the service
	Description string `json:"description,omitempty"`
	// List of the services versions
	Versions []ServiceVersion `json:"versions,omitempty"`
	// Number of versions the service has
	VersionCount float64 `json:"version_count,omitempty"`
}

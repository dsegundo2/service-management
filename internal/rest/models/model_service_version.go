package models

import (
	"time"
)

type ServiceVersion struct {
	// version number of this service version. must follow number.number.number format
	Version string `json:"version,omitempty"`
	// parent id that the service version belongs to
	ServiceId string    `json:"service_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// proof of concept random field on a version
	ServiceExampleField string `json:"service_example_field,omitempty"`
}

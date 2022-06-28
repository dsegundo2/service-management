package servicedb

import (
	"time"

	"gorm.io/gorm"
)

type Service struct {
	ID string `json:"id"`
	TimeModel

	// Service Level feilds
	Title       string `gorm:"unique" json:"title"`
	Description string `json:"description"`

	// Reference to Versions
	Versions     []*ServiceVersion `gorm:"foreignKey:ServiceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"versions"`
	VersionCount int64             `gorm:"-:all" json:"version_count"`
}

type ServiceVersion struct {
	TimeModel
	// Composite Primary Key with Version & Parent Service ID
	Version   string `gorm:"primaryKey"`
	ServiceID string `gorm:"primaryKey" json:"service_id"`

	// Service Specific fields. Other fields to go here
	ServiceExampleField string `json:"service_example_field"`
}

// Similar to Gorm.Model but without the uint ID
type TimeModel struct {
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

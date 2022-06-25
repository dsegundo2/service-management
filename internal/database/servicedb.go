package servicedb

import (
	"time"

	"gorm.io/gorm"
)

//
type ServiceDbApi interface {
	CreateService(service *Service) error
	ReadService(service *Service) error
	UpdateService(service *Service) error

	// Hard Delete a service from the database
	DeleteService(service *Service) error
}

type ServiceDB struct {
	*gorm.DB
}

func New(db *gorm.DB) *ServiceDB {
	return &ServiceDB{db}
}

func (db *ServiceDB) Migrate() error {
	return db.AutoMigrate(
		&Service{},
		&ServiceInstance{},
	)
}

type Service struct {
	gorm.Model
	Versions []*ServiceInstance `gorm:"foreignKey:ServiceID" json:"versions"`
}

type ServiceInstance struct {
	// Composite Primary Key with Version & Parent Service ID
	Version   string `gorm:"primaryKey"`
	ServiceID string `gorm:"primaryKey" json:"service_id"`

	// Time Info automatically updated by Gorm
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Service Specific fields
	Title       string `gorm:"unique" json:"title"`
	Description string `json:"description"`
}

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
	DeleteService(serviceID string) error

	/*	Service Versions	*/
	// Creates a version of a service, valid ServiceID required
	CreateServiceVersion(version *ServiceInstance) error
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
	ID string `json:"id"`
	TimeModel

	// Service Level feilds
	Title       string `gorm:"unique" json:"title"`
	Description string `json:"description"`

	// Reference to Versions
	Versions []*ServiceInstance `gorm:"foreignKey:ServiceID,OnDelete:CASCADE;" json:"versions"`
}

type ServiceInstance struct {
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

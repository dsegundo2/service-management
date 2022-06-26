package servicedb

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (db *ServiceDB) CreateService(service *Service) error {
	return db.Create(service).Error
}

func (db *ServiceDB) ReadService(service *Service) error {
	return db.Preload(clause.Associations).First(service).Error
}

func (db *ServiceDB) UpdateService(service *Service) error {
	return db.Model(service).Select("title", "description").Updates(service).Error
}

// Hard Delete a service from the database
func (db *ServiceDB) DeleteService(serviceID string) error {
	return db.Delete(serviceID).Error
}

// Creates a version of a service, valid ServiceID required
func (db *ServiceDB) CreateServiceVersion(version *ServiceInstance) error {
	return db.Create(version).Error
}

/*		DB Hooks	*/
// Create a uuid for services without ids
func (s *Service) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == "" {
		s.ID = uuid.NewString()
	}
	return
}

func (si *ServiceInstance) BeforeCreate(tx *gorm.DB) (err error) {
	// Regex ^(\d+\.)?(\d+\.)?(\*|\d+)$ to check for 0.0.0 formatting on version
	matched, err := regexp.MatchString("^(\\d+\\.)?(\\d+\\.)?(\\*|\\d+)$", si.Version)
	if !matched || err != nil {
		return fmt.Errorf("version format not supported, must follow 0.0.0 semantic")
	}
	return
}

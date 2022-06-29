package servicedb

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*		Service Methods		*/

func (db *ServiceDB) CreateService(service *Service) error {
	return db.Create(service).Error
}

// Returns a list of services. Filter used to look for only services with the title that has a substring filter
// Limit and offset used for pagination. Results Ordered using sort string. Total matched queries returned
func (db *ServiceDB) ListServices(filter, sort string, limit, offset int) (serviceList []*Service, totalCount int64, err error) {
	err = db.Model(&Service{}).Where("Title LIKE ?", fmt.Sprintf("%%%s%%", filter)).Count(&totalCount).Offset(offset).Limit(limit).Order(sort).Find(&serviceList).Error
	logrus.WithError(err).WithFields(logrus.Fields{
		"Count": totalCount,
		"List":  serviceList,
	}).Debug("List services query")
	return
}

// Reads a service from the database. Loads the Versions if preload is true
func (db *ServiceDB) ReadService(service *Service, preload bool) error {
	if service.ID == "" {
		return fmt.Errorf("service not found")
	}
	var err error
	if preload {
		err = db.Preload(clause.Associations).First(service).Error
	} else {
		err = db.First(service).Error
	}
	if err != nil {
		return err
	}

	return nil
}

func (db *ServiceDB) UpdateService(service *Service) error {
	return db.Model(service).Select("title", "description").Updates(service).Error
}

// Hard Delete a service from the database
func (db *ServiceDB) DeleteService(serviceID string) error {
	return db.Unscoped().Delete(&Service{ID: serviceID}).Error
}

/*		Service Versions		*/

// Creates a version of a service, valid ServiceID required
func (db *ServiceDB) CreateServiceVersion(version *ServiceVersion) error {
	return db.Create(version).Error

}

func (db *ServiceDB) ReadServiceVersion(version *ServiceVersion) error {
	if version.ServiceID == "" || version.Version == "" {
		return fmt.Errorf("service not found")
	}
	return db.First(version).Error
}

func (db *ServiceDB) UpdateServiceVersion(version *ServiceVersion) error {
	return db.Model(version).Select("version", "service_id", "service_example_field").Updates(version).Error
}

func (db *ServiceDB) DeleteServiceVersion(serviceID, version string) error {
	return db.Unscoped().Delete(&ServiceVersion{
		Version:   version,
		ServiceID: serviceID,
	}).Error
}

/*			DB Hooks			*/
// Create a uuid for services without ids
func (s *Service) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == "" {
		s.ID = uuid.NewString()
	}
	return
}

// Add the version count when loading a service
func (s *Service) AfterFind(tx *gorm.DB) (err error) {
	association := tx.Model(&s).Association("Versions")
	if association.Error != nil {
		return association.Error
	}
	s.VersionCount = association.Count()
	return
}

func (si *ServiceVersion) BeforeCreate(tx *gorm.DB) (err error) {
	// Regex ^(\d+\.)?(\d+\.)?(\*|\d+)$ to check for 0.0.0 formatting on version
	matched, err := regexp.MatchString("^(\\d+\\.)?(\\d+\\.)?(\\*|\\d+)$", si.Version)
	if !matched || err != nil {
		return fmt.Errorf("version format not supported, must follow 0.0.0 semantic")
	}
	return
}

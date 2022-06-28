package servicedb

import (
	"gorm.io/gorm"
)

//
type ServiceDbApi interface {
	CreateService(service *Service) error
	ReadService(service *Service, preload bool) error
	UpdateService(service *Service) error
	ListServices(filter, sort string, limit, offset int) (serviceList []*Service, totalCount int64, err error)
	// Hard Delete a service from the database
	DeleteService(serviceID string) error

	/*	Service Versions	*/

	// Creates a version of a service, valid ServiceID required
	CreateServiceVersion(version *ServiceVersion) error
	ReadServiceVersion(version *ServiceVersion) error
	UpdateServiceVersion(version *ServiceVersion) error
	DeleteServiceVersion(serviceID, version string) error
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
		&ServiceVersion{},
	)
}

//go:build integration

package servicedb

import (
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	testLog = logrus.NewEntry(logrus.StandardLogger())

	testDB           *ServiceDB
	servicesToDelete = []string{} // IDs of the services to delete

	testServiceID string
)

func Test_Service_Database(t *testing.T) {
	testLog.Logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	// Setup & migrate database
	dbUrl := os.Getenv("SERVICE_MANAGEMENT_DB_CONNECTION_URL")

	postgresDB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		testLog.WithField("Connection String", dbUrl).
			Fatal("could not create connect to database. make sure to export a SERVICE_MANAGEMENT_DB_CONNECTION_URL variable")
	}
	testDB = New(postgresDB)

	// Defer closing the database
	db, err := postgresDB.DB()
	if err != nil {
		testLog.WithField("Connection String", dbUrl).Fatal("could not get sql db")
	}
	defer db.Close()

	// Test Service
	t.Run("MigrateDatabase", assertMigration)
	t.Cleanup(deleteServices)
	t.Run("ServiceCRUD", testServiceFunctions)

	// Test Versions

}

func assertMigration(t *testing.T) {
	err := testDB.Migrate()
	require.NoError(t, err)
}

// Tests Create, Read, Update functionality on the service and sets the testServiceID field
func testServiceFunctions(t *testing.T) {
	/*	Create Service	*/
	serviceTitle := "LocateUsUnique-" + uuid.NewString()
	serviceDescription := "A service that provides support for locating the business"
	createService := &Service{
		Title:       serviceTitle,
		Description: serviceDescription,
	}
	err := testDB.CreateService(createService)
	testLog.WithField("CreatedService", createService).Info("created service")
	require.NoError(t, err)
	require.NotEmpty(t, createService.ID)
	require.NotEqual(t, time.Time{}, createService.CreatedAt)
	require.NotEqual(t, time.Time{}, createService.UpdatedAt)
	require.Nil(t, createService.DeletedAt)
	require.Equal(t, serviceTitle, createService.Title)
	require.Equal(t, serviceDescription, createService.Description)
	testServiceID = createService.ID

	/*	Update Service	*/
	updateDescription := "this is a new description"
	updateService := &Service{
		ID:          testServiceID,
		Title:       serviceTitle,
		Description: updateDescription,
	}
	err = testDB.UpdateService(updateService)
	testLog.WithField("UpdateService", updateService).Info("updated service")
	require.NoError(t, err)
	require.True(t, updateService.UpdatedAt.After(createService.UpdatedAt), "UpdatedAt should be changed")
	require.Equal(t, updateDescription, updateService.Description)
	require.Equal(t, serviceTitle, updateService.Title)

	/*	Read Service	*/
	getService := &Service{
		ID: createService.ID,
	}
	err = testDB.ReadService(getService)
	testLog.WithField("ReadService", getService).Info("read service")
	require.NoError(t, err)
	require.Equal(t, createService.CreatedAt, getService.CreatedAt)
	require.Equal(t, updateService.UpdatedAt, getService.UpdatedAt)
	require.Nil(t, getService.DeletedAt)
	require.Equal(t, serviceTitle, getService.Title)
	require.Equal(t, updateDescription, getService.Description)
}

// Delete the test Services from the database
func deleteServices() {
	for _, id := range servicesToDelete {
		if err := testDB.DeleteService(id); err != nil {
			testLog.WithField("Service ID", id).Warn("error when cleaning up test data and deleting service")
		}
	}
}

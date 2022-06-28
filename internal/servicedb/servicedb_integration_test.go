//go:build integration

package servicedb

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
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

	testServiceID string

	// For Clean up function
	dbToClose        *sql.DB
	testDB           *ServiceDB
	servicesToDelete = []string{}
)

func Test_Service_Database(t *testing.T) {
	testLog.Logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	// testLog.Logger.SetLevel()

	// Setup & migrate database
	dbUrl := os.Getenv("SERVICE_MANAGEMENT_DB_CONNECTION_URL")

	postgresDB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		FullSaveAssociations: true,
		PrepareStmt:          true,
	})
	if err != nil {
		testLog.WithField("Connection String", dbUrl).
			Fatal("could not create connect to database. make sure to export a SERVICE_MANAGEMENT_DB_CONNECTION_URL variable")
	}
	testDB = New(postgresDB)

	// Defer closing the database
	dbToClose, err = postgresDB.DB()
	if err != nil {
		testLog.WithField("Connection String", dbUrl).Fatal("could not get sql db")
	}

	// Test Service
	t.Run("MigrateDatabase", assertMigration)
	t.Cleanup(deleteServices)
	t.Run("ServiceCRUD", testServiceFunctions)
	t.Run("ServiceVersionCRUD", testServiceVersionCRUDFunctions)
	t.Run("ListServicesTest", testListServices)
}

func assertMigration(t *testing.T) {
	err := testDB.Migrate()
	require.NoError(t, err)
}

// Tests Create, Read, Update functionality on the service and sets the testServiceID field
func testServiceFunctions(t *testing.T) {
	/*	Create Service	*/

	// Successful Create
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

	// Check for same name error
	err = testDB.CreateService(createService)
	require.Error(t, err)

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
	err = testDB.ReadService(getService, false)
	testLog.WithField("ReadService", getService).Info("read service")
	require.NoError(t, err)
	require.Equal(t, createService.CreatedAt, getService.CreatedAt)
	require.Equal(t, updateService.UpdatedAt, getService.UpdatedAt)
	require.Nil(t, getService.DeletedAt)
	require.Equal(t, serviceTitle, getService.Title)
	require.Equal(t, updateDescription, getService.Description)

	// Check for read error for empty and fake id
	x := &Service{
		ID: "NOT A REAL ID",
	}
	err = testDB.ReadService(x, false)
	require.Error(t, err)
	err = testDB.ReadService(&Service{}, true)
	require.Error(t, err)
}

func testServiceVersionCRUDFunctions(t *testing.T) {
	/*		CreateServiceVersion Checks		*/
	firstVersion := "0.0.1"
	firstVersionField := "same field 1"
	// Create 2 Versions
	createServiceVersion1 := &ServiceVersion{
		Version:             firstVersion,
		ServiceID:           testServiceID,
		ServiceExampleField: firstVersionField,
	}
	err := testDB.CreateServiceVersion(createServiceVersion1)
	testLog.WithField("CreateServiceVersion1", createServiceVersion1).Info("CreateServiceVersion response")
	require.NoError(t, err)

	createServiceVersion2 := &ServiceVersion{
		Version:             "0.0.2",
		ServiceID:           testServiceID,
		ServiceExampleField: "same field 2",
	}
	err = testDB.CreateServiceVersion(createServiceVersion2)
	require.NoError(t, err)

	// Test Error Create situations
	err = testDB.CreateServiceVersion(&ServiceVersion{
		Version:   "BADVERSION",
		ServiceID: testServiceID,
	})
	require.Error(t, err, "expect error on invalid version semantics")

	err = testDB.CreateServiceVersion(&ServiceVersion{
		Version:   "0.0.5",
		ServiceID: "FAKE SERVICE",
	})
	require.Error(t, err, "expect error on bad service id")

	/*		UpdateServiceVersion Checks		*/
	updatedServiceField := "this is an updated value"
	updateVersion := &ServiceVersion{
		Version:             firstVersion,
		ServiceID:           testServiceID,
		ServiceExampleField: updatedServiceField,
	}
	err = testDB.UpdateServiceVersion(updateVersion)
	testLog.WithField("UpdateVersion", updateVersion).Info("UpdateServiceVersion response")
	require.NoError(t, err)
	require.True(t, updateVersion.UpdatedAt.After(createServiceVersion1.UpdatedAt))
	require.Equal(t, updatedServiceField, updateVersion.ServiceExampleField)

	// Check for error
	err = testDB.UpdateServiceVersion(&ServiceVersion{})
	require.Error(t, err, "should have an error trying to update empty obj")

	/*		ReadServiceVersion Checks		*/
	getVersion := &ServiceVersion{
		Version:   firstVersion,
		ServiceID: testServiceID,
	}
	err = testDB.ReadServiceVersion(getVersion)
	testLog.WithField("GetVersion", getVersion).Info("ReadServiceVersion response")
	require.NoError(t, err)
	require.Equal(t, updatedServiceField, getVersion.ServiceExampleField)

	// Check for expectde read errors
	x := &ServiceVersion{}
	err = testDB.ReadServiceVersion(x)
	testLog.WithField("GetVersion", x).Info("ReadServiceVersion response")

	require.Error(t, err)
	err = testDB.ReadServiceVersion(&ServiceVersion{
		Version:   firstVersion,
		ServiceID: "bad service id",
	})
	require.Error(t, err)

	// Check for version count on Service
	parentService := &Service{
		ID: testServiceID,
	}
	err = testDB.ReadService(parentService, false)
	testLog.WithField("Service", parentService).Info("Read Service response")
	require.NoError(t, err)
	require.Len(t, parentService.Versions, 0, "should not load versions")
	require.EqualValues(t, 2, parentService.VersionCount, "should have 2 services")

	/*		DeleteServiceVersion Checks		*/
	err = testDB.DeleteServiceVersion(testServiceID, firstVersion)
	require.NoError(t, err)

	// Check that the count went down 1
	err = testDB.ReadService(parentService, true)
	testLog.WithField("Service", parentService).Info("Read Service response after deleting version")
	require.NoError(t, err)
	require.EqualValues(t, 1, parentService.VersionCount, "should have 1 service now")
	require.Len(t, parentService.Versions, 1)
	require.NotEqual(t, parentService.Versions[0].Version, firstVersion)
}

func testListServices(t *testing.T) {
	/*	 Populate DB with Services		*/

	// Successful Create multiple
	titleSubString := uuid.NewString()
	for i := 0; i < 10; i++ {
		createService := &Service{
			Title:       "Service Title " + titleSubString + uuid.NewString(),
			Description: "ServiceDescription",
		}
		err := testDB.CreateService(createService)
		require.NoError(t, err)
		servicesToDelete = append(servicesToDelete, createService.ID)
	}
	for i := 0; i < 5; i++ {
		createService := &Service{
			Title:       "Service Title2 " + uuid.NewString(),
			Description: "ServiceDescription",
		}
		err := testDB.CreateService(createService)
		require.NoError(t, err)
		servicesToDelete = append(servicesToDelete, createService.ID)

		err = testDB.CreateServiceVersion(&ServiceVersion{
			Version:             "0.0." + fmt.Sprint(i),
			ServiceID:           createService.ID,
			ServiceExampleField: "sample fieldd",
		})
	}

	/*	List Services with filter, sorting and pagination	*/
	t.Run("TestPagination", func(t *testing.T) {
		list, count, err := testDB.ListServices("", "", 10, 0)
		require.NoError(t, err)
		require.NotEmpty(t, list)
		require.LessOrEqual(t, int64(15), count)
		require.Len(t, list, 10)
	})
	t.Run("TestStringFilterSearch", func(t *testing.T) {
		list, count, err := testDB.ListServices(titleSubString, "", 8, 0)
		require.NoError(t, err)
		require.NotEmpty(t, list)
		require.EqualValues(t, 10, count)
		require.Len(t, list, 8)
	})
	t.Run("TestSortingByTitle", func(t *testing.T) {
		list, _, err := testDB.ListServices("", "title desc", 15, 0)
		require.NoError(t, err)
		require.NotEmpty(t, list)
		require.Len(t, list, 15)
		// Check for descending order
		for i := 0; i < len(list)-1; i++ {
			require.GreaterOrEqual(t, strings.ToLower(list[i].Title), strings.ToLower(list[i+1].Title))
		}

		// Check Ascending
		list, _, err = testDB.ListServices("", "title", 15, 0)
		require.NoError(t, err)
		require.NotEmpty(t, list)
		require.Len(t, list, 15)
		// Check for descending order
		for i := 0; i < len(list)-1; i++ {
			require.LessOrEqual(t, strings.ToLower(list[i].Title), strings.ToLower(list[i+1].Title))
		}
	})

}

// Delete the test Services from the database
func deleteServices() {
	defer dbToClose.Close()
	for _, id := range servicesToDelete {
		if err := testDB.DeleteService(id); err != nil {
			testLog.WithError(err).WithField("Service ID", id).Warn("error when cleaning up test data and deleting service")
		}
	}
}

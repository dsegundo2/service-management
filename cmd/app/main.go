package main

import (
	"os"

	"github.com/dsegundo2/service-management/internal/rest"
	"github.com/dsegundo2/service-management/internal/servicedb"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Setup & migrate database
	dbUrl := os.Getenv("SERVICE_MANAGEMENT_DB_CONNECTION_URL")

	postgresDB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		logrus.WithField("Connection String", dbUrl).
			Fatal("could not create connect to database. make sure to export a SERVICE_MANAGEMENT_DB_CONNECTION_URL variable")
	}
	serviceDatabase := servicedb.New(postgresDB)

	// Defer closing the database
	db, err := postgresDB.DB()
	if err != nil {
		logrus.WithField("Connection String", dbUrl).Fatal("could not get sql db")
	}
	defer db.Close()

	if err := serviceDatabase.Migrate(); err != nil {
		logrus.WithField("Connection String", dbUrl).Fatal("error when migrating database")
	}

	// Create Logget
	log := logrus.NewEntry(logrus.StandardLogger())
	log.Logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	// Create Router
	addr := os.Getenv("SERVICE_MANAGEMENT_ADDRESS")
	if addr == "" {
		log.Fatal("SERVICE_MANAGEMENT_ADDRESS required as environment variable")
	}
	server := rest.New(addr, serviceDatabase, log)

	server.Listen()
}

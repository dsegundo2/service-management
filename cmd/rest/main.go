package main

import (
	"os"

	"github.com/dsegundo2/service-management/internal/database/servicedb"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Setup & migrate database
	dbUrl := os.Getenv("SERVICE_MANAGEMENT_DB_CONNECTION_URL")

	postgresDB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		logrus.WithField("Connection String", dbUrl).Fatal("could not create connect to database")
	}
	serviceDatabase := servicedb.New(postgresDB)

	// Defer closing the database
	db, err := postgresDB.DB()
	if err != nil {
		logrus.WithField("Connection String", dbUrl).Fatal("could not get sql db")
	}
	defer db.Close()

	serviceDatabase.AutoMigrate()
}

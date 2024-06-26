package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// define the database as a global variable
var DB *gorm.DB

// connect to the database
func ConnectDb() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("store.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Connected to Database!")

	// Migrate the schema
	// Need to add all models to the migrate function after created
	err = DB.AutoMigrate()
	if err != nil {
		return err
	}

	return nil
}

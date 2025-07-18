package db

import (
	calculationserver "Calculator/interal/calculationServer"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=qwerty dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database %v", err)
	}
	if err := db.AutoMigrate(&calculationserver.Calculation{}); err != nil {
		log.Fatalf("Could not connect to database %v", err)
	}
	return db, nil
}

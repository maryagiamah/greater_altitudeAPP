package config

import (
	"github.com/joho/godotenv"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
	"log"
)

var H *utils.DBHandler

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	H, err = utils.CreateDBHandler()
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	log.Println("Database handler Initialized successfully")

	if err := H.DB.AutoMigrate(
		&models.User{},
		&models.Parent{},
		&models.Program{},
		&models.Class{},
		&models.Pupil{},
		&models.Enrollment{},
		&models.News{},
		&models.Event{},
		&models.Staff{},
	); err != nil {
		H.Logger.Fatal("Failed to migrate tables: ", err)
	}

	H.Logger.Println("Database migration completed successfully")
}

package config

import (
    "log"
    "github.com/joho/godotenv"
    "greaterAltitudeapp/utils"
    "greaterAltitudeapp/models"
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

    if err := H.DB.AutoMigrate(&models.User{}, &models.Parent{}); err != nil {
	    H.Logger.Fatal("Failed to migrate tables: ", err)
    }

    H.Logger.Println("Database migration completed successfully")
}

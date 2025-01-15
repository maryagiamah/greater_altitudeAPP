package config

import (
	"github.com/joho/godotenv"
	"greaterAltitudeapp/models"
	"greaterAltitudeapp/utils"
	"log"
	"os"
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

	if err := H.DB.Migrator().DropTable(
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
		log.Fatal("Failed to drop tables: ", err)
	}

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

	hashedPassword, err := utils.HashPassword(os.Getenv("ADMIN_PWD"))
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	adminUser := models.User{
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: hashedPassword,
		Role:     "admin",
		Mobile:   os.Getenv("ADMIN_MOBILE"),
	}

	if err := H.DB.Create(&adminUser).Error; err != nil {
		log.Fatal("Failed to create admin user: ", err)
	}
	H.Logger.Println("Database migration completed successfully")
}

func CloseDB() {
	sqlDB, err := H.DB.DB()

	if err != nil {
		log.Fatal("Failed to get database: ", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatal("Failed to close database: ", err)
	}
}

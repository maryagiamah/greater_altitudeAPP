package utils

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBHandler struct {
	DB     *gorm.DB
	Logger *log.Logger
}

func CreateDBHandler() (*DBHandler, error) {
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_name := os.Getenv("DB_NAME")
	db_passwd := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", db_host, db_user, db_passwd, db_name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Can't connect to database: %w", err)
	}

	logger := createNewLogger()
	return &DBHandler{DB: db, Logger: logger}, nil
}

func createNewLogger() *log.Logger {
	return log.New(os.Stdout, "[DBHandler] ", log.LstdFlags)
}

func (h *DBHandler) MigrateAllTables(models ...interface{}) error {
	err := h.DB.AutoMigrate(models...)
	if err != nil {
		return fmt.Errorf("Failed to migrate tables")
	}
	return nil
}

package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"greaterAltitudeapp/models"
)

type dbHandler struct {
	DB     *gorm.DB
	Logger *log.Logger
	RDB    *redis.Client
}

var H *dbHandler

func createDBHandler() (*dbHandler, error) {
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_name := os.Getenv("DB_NAME")
	db_passwd := os.Getenv("DB_PASSWORD")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			Colorful:                  true,
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", db_host, db_user, db_passwd, db_name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		TranslateError: true,
	})

	if err != nil {
		return nil, fmt.Errorf("Can't connect to database: %w", err)
	}

	rdb, err := createRedisConnect()
	if err != nil {
		log.Fatalf("Error initializing Redis: %v", err)
	}
	logger := createNewLogger()

	return &dbHandler{DB: db, Logger: logger, RDB: rdb}, nil
}

func createNewLogger() *log.Logger {
	return log.New(os.Stdout, "[DBHandler] ", log.LstdFlags)
}

func createRedisConnect() (*redis.Client, error) {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}

	password := os.Getenv("REDIS_PASSWORD")

	db := 0
	if dbEnv := os.Getenv("REDIS_DB"); dbEnv != "" {
		parsedDB, err := strconv.Atoi(dbEnv)
		if err == nil {
			db = parsedDB
		}
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Redis: %w", err)
	}

	return rdb, nil
}

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	H, err = createDBHandler()
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	log.Println("Database handler Initialized successfully")

	if err := H.DB.Migrator().DropTable(
		&models.User{},
		&models.Message{},
		&models.Parent{},
		&models.Program{},
		&models.Activity{},
		&models.Class{},
		&models.Pupil{},
		&models.Invoice{},
		&models.Payment{},
		&models.Event{},
		&models.Staff{},
		&models.Report{},
		&models.Role{},
		&models.Permission{},
		"role_permissions",
	); err != nil {
		log.Fatalf("Failed to drop tables: %v", err)
	}

	if err := H.DB.AutoMigrate(
		&models.User{},
		&models.Message{},
		&models.Parent{},
		&models.Program{},
		&models.Activity{},
		&models.Class{},
		&models.Pupil{},
		&models.Invoice{},
		&models.Payment{},
		&models.Event{},
		&models.Staff{},
		&models.Report{},
		&models.Role{},
		&models.Permission{},
	); err != nil {
		H.Logger.Fatalf("Failed to migrate tables: %v", err)
	}

	hashedPassword, err := HashPassword(os.Getenv("ADMIN_PWD"))
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
		log.Fatalf("Failed to create admin user: %v", err)
	}
	H.Logger.Println("Database migration completed successfully")
}

func CloseDB() {
	sqlDB, err := H.DB.DB()

	if err != nil {
		log.Fatalf("Failed to get database: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}

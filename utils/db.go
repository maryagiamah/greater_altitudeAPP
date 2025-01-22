package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"github.com/redis/go-redis/v9"
	"greaterAltitudeapp/models"
	"gorm.io/driver/postgres"
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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", db_host, db_user, db_passwd, db_name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Can't connect to database: %w", err)
	}

	logger := createNewLogger()
	rdb := createRedisConnect()
	return &dbHandler{DB: db, Logger: logger, RDB: rdb}, nil
}

func createNewLogger() *log.Logger {
	return log.New(os.Stdout, "[DBHandler] ", log.LstdFlags)
}

func createRedisConnect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}

func InitDB() {
        err := godotenv.Load()
        if err != nil {
                log.Fatal("Error loading .env file")
        }

	H, err = createDBHandler()
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
                &models.Event{},
                &models.Staff{},
        ); err != nil {
                H.Logger.Fatal("Failed to migrate tables: ", err)
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

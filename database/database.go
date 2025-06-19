package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/saarthi123/saarthi-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// InitDB initializes the database connection only once and migrates models
func InitDB() {
	once.Do(func() {
		// Load .env file if it exists
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️ .env file not found, proceeding with system environment variables")
		}

		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			log.Fatal("❌ DATABASE_URL not set in environment")
		}
		fmt.Println("📡 Connecting to database:", dsn)

		dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("❌ Failed to connect to database: %v", err)
		}
		db = dbInstance
		fmt.Println("✅ Successfully connected to database")

		autoMigrateModels()
	})
}

// GetDB returns the active database instance
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("❌ Database not initialized. Call InitDB() in main.go before using DB.")
	}
	return db
}



// autoMigrateModels handles auto-migration for all models in the app
func autoMigrateModels() {
	err := db.AutoMigrate(
		&models.User{},
		&models.BankAccount{},
		&models.Mail{},
		&models.Diploma{},
		&models.UpcomingClass{},
		&models.AttendanceRecord{},
		&models.Notification{},
		&models.NotificationPreferences{},
		&models.Payment{},
		&models.Draft{},
		&models.CampusAttendance{},
		&models.FinancialTip{},
		&models.StudentProgress{},
		&models.User{},
		&models.Query{},
		&models.EmojiCategory{},
		&models.Emoji{},
	)
	if err != nil {
		log.Fatalf("❌ Failed to auto-migrate models: %v", err)
	}
	fmt.Println("📦 Models auto-migrated successfully")
}

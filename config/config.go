package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/saarthi123/saarthi-backend/models" // ✅ Import the models package
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	App  AppConfig
	once sync.Once
)

// AppConfig holds dynamic config for DB, AWS, OpenAI etc.
type AppConfig struct {
	DBURL     string
	AWSRegion string
	S3Bucket  string
	OpenAIKey string
}

// InitConfig loads environment-based config
func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env not found, using system environment variables")
	}

	App = AppConfig{
		DBURL:     os.Getenv("DATABASE_URL"),
		AWSRegion: os.Getenv("AWS_REGION"),
		S3Bucket:  os.Getenv("S3_BUCKET"),
		OpenAIKey: os.Getenv("OPENAI_API_KEY"),
	}

	if App.DBURL == "" {
		log.Fatal("❌ DATABASE_URL is missing from environment")
	}
}

// InitDB initializes database connection once
func InitDB() {
	once.Do(func() {
		var err error
		DB, err = gorm.Open(postgres.Open(App.DBURL), &gorm.Config{})
		if err != nil {
			log.Fatalf("❌ Failed to connect to database: %v", err)
		}
		fmt.Println("✅ Connected to PostgreSQL database")
	})
}

// CheckUserByPhone returns true if a user with the given phone exists
func CheckUserByPhone(phone string) (bool, error) {
	var user models.User
	err := DB.Where("phone_number = ?", phone).First(&user).Error
	if err != nil {
		// Record not found is not an error in this context
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/saarthi123/saarthi-backend/config"
	"github.com/saarthi123/saarthi-backend/routes"
)

func main() {
	// Load env vars first
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system environment variables")
	}

	// Initialize dynamic config and DB
	config.InitConfig()
	config.InitDB()
	

	// Setup Gin with CORS
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Update with frontend domain in production
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Register API routes
	routes.RegisterRoutes(router)

	// Define port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Saarthi backend running at http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}



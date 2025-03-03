package app

import (
	"github.com/gin-contrib/cors"
	"log"
	"os"

	"backend/internal/db"
	"backend/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/*
Start function starts application
-gets env variables
-connects to DB
-starts backend-service
*/
func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading config from environment variables.")
	}

	if err := db.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	if err := db.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	router := gin.Default()
	// To make server work without nginx
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	server.SetupRoutes(router)

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

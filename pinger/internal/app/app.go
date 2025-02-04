package app

import (
	"github.com/joho/godotenv"
	"log"
	"pinger/internal/config"
	"pinger/internal/service"
)

/*
Start function starts application

	-gets env variables
	-starts ping-service
*/
func Start() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found for pinger, using system ENV")
	}

	cfg := config.LoadConfig()

	service.RunPingerLoop(cfg)
}

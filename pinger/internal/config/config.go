package config

import (
	"os"
	"pinger/internal/models"
	"strconv"
	"time"
)

// LoadConfig gets environmental variables
func LoadConfig() *models.Config {
	backendURL := os.Getenv("BACKEND_URL")
	if backendURL == "" {
		backendURL = "http://backend:8080"
	}

	intervalStr := os.Getenv("PING_INTERVAL_SECONDS")
	if intervalStr == "" {
		intervalStr = "10"
	}
	intervalSec, err := strconv.Atoi(intervalStr)
	if err != nil {
		intervalSec = 10
	}

	return &models.Config{
		BackendURL:   backendURL,
		PingInterval: time.Duration(intervalSec) * time.Second,
	}
}

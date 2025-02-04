package config

import (
	"os"
	"strconv"
	"time"
)

// Config model
type Config struct {
	BackendURL   string
	PingInterval time.Duration
}

// LoadConfig gets environmental variables
func LoadConfig() *Config {
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

	return &Config{
		BackendURL:   backendURL,
		PingInterval: time.Duration(intervalSec) * time.Second,
	}
}

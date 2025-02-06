package service

import (
	"log"
	"time"

	"pinger/internal/models"
)

// RunPingerLoop continuously pings targets at intervals specified in cfg.PingInterval.
func RunPingerLoop(cfg *models.Config) {
	ticker := time.NewTicker(cfg.PingInterval)
	defer ticker.Stop()

	for {
		if err := processPing(cfg); err != nil {
			log.Printf("Error in processPing: %v", err)
		}
		<-ticker.C
	}
}

package service

import (
	"log"
	"time"

	"pinger/internal/models"
)

// processPing - one pass: get list of IP addresses from the backend, ping them, send results.
func processPing(cfg *models.Config) error {
	containers, err := fetchContainers(cfg.BackendURL)
	if err != nil {
		log.Println("Error fetching containers:", err)
		return err
	}

	for _, c := range containers {
		success := doPing(c.IPAddress)

		now := time.Now()
		updateData := map[string]interface{}{
			"last_ping_time": now,
		}
		if success {
			updateData["last_success_time"] = now
		}

		err := updateContainer(cfg.BackendURL, c.ID, updateData)
		if err != nil {
			log.Printf("Failed to update container %d: %v", c.ID, err)
		}
	}

	return nil
}

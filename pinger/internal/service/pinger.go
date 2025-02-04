package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	_ "strconv"
	"time"

	"pinger/internal/config"
	"pinger/internal/models"
)

// RunPingerLoop continuously pings targets at intervals specified in the configuration.
func RunPingerLoop(cfg *config.Config) {
	ticker := time.NewTicker(cfg.PingInterval)
	defer ticker.Stop()

	for {
		if err := processPing(cfg); err != nil {
			log.Printf("Error in processPing: %v", err)
		}
		<-ticker.C
	}
}

// processPing - one pass: get list IP addresses from backend, ping them, send results
func processPing(cfg *config.Config) error {
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

// fetchContainers does GET /api/containers
func fetchContainers(baseURL string) ([]models.Container, error) {
	resp, err := http.Get(baseURL + "/api/containers")
	if err != nil {
		log.Println("Error getting containers:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing body:", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET /api/containers returned status %d", resp.StatusCode)
	}

	var containers []models.Container
	if err := json.NewDecoder(resp.Body).Decode(&containers); err != nil {
		return nil, err
	}
	return containers, nil
}

// doPing — pings IP-address. Returns true, if ping is successful
func doPing(ip string) bool {
	cmd := exec.Command("ping", "-c", "1", "-w", "1", ip)

	err := cmd.Run()
	log.Println("From :", ip, "- ", err)
	return err == nil
}

// updateContainer — sends PUT /api/containers/:id
// updateData — map[string]interface{} for updating fields
func updateContainer(baseURL string, id uint, updateData map[string]interface{}) error {
	jsonBytes, err := json.Marshal(updateData)
	if err != nil {
		log.Println("Error marshalling updateData:", err)
		return err
	}

	url := fmt.Sprintf("%s/api/containers/%d", baseURL, id)
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		log.Println("Error creating PUT request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Body = http.NoBody

	req.Body = newRequestBody(jsonBytes)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error executing PUT request:", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing body:", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("PUT /api/containers/%d returned status %d", id, resp.StatusCode)
	}
	return nil
}

func newRequestBody(data []byte) io.ReadCloser {
	return io.NopCloser(bytes.NewReader(data))
}

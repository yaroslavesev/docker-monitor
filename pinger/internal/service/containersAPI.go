package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"pinger/internal/models"
)

// fetchContainers does a GET /api/containers
func fetchContainers(baseURL string) ([]models.Container, error) {
	resp, err := http.Get(baseURL + "/api/containers")
	if err != nil {
		log.Println("Error getting containers:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if cerr := Body.Close(); cerr != nil {
			log.Println("Error closing response body:", cerr)
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

// updateContainer sends PUT /api/containers/id
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

	req.Body = io.NopCloser(bytes.NewReader(jsonBytes))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error executing PUT request:", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		if cerr := Body.Close(); cerr != nil {
			log.Println("Error closing body:", cerr)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("PUT /api/containers/%d returned status %d", id, resp.StatusCode)
	}
	return nil
}

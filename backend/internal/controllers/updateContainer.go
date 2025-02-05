package controllers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// validateAndBindInput - validates input and binds JSON data to the container.
func ValidateAndBindInput(c *gin.Context, input *models.Container) error {
	if err := c.ShouldBindJSON(input); err != nil {
		return err
	}
	return nil
}

// updateContainerFields - updates container fields if provided.
func updateContainerFields(container *models.Container, input models.Container) {
	if input.IPAddress != "" {
		container.IPAddress = input.IPAddress
	}
	if !input.LastPingTime.IsZero() {
		container.LastPingTime = input.LastPingTime
	}
	if !input.LastSuccessTime.IsZero() {
		container.LastSuccessTime = input.LastSuccessTime
	}
	container.UpdatedAt = time.Now()
}

// UpdateContainer - updates container
func UpdateContainer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	container, dbError := repositories.FindById(id)
	if dbError != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "container not found"})
		return
	}

	var input models.Container
	if err := ValidateAndBindInput(c, &input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateContainerFields(container, input)

	if err := repositories.Save(container); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, container)
}

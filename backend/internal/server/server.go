package server

import (
	"backend/internal/db"
	"backend/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//CRUD operations

// GetAllContainers - returns all containers
func GetAllContainers(c *gin.Context) {
	var containers []models.Container
	result := db.DB.Find(&containers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, containers)
}

// CreateContainer - adds new IP-address
func CreateContainer(c *gin.Context) {
	var container models.Container
	if err := c.ShouldBindJSON(&container); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	container.CreatedAt = time.Now()
	container.UpdatedAt = time.Now()

	if err := db.DB.Create(&container).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, container)
}

// UpdateContainer - updates container
func UpdateContainer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var container models.Container
	if err := db.DB.First(&container, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "container not found"})
		return
	}

	var input models.Container
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	if err := db.DB.Save(&container).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, container)
}

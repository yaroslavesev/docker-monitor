package controllers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CreateContainer - adds new IP-address
func CreateContainer(c *gin.Context) {
	var container models.Container
	if err := ValidateAndBindInput(c, &container); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	container.CreatedAt = time.Now()
	container.UpdatedAt = time.Now()

	if err := repositories.Create(&container); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, container)
}

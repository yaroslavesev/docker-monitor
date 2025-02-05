package controllers

import (
	"backend/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReadAllContainers - returns all containers
func ReadAllContainers(c *gin.Context) {
	containers, dbError := repositories.GetAll()
	if dbError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbError.Error()})
		return
	}
	c.JSON(http.StatusOK, containers)
}

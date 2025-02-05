package controllers

import (
	"backend/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DeleteContainer - deletes a container by ID
func DeleteContainer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := repositories.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "container deleted successfully"})
}

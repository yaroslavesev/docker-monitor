package server

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes registers all routes
func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/containers", GetAllContainers)
		api.POST("/containers", CreateContainer)
		api.PUT("/containers/:id", UpdateContainer)
	}
}

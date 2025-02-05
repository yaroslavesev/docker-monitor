package server

import (
	"backend/internal/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes registers all routes
func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/containers", controllers.ReadAllContainers)
		api.POST("/containers", controllers.CreateContainer)
		api.PUT("/containers/:id", controllers.UpdateContainer)
		api.DELETE("/containers/:id", controllers.DeleteContainer)
	}
}

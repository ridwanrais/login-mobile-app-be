package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ridwanrais/login-mobile-app/internal/controller"
)

func SetupHealthsRoutes(router *gin.RouterGroup) {
	health := router.Group("/health")
	{
		health.GET("", controller.GetHealth)
	}
}

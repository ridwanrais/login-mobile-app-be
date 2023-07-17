package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ridwanrais/login-mobile-app/internal/controller"
)

func SetupHealthsRoutes(router *gin.RouterGroup) {
	accounts := router.Group("/health")
	{
		accounts.GET("", controller.GetHealth)
	}
}

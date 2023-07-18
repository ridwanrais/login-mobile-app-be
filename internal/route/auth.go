package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ridwanrais/login-mobile-app/internal/controller"
)

func SetupAuthsRoutes(router *gin.RouterGroup, controller controller.Controllers) {
	accounts := router.Group("/auth")
	{
		accounts.GET("/google/login", controller.HandleGoogleLogin)
		accounts.GET("/google/callback", controller.HandleGoogleLoginCallback)
	}
}

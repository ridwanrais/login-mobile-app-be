package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ridwanrais/login-mobile-app/internal/controller"
)

func SetupAuthsRoutes(router *gin.RouterGroup, controller controller.Controllers) {
	auth := router.Group("/auth")
	{
		auth.GET("/google/login", controller.HandleGoogleLogin)
		auth.GET("/google/callback", controller.HandleGoogleLoginCallback)
		auth.POST("/login", controller.Login)
	}
}

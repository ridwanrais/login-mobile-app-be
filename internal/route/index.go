package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ridwanrais/login-mobile-app/internal/controller"
	// "github.com/ridwanrais/login-mobile-app/internal/route"
)

func SetupRoutes(router *gin.Engine) {
	controller := controller.NewControllers()

	v1 := router.Group("/api/v1")

	{
		SetupHealthsRoutes(v1)

		SetupAccountsRoutes(v1, controller)
		SetupAuthsRoutes(v1, controller)
		// ...
	}
}
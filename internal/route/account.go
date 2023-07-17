package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ridwanrais/login-mobile-app/internal/controller"
)

func SetupAccountsRoutes(router *gin.RouterGroup, controller controller.Controllers) {
	accounts := router.Group("/accounts")
	{
		accounts.GET(":id", func(c *gin.Context) {
			fmt.Println("need implementation")
		})
		// accounts.GET("", ListAccounts)
		accounts.POST("", controller.AddAccount)
		// accounts.DELETE(":id", DeleteAccount)
		// accounts.PATCH(":id", UpdateAccount)
		// accounts.POST(":id/images", UploadAccountImage)
	}
}

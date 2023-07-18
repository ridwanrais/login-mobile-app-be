package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ridwanrais/login-mobile-app/internal/constants"
)

func (c *controllers) HandleGoogleLogin(ctx *gin.Context) {
	config := constants.GetGoogleOauthConfig()

	url := config.AuthCodeURL(os.Getenv("OAUTH_STATE_STRING"))

	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (c *controllers) HandleGoogleLoginCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Code is required"})
		return
	}

	accountID, err := c.usecase.HandleGoogleLoginCallback(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":   "ok",
		"accountID": accountID,
	})
}

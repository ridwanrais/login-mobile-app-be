package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ridwanrais/login-mobile-app/internal/config"
	"github.com/ridwanrais/login-mobile-app/internal/usecase"
)

type controllers struct {
	usecase usecase.Usecases
}

type Controllers interface {
	// account
	AddAccount(c *gin.Context)
	AccountVerificationCallback(ctx *gin.Context)

	// auth
	HandleGoogleLogin(ctx *gin.Context)
	HandleGoogleLoginCallback(ctx *gin.Context)
	Login(ctx *gin.Context)
}

func NewControllers() Controllers {
	return &controllers{
		usecase: config.GetUsecase()}
}

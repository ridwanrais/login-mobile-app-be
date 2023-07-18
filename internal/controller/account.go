package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	validator "github.com/ridwanrais/login-mobile-app/internal/validator/http"
)

func (c *controllers) AddAccount(ctx *gin.Context) {
	account := validator.AddAccountValidator(ctx)

	accountID, err := c.usecase.AddAccount(ctx, *account)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"accountID": accountID,
	})
}

func (c *controllers) AccountVerificationCallback(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Verification ID is required"})
		return
	}

	err := c.usecase.AccountVerificationCallback(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "ok",
		"accountID": "account is verified",
	})
}

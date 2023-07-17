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

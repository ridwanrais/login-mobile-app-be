package controller

import (
	"github.com/gin-gonic/gin"
)

func GetHealth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
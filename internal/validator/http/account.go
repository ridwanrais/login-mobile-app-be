package validator

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-ozzo/ozzo-validation"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/ridwanrais/login-mobile-app/internal/entity"
)

func AddAccountValidator(c *gin.Context) *entity.Account {
	var account entity.Account
	if err := c.ShouldBind(&account); err != nil {
		// Validation failed, handle the error
		if verr, ok := err.(validation.Errors); ok {
			// Validation errors occurred
			c.JSON(http.StatusBadRequest, gin.H{"error": verr.Error()})
			return nil
		}
		// Other errors occurred
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return nil
	}

	if err := validation.ValidateStruct(&account,
		validation.Field(&account.Username, validation.Required, validation.Length(5, 30)),
		validation.Field(&account.Email, validation.Required, is.Email),
		validation.Field(&account.Password, validation.Required),
	); err != nil {
		// Validation failed, handle the error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	return &account
}

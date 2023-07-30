package validator

import (
	"errors"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/ridwanrais/login-mobile-app/internal/entity"
)

func LoginValidator(c *gin.Context) (*entity.Account, error) {
	var account entity.Account
	if err := c.ShouldBind(&account); err != nil {
		// Validation failed, handle the error
		if verr, ok := err.(validation.Errors); ok {
			// Validation errors occurred
			return nil, verr
		}
		// Other errors occurred
		return nil, err
	}

	// Custom validation to ensure that either username or email is present
	if account.Username == "" && account.Email == "" {
		return nil, errors.New("Username or email is required")
	}

	if err := validation.ValidateStruct(&account,
		validation.Field(&account.Username, validation.Length(5, 30)),
		validation.Field(&account.Email, is.Email),
		validation.Field(&account.Password, validation.Required),
	); err != nil {
		// Validation failed, handle the error
		return nil, err
	}

	return &account, nil
}

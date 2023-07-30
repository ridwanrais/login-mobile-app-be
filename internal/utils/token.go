package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ridwanrais/login-mobile-app/internal/entity"
)

// Helper function to generate JWT token
func GenerateJWTToken(userID string) (*entity.JwtToken, error) { // Retrieve the JWT secret key and expiration time from environment variables
	secretKey := os.Getenv("JWT_SECRET_KEY")
	expirationTimeStr := os.Getenv("JWT_EXPIRATION_TIME")
	expirationTime, err := strconv.Atoi(expirationTimeStr)
	if err != nil {
		return nil, err
	}

	// Define the expiration time for the token
	expirationTimeDuration := time.Duration(expirationTime) * time.Hour
	expirationTimeUTC := time.Now().Add(expirationTimeDuration)

	// Create the JWT claims containing user ID and expiration time
	expiresAt := expirationTimeUTC
	issuedAt := time.Now()
	claims := jwt.StandardClaims{
		ExpiresAt: expiresAt.Unix(),
		IssuedAt:  issuedAt.Unix(),
		Subject:   userID,
	}

	// Sign the token with a secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &entity.JwtToken{
		Token: tokenString,
		ExpiresAt: expiresAt,
		IssuedAt: issuedAt,
	}, nil
}

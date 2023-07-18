package utils

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

func GenerateEmailVerificationUrl(verificationID string) string {
	baseURL := os.Getenv("BASE_URL") // Base URL of your application

	// Create a URL object and set the base URL
	u, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	// Set the path and query parameters for the verification URL
	u.Path = "/api/v1/accounts/verify"
	q := u.Query()
	q.Set("id", verificationID)
	u.RawQuery = q.Encode()

	verificationURL := u.String()

	fmt.Println("Verification URL:", verificationURL)

	return verificationURL
}

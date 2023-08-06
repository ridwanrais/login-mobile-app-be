package utils

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
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

func GenerateEmailVerificationUrlV2(verifyUrl, verificationID string) string { // Create a URL object and set the base URL
	u, err := url.Parse(verifyUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Append the verification ID as a path segment
	u.Path = path.Join(u.Path, verificationID)

	verificationURL := u.String()

	fmt.Println("Verification URL:", verificationURL)

	return verificationURL
}

package config

import (
	"context"
	"fmt"
	"log"
	"os"

	validator "github.com/ridwanrais/login-mobile-app/internal/validator/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() (*mongo.Client, error) {
	// Connect to MongoDB using the environment variable
	connectionString := os.Getenv("MONGO_CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("MONGO_CONNECTION_STRING environment variable not set")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// initialize validator for each mongo entity
	validator.SetupDocumentValidator(client)
	
	fmt.Println("Connected to MongoDB!")
	return client, nil
}

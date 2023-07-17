package validator

import (
	"context"
	"log"

	"github.com/ridwanrais/login-mobile-app/internal/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AccountCollectionValidator(client *mongo.Client) {
	// Access the "users" collection
	collection := client.Database(constants.DATABASE_PRIMARY).Collection(constants.COLLECTION_ACCOUNT)

	// Create a unique index on the "username" field
	usernameIndex := mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: options.Index().SetUnique(true),
	}
	if _, err := collection.Indexes().CreateOne(context.Background(), usernameIndex); err != nil {
		log.Fatal(err)
	}

	// Create a unique index on the "email" field
	emailIndex := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	if _, err := collection.Indexes().CreateOne(context.Background(), emailIndex); err != nil {
		log.Fatal(err)
	}
}

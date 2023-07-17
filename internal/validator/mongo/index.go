package validator

import "go.mongodb.org/mongo-driver/mongo"

func SetupDocumentValidator(client *mongo.Client) {
	AccountCollectionValidator(client)
}

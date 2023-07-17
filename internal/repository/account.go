package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/ridwanrais/login-mobile-app/internal/constants"
	"github.com/ridwanrais/login-mobile-app/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repositories) AddAccount(ctx context.Context, account entity.Account) (accountID string, err error) {
	coll := r.db.Collection(constants.COLLECTION_ACCOUNT)

	account = entity.NewAccount(account)
	result, err := coll.InsertOne(ctx, account)
	if err != nil {
		return "", err
	}

	accountID = result.InsertedID.(primitive.ObjectID).Hex()

	return accountID, nil
}

func (r *repositories) GetAccountByFields(ctx context.Context, fields map[string]interface{}) (account *entity.Account, err error) {
	coll := r.db.Collection(constants.COLLECTION_ACCOUNT)

	filter := bson.D{}
	for key, value := range fields {
		filter = append(filter, bson.E{Key: key, Value: value})
	}

	err = coll.FindOne(ctx, filter).Decode(&account)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("No document was found with the provided field(s)")
	}
	if err != nil {
		fmt.Println(err)
	}

	return account, nil
}

package repository

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/ridwanrais/login-mobile-app/internal/constants"
	"github.com/ridwanrais/login-mobile-app/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *repositories) AddAccount(ctx context.Context, account entity.Account) (accountID string, err error) {
	coll := r.db.Collection(constants.COLLECTION_ACCOUNT)

	now := time.Now()
	account.CreatedAt = now
	account.UpdatedAt = now

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

func (r *repositories) GetAccountByID(ctx context.Context, id string) (account *entity.Account, err error) {
	coll := r.db.Collection(constants.COLLECTION_ACCOUNT)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	err = coll.FindOne(ctx, filter).Decode(&account)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("No document was found with the provided ID")
	}
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *repositories) UpdateAccount(ctx context.Context, account entity.Account) (string, error) {
	coll := r.db.Collection(constants.COLLECTION_ACCOUNT)

	now := time.Now()
	account.UpdatedAt = now

	updateFields := bson.M{}
	accountValue := reflect.ValueOf(account)
	accountType := accountValue.Type()

	for i := 0; i < accountValue.NumField(); i++ {
		fieldValue := accountValue.Field(i)
		fieldType := accountType.Field(i)

		// Check if the field is set and not empty
		if fieldValue.IsValid() && !reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldType.Type).Interface()) {
			fieldName := fieldType.Tag.Get("bson")

			if fieldName != "_id,omitempty" {
				updateFields[fieldName] = fieldValue.Interface()
			}
		}
	}

	update := bson.M{
		"$set": updateFields,
	}

	filter := bson.M{"_id": account.MongoID}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", err
	}

	return account.MongoID.Hex(), nil
}

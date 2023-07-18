package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	MongoID        primitive.ObjectID `bson:"_id" json:"-"`
	Username       string             `bson:"username" validate:"required"`
	Email          string             `bson:"email" validate:"required,email"`
	Password       string             `bson:"password"`
	ProfilePhoto   string             `bson:"profilePhoto"`
	RegisterMethod string             `bson:"registerMethod"`
	IsVerified     bool               `bson:"isVerified"`
	CreatedAt      time.Time          `bson:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}

// useless for now but may become useful later
func NewAccount(account Account) Account {
	return account
}

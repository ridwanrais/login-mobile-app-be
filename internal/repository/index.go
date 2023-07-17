package repository

import (
	"context"

	"github.com/ridwanrais/login-mobile-app/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type repositories struct {
	db *mongo.Database
}

type Repositories interface {
	// account
	AddAccount(ctx context.Context, account entity.Account) (accountID string, err error)
	GetAccountByFields(ctx context.Context, fields map[string]interface{}) (account *entity.Account, err error)
}

func NewRepositories(d *mongo.Database) Repositories {
	return &repositories{db: d}
}

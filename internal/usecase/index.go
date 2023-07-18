package usecase

import (
	"context"

	"github.com/ridwanrais/login-mobile-app/internal/entity"
	"github.com/ridwanrais/login-mobile-app/internal/repository"
)

type usecases struct {
	repo repository.Repositories
}

type Usecases interface {
	// account
	AddAccount(ctx context.Context, account entity.Account) (accountID string, err error)
	AccountVerificationCallback(ctx context.Context, verificationID string) error
}

func NewUsecases(r repository.Repositories) Usecases {
	return &usecases{repo: r}
}

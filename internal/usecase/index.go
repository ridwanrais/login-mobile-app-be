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

	// auth
	HandleGoogleLoginCallback(ctx context.Context, code string) (accountID string, err error)
	Login(ctx context.Context, payload entity.Account) (*entity.JwtToken, error)
}

func NewUsecases(r repository.Repositories) Usecases {
	return &usecases{repo: r}
}

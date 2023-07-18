package usecase

import (
	"context"
	"log"

	"github.com/ridwanrais/login-mobile-app/internal/constants"
	"github.com/ridwanrais/login-mobile-app/internal/entity"

	// "golang.org/x/oauth2"

	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func (u *usecases) HandleGoogleLoginCallback(ctx context.Context, code string) (accountID string, err error) {
	config := constants.GetGoogleOauthConfig()

	token, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("Failed to exchange token: %v", err)
		return "", err
	}

	tokenSource := config.TokenSource(ctx, token)

	service, err := oauth2.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		log.Fatalf("Failed to create Google service: %v", err)
		return "", err
	}

	userInfo, err := service.Userinfo.Get().Do()
	if err != nil {
		log.Fatalf("Failed to retrieve user info: %v", err)
		return "", err
	}

	account, err := u.repo.GetAccountByFields(ctx, map[string]interface{}{
		"email": userInfo.Email,
	})

	if account == nil {
		accountID, err := u.repo.AddAccount(ctx, entity.Account{
			Username:       userInfo.Name,
			Email:          userInfo.Email,
			ProfilePhoto:   userInfo.Picture,
			RegisterMethod: constants.GOOGLE_REGISTER,
			IsVerified:     true,
		})
		if err != nil {
			log.Fatalf("Failed to register account: %v", err)
			return "", err
		}

		return accountID, nil
	} else {
		accountID, err := u.repo.UpdateAccount(ctx, entity.Account{
			MongoID:        account.MongoID,
			Username:       userInfo.Name,
			Email:          userInfo.Email,
			ProfilePhoto:   userInfo.Picture,
			RegisterMethod: constants.GOOGLE_REGISTER,
			IsVerified:     true,
		})

		if err != nil {
			log.Fatalf("Failed to register account: %v", err)
			return "", err
		}

		return accountID, nil
	}
}

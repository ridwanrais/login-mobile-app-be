package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/ridwanrais/login-mobile-app/internal/constants"
	"github.com/ridwanrais/login-mobile-app/internal/entity"
	"github.com/ridwanrais/login-mobile-app/internal/utils"

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

func (u *usecases) Login(ctx context.Context, payload entity.Account) (*entity.JwtToken, error) {

	account, err := u.repo.GetAccountByFields(ctx, map[string]interface{}{
		"email": payload.Email,
	})
	if err != nil {
		log.Fatalf("Failed to retrieve user info: %v", err)
		return nil, err
	}

	isCorrectPassword := utils.CheckPasswordHash(payload.Password, account.Password)
	if !isCorrectPassword {
		return nil, errors.New("Incorrect password")
	}

	// Generate the JWT token
	token, err := utils.GenerateJWTToken(account.MongoID.String())
	if err != nil {
		return nil, err
	}

	return token, nil
}

package usecase

import (
	"context"
	"strings"
	"time"

	"github.com/ridwanrais/login-mobile-app/internal/constants"
	"github.com/ridwanrais/login-mobile-app/internal/entity"
	"github.com/ridwanrais/login-mobile-app/internal/utils"
)

func (u *usecases) AddAccount(ctx context.Context, account entity.Account) (accountID string, err error) {
	// // validate if username is already taken
	// res, _ := u.repo.GetAccountByFields(ctx, map[string]interface{}{
	// 	"username": account.Username,
	// })
	// if res != nil {
	// 	return "", utils.CreateError(400, "username tidak tersedia")
	// }

	// fmt.Println("test 1", res);

	// // validate if email is already taken
	// res, _ = u.repo.GetAccountByFields(ctx, map[string]interface{}{
	// 	"email": account.Email,
	// })
	// if res != nil {
	// 	return "", utils.CreateError(400, "email sudah dipakai")
	// }
	// fmt.Println("test 2", res);

	hashedPassword, _ := utils.HashPassword(account.Password)

	now := time.Now()

	newAccount := entity.Account{
		Username:       account.Username,
		Email:          account.Email,
		Password:       hashedPassword,
		ProfilePhoto:   account.ProfilePhoto,
		RegisterMethod: constants.MANUAL_REGISTER,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	accountID, err = u.repo.AddAccount(ctx, newAccount)
	if err != nil {
		// validate if username is already taken
		if strings.Contains(err.Error(), "username") {
			return "", utils.CreateError(400, "username tidak tersedia")
		}

		// validate if email is already taken
		if strings.Contains(err.Error(), "email") {
			return "", utils.CreateError(400, "email sudah dipakai")
		}

		return "", utils.CreateError(500, err.Error())
	}

	return accountID, nil
}

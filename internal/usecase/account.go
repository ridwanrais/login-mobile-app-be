package usecase

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ridwanrais/login-mobile-app/internal/constants"
	"github.com/ridwanrais/login-mobile-app/internal/entity"
	"github.com/ridwanrais/login-mobile-app/internal/tools"
	"github.com/ridwanrais/login-mobile-app/internal/utils"
)

func (u *usecases) AddAccount(ctx context.Context, account entity.Account) (accountID string, err error) {
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
			return "", utils.CreateError(400, "username not available")
		}

		// validate if email is already taken
		if strings.Contains(err.Error(), "email") {
			return "", utils.CreateError(400, "email has already been used")
		}

		return "", utils.CreateError(500, err.Error())
	}

	go func() {
		err = SendAccountVerificationEmail(accountID, account.Email)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	return accountID, nil
}

func SendAccountVerificationEmail(accountID string, recipientEmail string) error {
	keyByte := []byte(os.Getenv("ACCOUNT_VERIFICATION_KEY"))
	verificationID, err := utils.EncryptString(keyByte, accountID)
	if err != nil {
		return err
	}

	verificationUrl := utils.GenerateEmailVerificationUrl(verificationID)

	recipientName := utils.GetEmailName(recipientEmail)
	htmlContent, err := utils.GenerateVerificationEmailHtml(recipientName, verificationUrl)
	if err != nil {
		return errors.New("error generating verification email")
	}

	err = tools.SendEmail(tools.SendEmailParams{
		RecipientEmail: recipientEmail,
		RecipientName:  recipientName,
		Subject:        "Email verification",
		HtmlContent:    htmlContent,
	})
	if err != nil {
		return errors.New("error sending verification email")
	}

	return nil
}

func (u *usecases) AccountVerificationCallback(ctx context.Context, verificationID string) error {
	keyByte := []byte(os.Getenv("ACCOUNT_VERIFICATION_KEY"))
	accountID, err := utils.DecryptString(keyByte, verificationID)
	if err != nil {
		return errors.New("invalid verification id")
	}

	account, _ := u.repo.GetAccountByID(ctx, accountID)
	if account == nil {
		return errors.New("account not found")
	}

	if account.IsVerified {
		return errors.New("account has already been verified")
	}

	err = u.repo.UpdateAccount(ctx, entity.Account{
		IsVerified: true,
		MongoID:    account.MongoID,
	})
	if err != nil {
		return errors.New("error verifivying account: " + err.Error())
	}

	return nil
}

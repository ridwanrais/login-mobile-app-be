package config

import (
	"log"
	"sync"

	"github.com/ridwanrais/login-mobile-app/internal/constants"
	"github.com/ridwanrais/login-mobile-app/internal/repository"
	"github.com/ridwanrais/login-mobile-app/internal/usecase"
)

var oneUc sync.Once
var uc usecase.Usecases

func GetUsecase() usecase.Usecases {
	oneUc.Do(func() {
		uc = usecase.NewUsecases(
			getRepository(),
		)
	})

	return uc
}

var repo repository.Repositories
var oneRepo sync.Once

func getRepository() repository.Repositories {
	mongoClient, err := ConnectToMongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	oneRepo.Do(func() {
		repo = repository.NewRepositories(mongoClient.Database(constants.DATABASE_PRIMARY))
	})

	return repo
}

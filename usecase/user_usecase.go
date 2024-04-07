package usecase

import (
	"sampleAPI/model"
	"sampleAPI/repository"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

// usecaseはrepositoryインターフェースだけに依存する
type userUsecase struct {
	ur repository.IUserRepository
}


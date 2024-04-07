package repository

import "sampleAPI/model"

// UserRepositoryのインターフェースを定義
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
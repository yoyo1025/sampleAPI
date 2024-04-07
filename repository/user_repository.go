package repository

import (
	"sampleAPI/model"

	"gorm.io/gorm"
)

// UserRepositoryのインターフェースを定義
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userResitory struct {
	db *gorm.DB
}
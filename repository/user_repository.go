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

// DBのインスタンスをDIするためのコンストラクタ
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userResitory{db}
}

func (ur *userResitory) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userResitory) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
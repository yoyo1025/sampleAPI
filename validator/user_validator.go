package validator

import "sampleAPI/model"

type IsUerValidator interface {
	UserValidate(user model.User) error
}
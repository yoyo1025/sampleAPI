package controller

import (
	"sampleAPI/usecase"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

// usecaseをDIするためのコンストラクタ
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}
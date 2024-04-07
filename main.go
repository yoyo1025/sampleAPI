package main

import (
	"sampleAPI/controller"
	"sampleAPI/db"
	"sampleAPI/repository"
	"sampleAPI/router"
	"sampleAPI/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewuUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
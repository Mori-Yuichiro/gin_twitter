package main

import (
	"gin-twitter/controllers"
	"gin-twitter/db"
	"gin-twitter/repositories"
	"gin-twitter/router"
	"gin-twitter/usecases"
	"gin-twitter/validators"
	"log"
)

func main() {
	db := db.NewDB()

	userValidator := validators.NewUserValidator()

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepository, userValidator)
	userController := controllers.NewUserController(userUsecase)

	r := router.NewRouter(userController)

	log.Println("Server Started")
	r.Run(":8080")
}

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
	tweetValidator := validators.NewTweetValidator()

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepository, userValidator)
	userController := controllers.NewUserController(userUsecase)

	tweetRepository := repositories.NewTweetRepository(db)
	tweetUsecase := usecases.NewTweetUsecase(tweetRepository, tweetValidator)
	tweetController := controllers.NewTweetController(tweetUsecase)

	r := router.NewRouter(userController, tweetController)

	log.Println("Server Started")
	r.Run(":8080")
}

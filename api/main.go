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
	commentValidator := validators.NewCommentValidator()

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepository, userValidator)
	userController := controllers.NewUserController(userUsecase)

	tweetRepository := repositories.NewTweetRepository(db)
	tweetUsecase := usecases.NewTweetUsecase(tweetRepository, tweetValidator)
	tweetController := controllers.NewTweetController(tweetUsecase)

	commentRepository := repositories.NewCommentRepository(db)
	commentUsecase := usecases.NewCommentUsecase(commentRepository, commentValidator)
	commentController := controllers.NewCommentController(commentUsecase)

	retweetRepository := repositories.NewRetweetRepository(db)
	retweetUsecase := usecases.NewRetweetUsecase(retweetRepository)
	retweetController := controllers.NewRetweetController(retweetUsecase)

	r := router.NewRouter(
		userController,
		tweetController,
		commentController,
		retweetController,
	)

	log.Println("Server Started")
	r.Run(":8080")
}

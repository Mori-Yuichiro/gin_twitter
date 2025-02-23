package main

import (
	"fmt"
	"gin-twitter/controllers"
	"gin-twitter/db"
	"gin-twitter/repositories"
	"gin-twitter/router"
	"gin-twitter/usecases"
	"gin-twitter/validators"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func main() {
	db := db.NewDB()
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		fmt.Println(err.Error())
	}

	imageRepository := repositories.NewImageRepository(cld)
	imageUsecase := usecases.NewImageUsecase(imageRepository)
	imageController := controllers.NewImageController(imageUsecase)

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

	favoriteRepository := repositories.NewFavoriteRepository(db)
	favoriteUsecase := usecases.NewFavoriteUsecase(favoriteRepository)
	favoriteController := controllers.NewFavoriteController(favoriteUsecase)

	bookmarkRepository := repositories.NewBookmarkRepository(db)
	bookmarkUsecase := usecases.NewBookmarkUsecase(bookmarkRepository)
	bookmarkController := controllers.NewBookmarkController(bookmarkUsecase)

	r := router.NewRouter(
		imageController,
		userController,
		tweetController,
		commentController,
		retweetController,
		favoriteController,
		bookmarkController,
	)

	log.Println("Server Started")
	r.Run(":8080")
}

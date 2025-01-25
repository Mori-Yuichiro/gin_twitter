package usecases

import (
	"gin-twitter/models"
	"gin-twitter/repositories"
)

type IFavoriteUsecase interface {
	CreateFavorite(favorite models.Favorite) error
	DeleteFavorite(userId, tweetId uint) error
}

type favoriteUsecase struct {
	fr repositories.IFavoriteRepository
}

func NewFavoriteUsecase(fr repositories.IFavoriteRepository) IFavoriteUsecase {
	return &favoriteUsecase{fr}
}

func (fu *favoriteUsecase) CreateFavorite(favorite models.Favorite) error {
	if err := fu.fr.CreateFavorite(&favorite); err != nil {
		return err
	}
	return nil
}

func (fu *favoriteUsecase) DeleteFavorite(userId, tweetId uint) error {
	if err := fu.fr.DeleteFavorite(userId, tweetId); err != nil {
		return err
	}
	return nil
}

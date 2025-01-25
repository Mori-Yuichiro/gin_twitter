package repositories

import (
	"fmt"
	"gin-twitter/models"

	"gorm.io/gorm"
)

type IFavoriteRepository interface {
	CreateFavorite(favorite *models.Favorite) error
	DeleteFavorite(userId, tweetId uint) error
}

type favoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) IFavoriteRepository {
	return &favoriteRepository{db}
}

func (fr *favoriteRepository) CreateFavorite(favorite *models.Favorite) error {
	if err := fr.db.Create(favorite).Error; err != nil {
		return err
	}
	return nil
}

func (fr *favoriteRepository) DeleteFavorite(userId, tweetId uint) error {
	result := fr.db.Where("tweet_id=? AND user_id=?", tweetId, userId).Delete(&models.Favorite{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

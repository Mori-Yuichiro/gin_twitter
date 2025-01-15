package repositories

import (
	"fmt"
	"gin-twitter/models"

	"gorm.io/gorm"
)

type ITweetRepository interface {
	CreateTweet(tweet *models.Tweet) error
	GetAllTweet(tweets *[]models.Tweet) error
	DeleteTweet(tweetId, userId uint) error
}

type tweetRepository struct {
	db *gorm.DB
}

func NewTweetRepository(db *gorm.DB) ITweetRepository {
	return &tweetRepository{db}
}

func (tr *tweetRepository) CreateTweet(tweet *models.Tweet) error {
	if err := tr.db.Create(tweet).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tweetRepository) GetAllTweet(tweets *[]models.Tweet) error {
	if err := tr.db.Joins("User").Order("created_at DESC").Find(tweets).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tweetRepository) DeleteTweet(tweetId, userId uint) error {
	result := tr.db.Where("id=? AND user_id=?", tweetId, userId).Delete(&models.Tweet{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

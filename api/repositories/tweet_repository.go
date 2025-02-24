package repositories

import (
	"fmt"
	"gin-twitter/models"

	"gorm.io/gorm"
)

type ITweetRepository interface {
	CreateTweet(tweet *models.Tweet) error
	GetAllTweet(tweets *[]models.Tweet) error
	GetTweetById(tweet *models.Tweet, tweetId uint) error
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
	if err := tr.db.Joins("User").Preload("Retweets").Preload("Favorites").Preload("Bookmarks").Order("created_at DESC").Find(tweets).Error; err != nil {
		return err
	}
	return nil
}

func (tr *tweetRepository) GetTweetById(tweet *models.Tweet, tweetId uint) error {
	if err := tr.db.Joins("User").Preload("Comments").Preload("Comments.User").Preload("Retweets").Preload("Favorites").Preload("Bookmarks").Where("tweets.id=?", tweetId).Order("created_at DESC").First(tweet).Error; err != nil {
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

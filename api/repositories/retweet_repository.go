package repositories

import (
	"fmt"
	"gin-twitter/models"

	"gorm.io/gorm"
)

type IRetweetRepository interface {
	CreateRetweet(retweet *models.Retweet) error
	DeleteRetweet(userId, tweetId uint) error
}

type retweetRepository struct {
	db *gorm.DB
}

func NewRetweetRepository(db *gorm.DB) IRetweetRepository {
	return &retweetRepository{db}
}

func (rr *retweetRepository) CreateRetweet(retweet *models.Retweet) error {
	if err := rr.db.Create(retweet).Error; err != nil {
		return err
	}
	return nil
}

func (rr *retweetRepository) DeleteRetweet(userId, tweetId uint) error {
	result := rr.db.Where("tweet_id=? AND user_id=?", tweetId, userId).Delete(&models.Retweet{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

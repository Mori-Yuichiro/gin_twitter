package repositories

import (
	"fmt"
	"gin-twitter/models"

	"gorm.io/gorm"
)

type IBookmarkRepository interface {
	GetBookmarksByUserId(bookmarks *[]models.Bookmark, userId uint) error
	CreateBookmark(bookmark *models.Bookmark) error
	DeleteBookmark(userId, tweetId uint) error
}

type bookmarkRepository struct {
	db *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) IBookmarkRepository {
	return &bookmarkRepository{db}
}

func (br *bookmarkRepository) GetBookmarksByUserId(bookmarks *[]models.Bookmark, userId uint) error {
	if err := br.db.Preload("Tweet").Preload("Tweet.User").Preload("Tweet.Comments").Preload("Tweet.Comments.User").Preload("Tweet.Retweets").Preload("Tweet.Favorites").Preload("Tweet.Bookmarks").Where("user_id=?", userId).Find(bookmarks).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookmarkRepository) CreateBookmark(bookmark *models.Bookmark) error {
	if err := br.db.Create(bookmark).Error; err != nil {
		return err
	}
	return nil
}

func (br *bookmarkRepository) DeleteBookmark(userId, tweetId uint) error {
	result := br.db.Where("tweet_id=? AND user_id=?", tweetId, userId).Delete(&models.Bookmark{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

package repositories

import (
	"fmt"
	"gin-twitter/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	GetUserByEmail(user *models.User, email string) error
	GetUserByUserId(user *models.User, userId uint) error
	CreateUser(user *models.User) error
	UpdateUser(user *models.User, userId uint) error
}

type userRespository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRespository{db}
}

func (ur *userRespository) GetUserByEmail(user *models.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRespository) GetUserByUserId(user *models.User, userId uint) error {
	if err := ur.db.Preload("Tweets", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at desc")
	}).Preload("Tweets.User").Preload("Tweets.Retweets").Preload("Tweets.Favorites").Preload("Tweets.Bookmarks").Preload("Comments").Preload("Comments.User").Preload("Retweets").Preload("Retweets.Tweet").Preload("Retweets.Tweet.User").Preload("Retweets.Tweet.Retweets").Preload("Retweets.Tweet.Favorites").Preload("Retweets.Tweet.Bookmarks").Preload("Favorites").Preload("Favorites.Tweet").Preload("Favorites.Tweet.User").Preload("Favorites.Tweet.Retweets").Preload("Favorites.Tweet.Favorites").Preload("Favorites.Tweet.Bookmarks").Preload("Followers", func(db *gorm.DB) *gorm.DB {
		return db.Where("followed_id=?", userId)
	}).Preload("Followeds", func(db *gorm.DB) *gorm.DB {
		return db.Where("follower_id=?", userId)
	}).Where("id=?", userId).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRespository) CreateUser(user *models.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRespository) UpdateUser(user *models.User, userId uint) error {
	result := ur.db.Model(user).Clauses(clause.Returning{}).Where("id=?", userId).Updates(map[string]interface{}{
		"avator":        user.Avator,
		"display_name":  user.DisplayName,
		"profile_image": user.ProfileImage,
		"bio":           user.Bio,
		"location":      user.Location,
		"website":       user.Website,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

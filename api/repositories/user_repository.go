package repositories

import (
	"gin-twitter/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *models.User, email string) error
	GetUserByUserId(user *models.User, userId uint) error
	CreateUser(user *models.User) error
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
	}).Preload("Tweets.User").Preload("Tweets.Retweets").Preload("Tweets.Favorites").Where("id=?", userId).First(user).Error; err != nil {
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

package models

import "time"

type Favorite struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"userId" gorm:"not null; uniqueIndex:idx_user_tweet_fa"`
	TweetId   uint      `json:"tweetId" gorm:"not null; uniqueIndex:idx_user_tweet_fa"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Tweet     Tweet     `json:"tweet" gorm:"foreignKey:TweetId; constraint:OnDelete:CASCADE"`
}

type FavoriteResponse struct {
	ID        uint          `json:"id" gorm:"primaryKey"`
	UserId    uint          `json:"userId" gorm:"not null"`
	TweetId   uint          `json:"tweetId" gorm:"not null"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Tweet     TweetResponse `json:"tweet"`
}

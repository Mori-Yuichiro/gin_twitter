package models

import "time"

type Bookmark struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"userId" gorm:"not null; uniqueIndex:idx_user_tweet_bo"`
	TweetId   uint      `json:"tweetId" gorm:"not null; uniqueIndex:idx_user_tweet_bo"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Tweet     Tweet     `json:"tweet" gorm:"foreignKey:TweetId; constraint:OnDelete:CASCADE"`
}

type BookmarkResponse struct {
	ID        uint          `json:"id" gorm:"primaryKey"`
	UserId    uint          `json:"userId" gorm:"not null"`
	TweetId   uint          `json:"tweetId" gorm:"not null"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Tweet     TweetResponse `json:"tweet"`
}

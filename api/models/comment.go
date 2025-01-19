package models

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Comment   string    `json:"comment" gorm:"not null"`
	UserId    uint      `json:"userId" gorm:"not null"`
	TweetId   uint      `json:"tweetId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Tweet     Tweet     `json:"tweet" gorm:"foreignKey:TweetId; constraint:OnDelete:CASCADE"`
}

type CommentReponse struct {
	ID        uint          `json:"id" gorm:"primaryKey"`
	Comment   string        `json:"comment" gorm:"not null"`
	UserId    uint          `json:"userId" gorm:"not null"`
	TweetId   uint          `json:"tweetId" gorm:"not null"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	User      UserResponse  `json:"user"`
	Tweet     TweetResponse `json:"tweet"`
}

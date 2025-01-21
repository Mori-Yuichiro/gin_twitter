package models

import "time"

type Tweet struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"not null"`
	UserId    uint      `json:"userId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Comments  []Comment `json:"comments"`
}

type TweetResponse struct {
	ID        uint             `json:"id" gorm:"primaryKey"`
	Content   string           `json:"content" gorm:"not null"`
	UserId    uint             `json:"userId" gorm:"not null"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	User      UserResponse     `json:"user"`
	Comments  []CommentReponse `json:"comments"`
}

package models

import "time"

type User struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	Name         string     `json:"name" gorm:"not null"`
	Email        string     `json:"email" gorm:"not null; unique"`
	Password     string     `json:"password" gorm:"not null"`
	Avator       string     `json:"avator"`
	DisplayName  string     `json:"displayName"`
	ProfileImage string     `json:"profileImage"`
	Bio          string     `json:"bio"`
	Location     string     `json:"location"`
	Website      string     `json:"website"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	Tweets       []Tweet    `json:"tweets"`
	Comments     []Comment  `json:"comments"`
	Retweets     []Retweet  `json:"retweets"`
	Favorites    []Favorite `json:"favorites"`
}

type UserResponse struct {
	ID           uint               `json:"id" gorm:"primaryKey"`
	Name         string             `json:"name" gorm:"not null"`
	Email        string             `json:"email" gorm:"not null; unique"`
	Password     string             `json:"password" gorm:"not null"`
	Avator       string             `json:"avator"`
	DisplayName  string             `json:"displayName"`
	ProfileImage string             `json:"profileImage"`
	Bio          string             `json:"bio"`
	Location     string             `json:"location"`
	Website      string             `json:"website"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"`
	Tweets       []TweetResponse    `json:"tweets"`
	Comments     []CommentReponse   `json:"comments"`
	Retweets     []RetweetResponse  `json:"retweets"`
	Favorites    []FavoriteResponse `json:"favorites"`
}

package models

import (
	"errors"
	"time"
)

type Relationship struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	FollowerId uint      `json:"followerId" gorm:"not null; uniqueIndex:idx_follow_follower"`
	FollowedId uint      `json:"followedId" gorm:"not null; uniqueIndex:idx_follow_follower"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Follower   User      `json:"follower" gorm:"foreignKey:FollowerId; constraint:OnDelete:CASCADE"`
	Followed   User      `json:"followed" gorm:"foreignKey:FollowedId; constraint:OnDelete:CASCADE"`
}

type RelationshipResponse struct {
	ID         uint         `json:"id" gorm:"primaryKey"`
	FollowerId uint         `json:"followerId" gorm:"not null; uniqueIndex:idx_follow_follower"`
	FollowedId uint         `json:"followedId" gorm:"not null; uniqueIndex:idx_follow_follower"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
	Follower   UserResponse `json:"follower"`
	Followed   UserResponse `json:"followed"`
}

func (r *Relationship) BeforeCreate() error {
	if r.FollowerId == r.FollowedId {
		return errors.New("follower and followed cannot be the same")
	}

	return nil
}

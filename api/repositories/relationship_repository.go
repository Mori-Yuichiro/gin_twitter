package repositories

import (
	"fmt"
	"gin-twitter/models"

	"gorm.io/gorm"
)

type IRelationshipRepository interface {
	CreateRelationship(relationship *models.Relationship) error
	DeleteRelationship(followerId, followedId uint) error
}

type relationshipRepository struct {
	db *gorm.DB
}

func NewRelationshipRepository(db *gorm.DB) IRelationshipRepository {
	return &relationshipRepository{db}
}

func (rr *relationshipRepository) CreateRelationship(relationship *models.Relationship) error {
	if err := rr.db.Create(relationship).Error; err != nil {
		return err
	}
	return nil
}

func (rr *relationshipRepository) DeleteRelationship(followerId, followedId uint) error {
	result := rr.db.Where("follower_id=? AND followed_id=?", followerId, followedId).Delete(&models.Relationship{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

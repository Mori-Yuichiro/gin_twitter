package repositories

import (
	"gin-twitter/models"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(comment *models.Comment) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &commentRepository{db}
}

func (cr *commentRepository) CreateComment(comment *models.Comment) error {
	if err := cr.db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

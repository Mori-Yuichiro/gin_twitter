package usecases

import (
	"gin-twitter/models"
	"gin-twitter/repositories"
	"gin-twitter/validators"
)

type ICommentUsecase interface {
	CreateComment(comment models.Comment) error
}

type commentUsecase struct {
	cr repositories.ICommentRepository
	cv validators.ICommentValidator
}

func NewCommentUsecase(cr repositories.ICommentRepository, cv validators.ICommentValidator) ICommentUsecase {
	return &commentUsecase{cr, cv}
}

func (cu *commentUsecase) CreateComment(comment models.Comment) error {
	if err := cu.cv.CreateCommentValidator(comment); err != nil {
		return err
	}

	if err := cu.cr.CreateComment(&comment); err != nil {
		return err
	}
	return nil
}

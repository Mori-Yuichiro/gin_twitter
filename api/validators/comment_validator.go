package validators

import (
	"gin-twitter/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ICommentValidator interface {
	CreateCommentValidator(comment models.Comment) error
}

type commentValidator struct{}

func NewCommentValidator() ICommentValidator {
	return &commentValidator{}
}

func (cv *commentValidator) CreateCommentValidator(comment models.Comment) error {
	return validation.ValidateStruct(&comment,
		validation.Field(
			&comment.Comment,
			validation.Required.Error("comment is required"),
			validation.RuneLength(1, 140).Error("limit min 1 max 140 char"),
		),
	)
}

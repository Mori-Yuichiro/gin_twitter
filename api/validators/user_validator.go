package validators

import (
	"gin-twitter/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	SignUpValidator(user models.User) error
	LogInValidator(user models.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) SignUpValidator(user models.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(8, 30).Error("limit min 8 max 30 char"),
		),
		validation.Field(
			&user.Name,
			validation.Required.Error("name is required"),
		),
	)
}

func (uv *userValidator) LogInValidator(user models.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(8, 30).Error("limit min 8 max 30 char"),
		),
	)
}

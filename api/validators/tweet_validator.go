package validators

import (
	"gin-twitter/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITweetValidator interface {
	CreateTweetValidator(tweet models.Tweet) error
}

type tweetValidator struct{}

func NewTweetValidator() ITweetValidator {
	return &tweetValidator{}
}

func (tv *tweetValidator) CreateTweetValidator(tweet models.Tweet) error {
	return validation.ValidateStruct(&tweet,
		validation.Field(
			&tweet.Content,
			validation.Required.Error("content is required"),
			validation.RuneLength(1, 140).Error("limit min 1 max 140 char"),
		),
	)
}

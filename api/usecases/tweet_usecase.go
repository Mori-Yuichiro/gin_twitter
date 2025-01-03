package usecases

import (
	"gin-twitter/models"
	"gin-twitter/repositories"
	"gin-twitter/validators"
)

type ITweetUsecase interface {
	CreateTweet(tweet models.Tweet) error
}

type tweetUsecase struct {
	tr repositories.ITweetRepository
	tv validators.ITweetValidator
}

func NewTweetUsecase(tr repositories.ITweetRepository, tv validators.ITweetValidator) ITweetUsecase {
	return &tweetUsecase{tr, tv}
}

func (tu *tweetUsecase) CreateTweet(tweet models.Tweet) error {
	if err := tu.tv.CreateTweetValidator(tweet); err != nil {
		return err
	}

	if err := tu.tr.CreateTweet(&tweet); err != nil {
		return err
	}

	return nil
}

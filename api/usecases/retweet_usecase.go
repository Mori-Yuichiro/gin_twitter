package usecases

import (
	"gin-twitter/models"
	"gin-twitter/repositories"
)

type IRetweetUsecase interface {
	CreateRetweet(retweet models.Retweet) error
	DeleteRetweet(userId, tweetId uint) error
}

type retweetUsecase struct {
	rr repositories.IRetweetRepository
}

func NewRetweetUsecase(rr repositories.IRetweetRepository) IRetweetUsecase {
	return &retweetUsecase{rr}
}

func (ru *retweetUsecase) CreateRetweet(retweet models.Retweet) error {
	if err := ru.rr.CreateRetweet(&retweet); err != nil {
		return err
	}
	return nil
}

func (ru *retweetUsecase) DeleteRetweet(userId, tweetId uint) error {
	if err := ru.rr.DeleteRetweet(userId, tweetId); err != nil {
		return err
	}
	return nil
}

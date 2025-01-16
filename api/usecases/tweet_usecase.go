package usecases

import (
	"gin-twitter/models"
	"gin-twitter/repositories"
	"gin-twitter/validators"
)

type ITweetUsecase interface {
	CreateTweet(tweet models.Tweet) error
	GetAllTweet() ([]models.TweetResponse, error)
	GetTweetById(tweetId uint) (models.TweetResponse, error)
	DeleteTweet(tweetId, userId uint) error
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

func (tu *tweetUsecase) GetAllTweet() ([]models.TweetResponse, error) {
	tweets := []models.Tweet{}
	if err := tu.tr.GetAllTweet(&tweets); err != nil {
		return []models.TweetResponse{}, err
	}

	resTweets := []models.TweetResponse{}
	for _, v := range tweets {
		user := models.UserResponse{
			ID:           v.User.ID,
			Name:         v.User.Name,
			Email:        v.User.Email,
			Password:     v.User.Password,
			Avator:       v.User.Avator,
			DisplayName:  v.User.DisplayName,
			ProfileImage: v.User.ProfileImage,
			Bio:          v.User.Bio,
			Location:     v.User.Location,
			Website:      v.User.Website,
			CreatedAt:    v.User.CreatedAt,
			UpdatedAt:    v.User.UpdatedAt,
		}
		tweet := models.TweetResponse{
			ID:        v.ID,
			Content:   v.Content,
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User:      user,
		}
		resTweets = append(resTweets, tweet)
	}

	return resTweets, nil
}

func (tu *tweetUsecase) GetTweetById(tweetId uint) (models.TweetResponse, error) {
	tweet := models.Tweet{}
	if err := tu.tr.GetTweetById(&tweet, tweetId); err != nil {
		return models.TweetResponse{}, err
	}

	user := models.UserResponse{
		ID:           tweet.User.ID,
		Name:         tweet.User.Name,
		Email:        tweet.User.Email,
		Password:     tweet.User.Password,
		Avator:       tweet.User.Avator,
		DisplayName:  tweet.User.DisplayName,
		ProfileImage: tweet.User.ProfileImage,
		Bio:          tweet.User.Bio,
		Location:     tweet.User.Location,
		Website:      tweet.User.Website,
		CreatedAt:    tweet.User.CreatedAt,
		UpdatedAt:    tweet.User.UpdatedAt,
	}

	resTweet := models.TweetResponse{
		ID:        tweet.ID,
		Content:   tweet.Content,
		UserId:    tweet.UserId,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
		User:      user,
	}

	return resTweet, nil
}

func (tu *tweetUsecase) DeleteTweet(tweetId, userId uint) error {
	if err := tu.tr.DeleteTweet(tweetId, userId); err != nil {
		return err
	}
	return nil
}

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

		comments := []models.CommentReponse{}
		if len(v.Comments) > 0 {
			for _, comm := range v.Comments {
				comment := models.CommentReponse{
					ID:        comm.ID,
					Comment:   comm.Comment,
					UserId:    comm.UserId,
					TweetId:   comm.TweetId,
					CreatedAt: comm.CreatedAt,
					UpdatedAt: comm.UpdatedAt,
					User:      models.UserResponse(comm.User),
				}
				comments = append(comments, comment)
			}
		}

		retweets := []models.RetweetResponse{}
		if len(v.Retweets) > 0 {
			for _, ret := range v.Retweets {
				retweet := models.RetweetResponse{
					ID:        ret.ID,
					UserId:    ret.UserId,
					TweetId:   ret.TweetId,
					CreatedAt: ret.CreatedAt,
					UpdatedAt: ret.UpdatedAt,
				}
				retweets = append(retweets, retweet)
			}
		}

		tweet := models.TweetResponse{
			ID:        v.ID,
			Content:   v.Content,
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User:      user,
			Comments:  comments,
			Retweets:  retweets,
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

	comments := []models.CommentReponse{}
	if len(tweet.Comments) > 0 {
		for _, v := range tweet.Comments {
			comment := models.CommentReponse{
				ID:        v.ID,
				Comment:   v.Comment,
				UserId:    v.UserId,
				TweetId:   v.TweetId,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				User:      models.UserResponse(v.User),
			}
			comments = append(comments, comment)
		}
	}

	retweets := []models.RetweetResponse{}
	if len(tweet.Retweets) > 0 {
		for _, ret := range tweet.Retweets {
			retweet := models.RetweetResponse{
				ID:        ret.ID,
				UserId:    ret.UserId,
				TweetId:   ret.TweetId,
				CreatedAt: ret.CreatedAt,
				UpdatedAt: ret.UpdatedAt,
			}
			retweets = append(retweets, retweet)
		}
	}

	resTweet := models.TweetResponse{
		ID:        tweet.ID,
		Content:   tweet.Content,
		UserId:    tweet.UserId,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
		User:      models.UserResponse(tweet.User),
		Comments:  comments,
		Retweets:  retweets,
	}

	return resTweet, nil
}

func (tu *tweetUsecase) DeleteTweet(tweetId, userId uint) error {
	if err := tu.tr.DeleteTweet(tweetId, userId); err != nil {
		return err
	}
	return nil
}

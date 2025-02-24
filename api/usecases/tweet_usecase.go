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
				commUser := models.UserResponse{
					ID:           comm.User.ID,
					Name:         comm.User.Name,
					Email:        comm.User.Email,
					Password:     comm.User.Password,
					Avator:       comm.User.Avator,
					DisplayName:  comm.User.DisplayName,
					ProfileImage: comm.User.ProfileImage,
					Bio:          comm.User.Bio,
					Location:     comm.User.Location,
					Website:      comm.User.Website,
					CreatedAt:    comm.User.CreatedAt,
					UpdatedAt:    comm.User.UpdatedAt,
				}
				comment := models.CommentReponse{
					ID:        comm.ID,
					Comment:   comm.Comment,
					UserId:    comm.UserId,
					TweetId:   comm.TweetId,
					CreatedAt: comm.CreatedAt,
					UpdatedAt: comm.UpdatedAt,
					User:      commUser,
					// User:      models.UserResponse(comm.User),
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

		favorites := []models.FavoriteResponse{}
		if len(v.Favorites) > 0 {
			for _, fav := range v.Favorites {
				favorite := models.FavoriteResponse{
					ID:        fav.ID,
					UserId:    fav.UserId,
					TweetId:   fav.TweetId,
					CreatedAt: fav.CreatedAt,
					UpdatedAt: fav.UpdatedAt,
				}
				favorites = append(favorites, favorite)
			}
		}

		bookmarks := []models.BookmarkResponse{}
		if len(v.Bookmarks) > 0 {
			for _, book := range v.Bookmarks {
				bookmark := models.BookmarkResponse{
					ID:        book.ID,
					UserId:    book.UserId,
					TweetId:   book.TweetId,
					CreatedAt: book.CreatedAt,
					UpdatedAt: book.UpdatedAt,
				}
				bookmarks = append(bookmarks, bookmark)
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
			Favorites: favorites,
			Bookmarks: bookmarks,
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

	comments := []models.CommentReponse{}
	if len(tweet.Comments) > 0 {
		for _, v := range tweet.Comments {
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
			comment := models.CommentReponse{
				ID:        v.ID,
				Comment:   v.Comment,
				UserId:    v.UserId,
				TweetId:   v.TweetId,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				User:      user,
				// User:      models.UserResponse(v.User),
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

	favorites := []models.FavoriteResponse{}
	if len(tweet.Favorites) > 0 {
		for _, fav := range tweet.Favorites {
			favorite := models.FavoriteResponse{
				ID:        fav.ID,
				UserId:    fav.UserId,
				TweetId:   fav.TweetId,
				CreatedAt: fav.CreatedAt,
				UpdatedAt: fav.UpdatedAt,
			}
			favorites = append(favorites, favorite)
		}
	}

	bookmarks := []models.BookmarkResponse{}
	if len(tweet.Bookmarks) > 0 {
		for _, book := range tweet.Bookmarks {
			bookmark := models.BookmarkResponse{
				ID:        book.ID,
				UserId:    book.UserId,
				TweetId:   book.TweetId,
				CreatedAt: book.CreatedAt,
				UpdatedAt: book.UpdatedAt,
			}
			bookmarks = append(bookmarks, bookmark)
		}
	}

	resTweet := models.TweetResponse{
		ID:        tweet.ID,
		Content:   tweet.Content,
		UserId:    tweet.UserId,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
		User:      user,
		Comments:  comments,
		Retweets:  retweets,
		Favorites: favorites,
		Bookmarks: bookmarks,
		// User:      models.UserResponse(tweet.User),
	}

	return resTweet, nil
}

func (tu *tweetUsecase) DeleteTweet(tweetId, userId uint) error {
	if err := tu.tr.DeleteTweet(tweetId, userId); err != nil {
		return err
	}
	return nil
}

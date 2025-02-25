package usecases

import (
	"gin-twitter/models"
	"gin-twitter/repositories"
)

type IBookmarkUsecase interface {
	GetBookmarksByUserId(userId uint) ([]models.BookmarkResponse, error)
	CreateBookmark(bookmark models.Bookmark) error
	DeleteBookmark(userId, tweetId uint) error
}

type bookmarkUsecase struct {
	br repositories.IBookmarkRepository
}

func NewBookmarkUsecase(br repositories.IBookmarkRepository) IBookmarkUsecase {
	return &bookmarkUsecase{br}
}

func (bu *bookmarkUsecase) GetBookmarksByUserId(userId uint) ([]models.BookmarkResponse, error) {
	bookmarks := []models.Bookmark{}
	if err := bu.br.GetBookmarksByUserId(&bookmarks, userId); err != nil {
		return []models.BookmarkResponse{}, err
	}

	resBookmarks := []models.BookmarkResponse{}
	for _, v := range bookmarks {
		user := models.UserResponse{
			ID:           v.Tweet.User.ID,
			Name:         v.Tweet.User.Name,
			Email:        v.Tweet.User.Email,
			Password:     v.Tweet.User.Password,
			Avator:       v.Tweet.User.Avator,
			DisplayName:  v.Tweet.User.DisplayName,
			ProfileImage: v.Tweet.User.ProfileImage,
			Bio:          v.Tweet.User.Bio,
			Location:     v.Tweet.User.Location,
			Website:      v.Tweet.User.Website,
			CreatedAt:    v.Tweet.User.CreatedAt,
			UpdatedAt:    v.Tweet.User.UpdatedAt,
		}

		comments := []models.CommentReponse{}
		if len(v.Tweet.Comments) > 0 {
			for _, comm := range v.Tweet.Comments {
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
				}
				comments = append(comments, comment)
			}
		}

		retweets := []models.RetweetResponse{}
		if len(v.Tweet.Retweets) > 0 {
			for _, ret := range v.Tweet.Retweets {
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
		if len(v.Tweet.Favorites) > 0 {
			for _, fav := range v.Tweet.Favorites {
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
		if len(v.Tweet.Bookmarks) > 0 {
			for _, book := range v.Tweet.Bookmarks {
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
			ID:        v.Tweet.ID,
			Content:   v.Tweet.Content,
			UserId:    v.Tweet.UserId,
			CreatedAt: v.Tweet.CreatedAt,
			UpdatedAt: v.Tweet.UpdatedAt,
			User:      user,
			Comments:  comments,
			Retweets:  retweets,
			Favorites: favorites,
			Bookmarks: bookmarks,
		}
		bookmark := models.BookmarkResponse{
			ID:        v.ID,
			UserId:    v.UserId,
			TweetId:   v.TweetId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Tweet:     tweet,
		}

		resBookmarks = append(resBookmarks, bookmark)
	}

	return resBookmarks, nil
}

func (bu *bookmarkUsecase) CreateBookmark(bookmark models.Bookmark) error {
	if err := bu.br.CreateBookmark(&bookmark); err != nil {
		return err
	}
	return nil
}

func (bu *bookmarkUsecase) DeleteBookmark(userId, tweetId uint) error {
	if err := bu.br.DeleteBookmark(userId, tweetId); err != nil {
		return err
	}
	return nil
}

package usecases

import (
	"gin-twitter/models"
	"gin-twitter/repositories"
	"gin-twitter/validators"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user models.User) error
	LogIn(user models.User) (string, error)
	GetUserByUserId(userId uint) (models.UserResponse, error)
	UpdateUser(user models.User, userId uint) error
}

type userUsecase struct {
	ur repositories.IUserRepository
	uv validators.IUserValidator
}

func NewUserUsecase(ur repositories.IUserRepository, uv validators.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user models.User) error {
	if err := uu.uv.SignUpValidator(user); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hash),
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return err
	}
	return nil
}

func (uu *userUsecase) LogIn(user models.User) (string, error) {
	if err := uu.uv.LogInValidator(user); err != nil {
		return "", err
	}

	storeUser := models.User{}
	if err := uu.ur.GetUserByEmail(&storeUser, user.Email); err != nil {
		return "", err
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(storeUser.Password),
		[]byte(user.Password),
	)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": storeUser.ID,
			"exp":    time.Now().Add(time.Hour * 5).Unix(),
		},
	)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	return tokenString, nil
}

func (uu *userUsecase) GetUserByUserId(userId uint) (models.UserResponse, error) {
	user := models.User{}
	if err := uu.ur.GetUserByUserId(&user, userId); err != nil {
		return models.UserResponse{}, err
	}

	tweets := []models.TweetResponse{}
	if len(user.Tweets) > 0 {
		for _, v := range user.Tweets {
			tweetUser := models.UserResponse{
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
				User:      tweetUser,
				Retweets:  retweets,
				Favorites: favorites,
				Bookmarks: bookmarks,
			}

			tweets = append(tweets, tweet)
		}
	}

	comments := []models.CommentReponse{}
	if len(user.Comments) > 0 {
		for _, comm := range user.Comments {
			user := models.UserResponse{
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
				User:      user,
			}
			comments = append(comments, comment)
		}
	}

	retweets := []models.RetweetResponse{}
	if len(user.Retweets) > 0 {
		for _, ret := range user.Retweets {
			retTwUser := models.UserResponse{
				ID:           ret.Tweet.User.ID,
				Name:         ret.Tweet.User.Name,
				Email:        ret.Tweet.User.Email,
				Password:     ret.Tweet.User.Password,
				Avator:       ret.Tweet.User.Avator,
				DisplayName:  ret.Tweet.User.DisplayName,
				ProfileImage: ret.Tweet.User.ProfileImage,
				Bio:          ret.Tweet.User.Bio,
				Location:     ret.Tweet.User.Location,
				Website:      ret.Tweet.User.Website,
				CreatedAt:    ret.Tweet.User.CreatedAt,
				UpdatedAt:    ret.Tweet.User.UpdatedAt,
			}

			retTwRets := []models.RetweetResponse{}
			if len(ret.Tweet.Retweets) > 0 {
				for _, retTweetRet := range ret.Tweet.Retweets {
					retTwRet := models.RetweetResponse{
						ID:        retTweetRet.ID,
						UserId:    retTweetRet.UserId,
						TweetId:   retTweetRet.TweetId,
						CreatedAt: retTweetRet.CreatedAt,
						UpdatedAt: retTweetRet.UpdatedAt,
					}
					retTwRets = append(retTwRets, retTwRet)
				}
			}

			retTwFavs := []models.FavoriteResponse{}
			if len(ret.Tweet.Favorites) > 0 {
				for _, retTweetFav := range ret.Tweet.Favorites {
					retTwFav := models.FavoriteResponse{
						ID:        retTweetFav.ID,
						UserId:    retTweetFav.UserId,
						TweetId:   retTweetFav.TweetId,
						CreatedAt: retTweetFav.CreatedAt,
						UpdatedAt: retTweetFav.UpdatedAt,
					}
					retTwFavs = append(retTwFavs, retTwFav)
				}
			}

			retTwBooks := []models.BookmarkResponse{}
			if len(ret.Tweet.Bookmarks) > 0 {
				for _, retTweetBook := range ret.Tweet.Bookmarks {
					favTwBook := models.BookmarkResponse{
						ID:        retTweetBook.ID,
						UserId:    retTweetBook.UserId,
						TweetId:   retTweetBook.TweetId,
						CreatedAt: retTweetBook.CreatedAt,
						UpdatedAt: retTweetBook.UpdatedAt,
					}
					retTwBooks = append(retTwBooks, favTwBook)
				}
			}

			tweet := models.TweetResponse{
				ID:        ret.Tweet.ID,
				Content:   ret.Tweet.Content,
				UserId:    ret.Tweet.UserId,
				CreatedAt: ret.Tweet.CreatedAt,
				UpdatedAt: ret.Tweet.UpdatedAt,
				User:      retTwUser,
				Retweets:  retTwRets,
				Favorites: retTwFavs,
			}

			retweet := models.RetweetResponse{
				ID:        ret.ID,
				UserId:    ret.UserId,
				TweetId:   ret.TweetId,
				CreatedAt: ret.CreatedAt,
				UpdatedAt: ret.UpdatedAt,
				Tweet:     tweet,
			}

			retweets = append(retweets, retweet)
		}
	}

	favorites := []models.FavoriteResponse{}
	if len(user.Favorites) > 0 {
		for _, fav := range user.Favorites {
			favTwUser := models.UserResponse{
				ID:           fav.Tweet.User.ID,
				Name:         fav.Tweet.User.Name,
				Email:        fav.Tweet.User.Email,
				Password:     fav.Tweet.User.Password,
				Avator:       fav.Tweet.User.Avator,
				DisplayName:  fav.Tweet.User.DisplayName,
				ProfileImage: fav.Tweet.User.ProfileImage,
				Bio:          fav.Tweet.User.Bio,
				Location:     fav.Tweet.User.Location,
				Website:      fav.Tweet.User.Website,
				CreatedAt:    fav.Tweet.User.CreatedAt,
				UpdatedAt:    fav.Tweet.User.UpdatedAt,
			}

			favTwRets := []models.RetweetResponse{}
			if len(fav.Tweet.Retweets) > 0 {
				for _, favTweetRet := range fav.Tweet.Retweets {
					retTwRet := models.RetweetResponse{
						ID:        favTweetRet.ID,
						UserId:    favTweetRet.UserId,
						TweetId:   favTweetRet.TweetId,
						CreatedAt: favTweetRet.CreatedAt,
						UpdatedAt: favTweetRet.UpdatedAt,
					}
					favTwRets = append(favTwRets, retTwRet)
				}
			}

			favTwFavs := []models.FavoriteResponse{}
			if len(fav.Tweet.Favorites) > 0 {
				for _, favTweetFav := range fav.Tweet.Favorites {
					retTwFav := models.FavoriteResponse{
						ID:        favTweetFav.ID,
						UserId:    favTweetFav.UserId,
						TweetId:   favTweetFav.TweetId,
						CreatedAt: favTweetFav.CreatedAt,
						UpdatedAt: favTweetFav.UpdatedAt,
					}
					favTwFavs = append(favTwFavs, retTwFav)
				}
			}

			favTwBooks := []models.BookmarkResponse{}
			if len(fav.Tweet.Bookmarks) > 0 {
				for _, favTweetBook := range fav.Tweet.Bookmarks {
					favTwBook := models.BookmarkResponse{
						ID:        favTweetBook.ID,
						UserId:    favTweetBook.UserId,
						TweetId:   favTweetBook.TweetId,
						CreatedAt: favTweetBook.CreatedAt,
						UpdatedAt: favTweetBook.UpdatedAt,
					}
					favTwBooks = append(favTwBooks, favTwBook)
				}
			}

			tweet := models.TweetResponse{
				ID:        fav.Tweet.ID,
				Content:   fav.Tweet.Content,
				UserId:    fav.Tweet.UserId,
				CreatedAt: fav.Tweet.CreatedAt,
				UpdatedAt: fav.Tweet.UpdatedAt,
				User:      favTwUser,
				Retweets:  favTwRets,
				Favorites: favTwFavs,
				Bookmarks: favTwBooks,
			}

			favorite := models.FavoriteResponse{
				ID:        fav.ID,
				UserId:    fav.UserId,
				TweetId:   fav.TweetId,
				CreatedAt: fav.CreatedAt,
				UpdatedAt: fav.UpdatedAt,
				Tweet:     tweet,
			}

			favorites = append(favorites, favorite)
		}
	}

	resUser := models.UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		Avator:       user.Avator,
		DisplayName:  user.DisplayName,
		ProfileImage: user.ProfileImage,
		Bio:          user.Bio,
		Location:     user.Location,
		Website:      user.Website,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Tweets:       tweets,
		Comments:     comments,
		Retweets:     retweets,
		Favorites:    favorites,
	}

	return resUser, nil
}

func (uu *userUsecase) UpdateUser(user models.User, userId uint) error {
	if err := uu.ur.UpdateUser(&user, userId); err != nil {
		return err
	}
	return nil
}
